package hadoop

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	jsonparser "github.com/influxdata/telegraf/plugins/parsers/json"
)

var javaBeanNames = []string{
	"java.lang:type=Threading",
	"java.lang:type=OperatingSystem",
}

var JournalNode = "hadoop_journalnode"
var NameNode = "hadoop_namenode"
var DataNode = "hadoop_datanode"

type Hadoop struct {
	JournalNodes     []string `toml:"journal_nodes"`
	JournalBeanNames []string `toml:"journal_bean_names"`
	DataNodes        []string `toml:"data_nodes"`
	DataBeanNames    []string `toml:"data_bean_names"`
	NameNodes        []string `toml:"name_nodes"`
	NameBeanNames    []string `toml:"name_bean_names"`
	CollectAllBeans  bool     `toml:"collect_all_beans"`
}

func contains(a interface{}, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getNameType(name interface{}) (nameType string, err error) {
	stringName, _ := name.(string)

	typeMatch, _ := regexp.Compile(`type=(?P<type>\w+)`)
	typeFound := typeMatch.FindStringSubmatch(stringName)

	nameMatch, _ := regexp.Compile(`name=(?P<name>.*)`)
	nameFound := nameMatch.FindStringSubmatch(stringName)

	if len(typeFound) > 0 {
		nameType = typeFound[1]
	} else {
		if len(nameFound) > 0 {
			nameType = nameFound[1]
		} else {
			err = errors.New("Did not find Bean type or name")
		}
	}
	return
}

// SampleConfig returns a sample configuration block
func (h *Hadoop) SampleConfig() string {
	return `journalnodes = ["localhost:8080"]`
}

// Description just returns a short description of the Hadoop plugin
func (h *Hadoop) Description() string {
	return "Telegraf plugin for gathering metrics from Hadoop hosts"
}

// Gather() metrics from given list of Hadoop servers
func (h *Hadoop) Gather(acc telegraf.Accumulator) error {
	var wg sync.WaitGroup
	var errorChannel chan error
	errorChannel = make(chan error, len(h.JournalNodes)+len(h.NameNodes)+len(h.DataNodes))

	for _, node := range h.JournalNodes {
		wg.Add(1)
		go func(jnodes string) {
			defer wg.Done()
			acc.AddError(h.gatherMetrics(jnodes, JournalNode, acc))
		}(node)
	}

	for _, node := range h.NameNodes {
		wg.Add(1)
		go func(nnodes string) {
			defer wg.Done()
			acc.AddError(h.gatherMetrics(nnodes, NameNode, acc))
		}(node)
	}

	for _, node := range h.DataNodes {
		wg.Add(1)
		go func(dnodes string) {
			defer wg.Done()
			acc.AddError(h.gatherMetrics(dnodes, DataNode, acc))
		}(node)
	}

	wg.Wait()
	close(errorChannel)
	errorStrings := []string{}

	// Gather all errors for returning them at once
	for err := range errorChannel {
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
		}
	}

	if len(errorStrings) > 0 {
		return errors.New(strings.Join(errorStrings, "\n"))
	}

	return nil
}

var tr = &http.Transport{
	ResponseHeaderTimeout: time.Duration(3 * time.Second),
}

var client = &http.Client{
	Transport: tr,
	Timeout:   time.Duration(4 * time.Second),
}

func (h *Hadoop) gatherMetrics(endpoint string, nodeType string, acc telegraf.Accumulator) error {
	var jsonOut map[string][]interface{}

	var beanNames []string
	var prefixName string

	tags := map[string]string{
		"server": endpoint,
	}
	if nodeType == JournalNode {
		beanNames = h.JournalBeanNames
		prefixName = JournalNode
		tags["type"] = "journal"

	} else if nodeType == NameNode {
		beanNames = h.NameBeanNames
		prefixName = NameNode
		tags["type"] = "name"
	} else {
		beanNames = h.DataBeanNames
		prefixName = DataNode
		tags["type"] = "data"
	}
	if h.CollectAllBeans != true && len(beanNames) == 0 {
		err := errors.New("collect_all_beans was not set to true or one of the list of bean_names was not set")
		return err
	}

	resp, err := client.Get(endpoint + "/jmx")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(data), &jsonOut); err != nil {
		return errors.New("Error decoding JSON response")
	}

	for _, val := range jsonOut["beans"] {
		name := (val.(map[string]interface{})["name"])
		context := (val.(map[string]interface{})["tag.Context"])
		var beanName string
		if context != nil {
			beanName = context.(string)
		} else {
			beanName, err = getNameType(name)
			if err != nil {
				return err
			}
		}
		tags["bean_name"] = beanName
		jf := jsonparser.JSONFlattener{}
		err = jf.FlattenJSON("", val)
		if err != nil {
			return err
		}
		if h.CollectAllBeans == true {
			if len(jf.Fields) > 0 {
				acc.AddFields(prefixName+"_"+beanName, jf.Fields, tags)
			}

		} else if contains(name, beanNames) {
			if len(jf.Fields) > 0 {
				acc.AddFields(prefixName+"_"+beanName, jf.Fields, tags)
			}
		}
	}
	return nil
}

func init() {
	inputs.Add("hadoop", func() telegraf.Input {
		return &Hadoop{}
	})
}
