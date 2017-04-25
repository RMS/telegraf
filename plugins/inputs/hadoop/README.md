# Hadoop Input Plugin

This input plugin gathers metrics from Hadoop.
For more information, please check the [Hadoop: Metrics](https://hadoop.apache.org/docs/current/hadoop-project-dist/hadoop-common/Metrics.html) page.

## Configuration

```toml
# Telegraf plugin for gathering metrics from Hadoop hosts.

# Below is an example 
# Only pass the Bean Names if you want the default list to be overridden.
[[inputs.hadoop]]
  ## A list of Hadoop Journal Nodes. By default it will not collect data unless at least one node is specified.
  journal_nodes = ["http://localhost:8480"]
  ## Example of Metric Bean names to be collected:
  journal_bean_names = [
    "Hadoop:service=JournalNode,name=MetricsSystem,sub=Stats",
    "Hadoop:service=JournalNode,name=RpcDetailedActivityForPort8485",
    "Hadoop:service=JournalNode:service=JournalNode,name=RpcActivityForPort8485".
    "Hadoop:service=JournalNode,name=UgiMetrics",
    "Hadoop:service=JournalNode,name=JvmMetrics",
    "java.lang:type=Threading",
    "java.lang:type=OperatingSystem"
  ]
}
```

## Example of collecting all beans for Hadoop Name Nodes
```
[[inputs.hadoop]]
  ## A list of Hadoop Name Nodes and collect all beans.
  name_nodes = ["http://localhost:50070"]
  collect_all_beans = true
  ]
}
```
- **journal_nodes** option should contain a list of Hadoop Journal Nodes to collect metrics from.
- **name_nodes** option should contain a list of Hadoop Name Nodes to collect metrics from.
- **data_nodes** option should contain a list of Hadoop Data Nodes to collect metrics from.
- **journal_bean_names** option should contain a list of Bean Names to collect metrics from. This can not be set with collect_all_beans.
- **name_bean_names** option should contain a list of Bean Names to collect metrics from. This can not be set with collect_all_beans.
- **data_bean_names** option should contain a list of Bean Names to collect metrics from. This can not be set with collect_all_beans.
- **collect_all_beans** Default is false. This can not be set with the list of journal_bean_names, name_bean_names, and data_bean_names.

## Measurements & Fields

All metrics are stored in a few measurements:
- hadoop_journalnode
- hadoop_namenode
- hadoop_datanode


Tags:

- server: hadoop node address.
- bean_name: name of the bean of the collected data.
- type: The Node type - name, journal, and data.


Fields:

Hadoop can produce a lot of metrics, depending on the bean names you choose to collect.
Please refer to the hadoop Metrics documentation for further information.
This is not a comprehensive list of all the metrics that can be collected.

- hadoop_journalnode_metricssystem.NumActiveSinks
- hadoop_journalnode_metricssystem.SnapshotAvgTime
- hadoop_journalnode_metricssystem.PublishNumOps
- hadoop_journalnode_metricssystem.NumActiveSources
- hadoop_journalnode_metricssystem.NumAllSinks
- hadoop_journalnode_metricssystem.SnapshotNumOps
- hadoop_journalnode_metricssystem.PublishAvgTime
- hadoop_journalnode_metricssystem.DroppedPubAll
- hadoop_journalnode_metricssystem.NumAllSources
- hadoop_journalnode_rpcdetailed.GetJournalStateAvgTime
- hadoop_journalnode_rpcdetailed.NewEpochAvgTime
- hadoop_journalnode_rpcdetailed.JournalNumOps
- hadoop_journalnode_rpcdetailed.AcceptRecoveryNumOps
- hadoop_journalnode_rpcdetailed.StartLogSegmentAvgTime
- hadoop_journalnode_rpcdetailed.StartLogSegmentNumOps
- hadoop_journalnode_rpcdetailed.PrepareRecoveryNumOps
- hadoop_journalnode_rpcdetailed.IsFormattedAvgTime
- hadoop_journalnode_rpcdetailed.FinalizeLogSegmentAvgTime
- hadoop_journalnode_rpcdetailed.PrepareRecoveryAvgTime
- hadoop_journalnode_rpcdetailed.GetEditLogManifestNumOps
- hadoop_journalnode_rpcdetailed.GetJournalStateNumOps
- hadoop_journalnode_rpcdetailed.AcceptRecoveryAvgTime
- hadoop_journalnode_rpcdetailed.JournalAvgTime
- hadoop_journalnode_rpcdetailed.FormatAvgTime
- hadoop_journalnode_rpcdetailed.GetEditLogManifestAvgTime
- hadoop_journalnode_rpcdetailed.FormatNumOps
- hadoop_journalnode_rpcdetailed.NewEpochNumOps
- hadoop_journalnode_rpcdetailed.FinalizeLogSegmentNumOps
- hadoop_journalnode_rpcdetailed.IsFormattedNumOps
- hadoop_journalnode_rpc.RpcAuthorizationSuccesses
- hadoop_journalnode_rpc.NumOpenConnections
- hadoop_journalnode_rpc.SentBytes
- hadoop_journalnode_rpc.RpcAuthenticationSuccesses
- hadoop_journalnode_rpc.RpcQueueTimeAvgTime
- hadoop_journalnode_rpc.RpcProcessingTimeAvgTime
- hadoop_journalnode_rpc.RpcAuthorizationFailures
- hadoop_journalnode_rpc.RpcProcessingTimeNumOps
- hadoop_journalnode_rpc.RpcAuthenticationFailures
- hadoop_journalnode_rpc.CallQueueLength
- hadoop_journalnode_rpc.RpcQueueTimeNumOps
- hadoop_journalnode_rpc.ReceivedBytes
- hadoop_journalnode_Threading.DaemonThreadCount
- hadoop_journalnode_Threading.AllThreadIds_1
- hadoop_journalnode_Threading.AllThreadIds_9
- hadoop_journalnode_Threading.AllThreadIds_10
- hadoop_journalnode_Threading.AllThreadIds_0
- hadoop_journalnode_Threading.AllThreadIds_4
- hadoop_journalnode_Threading.PeakThreadCount=21
- hadoop_journalnode_Threading.ThreadCount
- hadoop_journalnode_Threading.TotalStartedThreadCount
- hadoop_journalnode_Threading.CurrentThreadUserTime
- hadoop_journalnode_Threading.CurrentThreadCpuTime
- hadoop_journalnode_OperatingSystem.OpenFileDescriptorCount
- hadoop_journalnode_OperatingSystem.AvailableProcessors
- hadoop_journalnode_OperatingSystem.FreeSwapSpaceSize
- hadoop_journalnode_OperatingSystem.ProcessCpuLoad
- hadoop_journalnode_OperatingSystem.TotalPhysicalMemorySize
- hadoop_journalnode_OperatingSystem.CommittedVirtualMemorySize
- hadoop_journalnode_OperatingSystem.FreePhysicalMemorySize
- hadoop_journalnode_OperatingSystem.SystemCpuLoad
- hadoop_journalnode_OperatingSystem.TotalSwapSpaceSize
- hadoop_journalnode_OperatingSystem.SystemLoadAverage
- hadoop_journalnode_OperatingSystem.MaxFileDescriptorCount
- hadoop_journalnode_OperatingSystem.ProcessCpuTime
- hadoop_journalnode_ugi.LoginSuccessNumOps
- hadoop_journalnode_ugi.LoginSuccessAvgTime
- hadoop_journalnode_ugi.LoginFailureAvgTime
- hadoop_journalnode_ugi.LoginFailureNumOps
- hadoop_journalnode_ugi.GetGroupsNumOps
- hadoop_journalnode_ugi.GetGroupsAvgTime
- hadoop_journalnode_jvm.ThreadsTerminated
- hadoop_journalnode_jvm.MemHeapUsedM
- hadoop_journalnode_jvm.ThreadsBlocked
- hadoop_journalnode_jvm.LogFatal
- hadoop_journalnode_jvm.LogError
- hadoop_journalnode_jvm.ThreadsRunnable
- hadoop_journalnode_jvm.ThreadsWaiting
- hadoop_journalnode_jvm.MemNonHeapCommittedM
- hadoop_journalnode_jvm.LogInfo
- hadoop_journalnode_jvm.GcTimeMillis
- hadoop_journalnode_jvm.GcCount
- hadoop_journalnode_jvm.ThreadsTimedWaiting
- hadoop_journalnode_jvm.LogWarn
- hadoop_journalnode_jvm.MemHeapCommittedM
- hadoop_journalnode_jvm.MemNonHeapMaxM
- hadoop_journalnode_jvm.MemNonHeapUsedM
- hadoop_journalnode_jvm.MemHeapMaxM
- hadoop_journalnode_jvm.MemMaxM
- hadoop_journalnode_jvm.ThreadsNew

- hadoop_namenode_FSNamesystemState.UnderReplicatedBlocks
- hadoop_namenode_FSNamesystemState.NumDecomDeadDataNodes
- hadoop_namenode_FSNamesystemState.NumDecommissioningDataNodes
- hadoop_namenode_FSNamesystemState.CapacityRemaining
- hadoop_namenode_FSNamesystemState.NumDeadDataNodes
- hadoop_namenode_FSNamesystemState.NumDecomLiveDataNodes
- hadoop_namenode_FSNamesystemState.VolumeFailuresTotal
- hadoop_namenode_FSNamesystemState.MaxObjects
- hadoop_namenode_FSNamesystemState.FilesTotal
- hadoop_namenode_FSNamesystemState.NumStaleDataNodes
- hadoop_namenode_FSNamesystemState.CapacityTotal
- hadoop_namenode_FSNamesystemState.ScheduledReplicationBlocks
- hadoop_namenode_FSNamesystemState.PendingDeletionBlocks
- hadoop_namenode_FSNamesystemState.EstimatedCapacityLostTotal
- hadoop_namenode_FSNamesystemState.CapacityUsed
- hadoop_namenode_FSNamesystemState.BlockDeletionStartTime
- hadoop_namenode_FSNamesystemState.NumStaleStorages
- hadoop_namenode_FSNamesystemState.NumLiveDataNodes
- hadoop_namenode_FSNamesystemState.PendingReplicationBlocks
- hadoop_namenode_FSNamesystemState.BlocksTotal
- hadoop_namenode_FSNamesystemState.TotalLoad
- hadoop_namenode_NameNodeInfo.NumberOfMissingBlocks
- hadoop_namenode_NameNodeInfo.BlockPoolUsedSpace
- hadoop_namenode_NameNodeInfo.DistinctVersionCount
- hadoop_namenode_NameNodeInfo.Threads
- hadoop_namenode_NameNodeInfo.Total
- hadoop_namenode_NameNodeInfo.PercentRemaining
- hadoop_namenode_NameNodeInfo.NumberOfMissingBlocksWithReplicationFactorOne
- hadoop_namenode_NameNodeInfo.Used
- hadoop_namenode_NameNodeInfo.PercentUsed
- hadoop_namenode_NameNodeInfo.CacheUsed
- hadoop_namenode_NameNodeInfo.Free
- hadoop_namenode_NameNodeInfo.CacheCapacity
- hadoop_namenode_NameNodeInfo.PercentBlockPoolUsed
- hadoop_namenode_NameNodeInfo.DistinctVersions_0_value
- hadoop_namenode_NameNodeInfo.TotalBlocks
- hadoop_namenode_NameNodeInfo.TotalFiles
- hadoop_namenode_NameNodeInfo.NonDfsUsedSpace
- hadoop_namenode_NameNodeStatus.LastHATransitionTime
- hadoop_namenode_dfs.UnderReplicatedBlocks
- hadoop_namenode_dfs.CapacityTotal
- hadoop_namenode_dfs.PendingDeletionBlocks
- hadoop_namenode_dfs.MissingBlocks
- hadoop_namenode_dfs.TotalFiles
- hadoop_namenode_dfs.Snapshots
- hadoop_namenode_dfs.MillisSinceLastLoadedEdits
- hadoop_namenode_dfs.ScheduledReplicationBlocks
- hadoop_namenode_dfs.LastWrittenTransactionId
- hadoop_namenode_dfs.ExcessBlocks
- hadoop_namenode_dfs.PendingDataNodeMessageCount
- hadoop_namenode_dfs.CorruptBlocks
- hadoop_namenode_dfs.BlocksTotal
- hadoop_namenode_dfs.CapacityUsedGB
- hadoop_namenode_dfs.CapacityRemaining
- hadoop_namenode_dfs.PendingReplicationBlocks
- hadoop_namenode_dfs.BlockCapacity
- hadoop_namenode_dfs.StaleDataNodes
- hadoop_namenode_dfs.MissingReplOneBlocks
- hadoop_namenode_dfs.TotalLoad
- hadoop_namenode_dfs.TransactionsSinceLastCheckpoint
- hadoop_namenode_dfs.CapacityUsed
- hadoop_namenode_dfs.PostponedMisreplicatedBlocks
- hadoop_namenode_dfs.TransactionsSinceLastLogRoll
- hadoop_namenode_dfs.FilesTotal
- hadoop_namenode_dfs.CapacityTotalGB
- hadoop_namenode_dfs.SnapshottableDirectories
- hadoop_namenode_dfs.LastCheckpointTime
- hadoop_namenode_dfs.CapacityRemainingGB
- hadoop_namenode_dfs.ExpiredHeartbeats
- hadoop_namenode_dfs.CapacityUsedNonDFS

- hadoop_datanode_FSDatasetState-null.Remaining
- hadoop_datanode_FSDatasetState-null.LastVolumeFailureDate
- hadoop_datanode_FSDatasetState-null.DfsUsed
- hadoop_datanode_FSDatasetState-null.CacheCapacity
- hadoop_datanode_FSDatasetState-null.CacheUsed
- hadoop_datanode_FSDatasetState-null.NumFailedVolumes
- hadoop_datanode_FSDatasetState-null.EstimatedCapacityLostTotal
- hadoop_datanode_FSDatasetState-null.Capacity
- hadoop_datanode_FSDatasetState-null.NumBlocksFailedToCache
- hadoop_datanode_FSDatasetState-null.NumBlocksFailedToUncache
- hadoop_datanode_FSDatasetState-null.NumBlocksCached


### Example Output:
```
$ telegraf -config ~/hadoop.conf -input-filter hadoop -test
* Plugin: inputs.hadoop, Collection 1
> hadoop_datanode_Threading,server=http://10.92.13.10:50075,type=data,bean_name=Threading AllThreadIds_67=24,AllThreadIds_1=2300,AllThreadIds_12=54,AllThreadIds_27=119,AllThreadIds_38=98,AllThreadIds_65=28,AllThreadIds_66=26,AllThreadIds_8=58,AllThreadIds_22=124,AllThreadIds_44=80,AllThreadIds_58=68,AllThreadIds_33=103,TotalStartedThreadCount=2291,AllThreadIds_14=52,AllThreadIds_26=120,AllThreadIds_28=118,AllThreadIds_32=104,AllThreadIds_61=32,AllThreadIds_62=31,AllThreadIds_5=61,AllThreadIds_21=125,AllThreadIds_45=79,AllThreadIds_57=70,AllThreadIds_64=29,AllThreadIds_15=51,AllThreadIds_24=122,AllThreadIds_29=117,AllThreadIds_30=106,AllThreadIds_48=76,AllThreadIds_53=66,CurrentThreadCpuTime=87279503699,AllThreadIds_2=619,AllThreadIds_3=63,AllThreadIds_7=59,AllThreadIds_40=96,AllThreadIds_39=97,AllThreadIds_70=2,AllThreadIds_51=73,AllThreadIds_13=53,AllThreadIds_34=102,AllThreadIds_47=77,AllThreadIds_50=74,AllThreadIds_23=123,AllThreadIds_41=82,AllThreadIds_56=71,AllThreadIds_49=75,AllThreadIds_52=72,AllThreadIds_18=48,AllThreadIds_25=121,AllThreadIds_31=105,AllThreadIds_46=78,AllThreadIds_54=69,CurrentThreadUserTime=62710000000,PeakThreadCount=88,AllThreadIds_6=60,AllThreadIds_10=56,AllThreadIds_35=101,AllThreadIds_37=99,AllThreadIds_42=83,AllThreadIds_63=30,AllThreadIds_68=4,AllThreadIds_71=1,AllThreadIds_4=62,AllThreadIds_20=126,AllThreadIds_36=100,AllThreadIds_43=81,AllThreadIds_55=25,AllThreadIds_19=139,AllThreadIds_60=65,ThreadCount=72,DaemonThreadCount=52,AllThreadIds_0=2316,AllThreadIds_11=55,AllThreadIds_16=50,AllThreadIds_17=49,AllThreadIds_9=57,AllThreadIds_59=67,AllThreadIds_69=3 1493134976000000000
> hadoop_datanode_OperatingSystem,server=http://10.92.13.10:50075,type=data,bean_name=OperatingSystem SystemCpuLoad=0.04716981132075472,MaxFileDescriptorCount=4096,SystemLoadAverage=0,FreeSwapSpaceSize=0,OpenFileDescriptorCount=331,ProcessCpuTime=261470000000,TotalPhysicalMemorySize=29503361024,TotalSwapSpaceSize=0,CommittedVirtualMemorySize=6230466560,FreePhysicalMemorySize=25873240064,ProcessCpuLoad=0.0015748031496062992,AvailableProcessors=8 1493134976000000000
> hadoop_datanode_dfs,bean_name=dfs,server=http://10.92.13.10:50075,type=data WritesFromLocalClient=0,BlocksCached=0,SendDataPacketBlockedOnNetworkNanosNumOps=0,WriteBlockOpNumOps=0,CopyBlockOpAvgTime=0,SendDataPacketBlockedOnNetworkNanosAvgTime=0,RamDiskBytesLazyPersisted=0,SendDataPacketTransferNanosAvgTime=0,BlockReportsAvgTime=3,RamDiskBytesWrite=0,RamDiskBlocksLazyPersistWindowMsNumOps=0,IncrementalBlockReportsNumOps=0,HeartbeatsAvgTime=1,FsyncCount=0,RamDiskBlocksEvicted=0,RamDiskBlocksEvictedWithoutRead=0,BlockChecksumOpNumOps=0,BlockReportsNumOps=9,BlocksRemoved=0,RamDiskBlocksLazyPersisted=0,ReadsFromRemoteClient=0,TotalWriteTime=0,ReadBlockOpAvgTime=0,RemoteBytesWritten=0,BlocksGetLocalPathInfo=0,BytesWritten=0,RamDiskBlocksEvictionWindowMsAvgTime=0,RemoteBytesRead=0,BlockVerificationFailures=0,BlockChecksumOpAvgTime=0,BytesRead=0,FsyncNanosAvgTime=0,FlushNanosNumOps=0,TotalReadTime=0,HeartbeatsNumOps=50228,RamDiskBlocksWriteFallback=0,CacheReportsAvgTime=0,ReplaceBlockOpAvgTime=0,BlocksVerified=0,RamDiskBlocksEvictionWindowMsNumOps=0,VolumeFailures=0,CacheReportsNumOps=0,DatanodeNetworkErrors=0,BlocksWritten=0,RamDiskBlocksWrite=0,FlushNanosAvgTime=0,PacketAckRoundTripTimeNanosNumOps=0,SendDataPacketTransferNanosNumOps=0,CopyBlockOpNumOps=0,IncrementalBlockReportsAvgTime=0,WritesFromRemoteClient=0,BlocksUncached=0,BlocksReplicated=0,FsyncNanosNumOps=0,RamDiskBlocksReadHits=0,ReadsFromLocalClient=0,WriteBlockOpAvgTime=0,BlocksRead=0,PacketAckRoundTripTimeNanosAvgTime=0,RamDiskBlocksDeletedBeforeLazyPersisted=0,ReadBlockOpNumOps=0,ReplaceBlockOpNumOps=0,RamDiskBlocksLazyPersistWindowMsAvgTime=0 1493134976000000000
> hadoop_datanode_jvm,type=data,bean_name=jvm,server=http://10.92.13.10:50075 MemNonHeapCommittedM=50.375,LogError=0,LogInfo=184,ThreadsBlocked=0,GcTimeMillis=792,GcTotalExtraSleepTime=261,ThreadsTimedWaiting=43,GcNumInfoThresholdExceeded=0,MemHeapUsedM=40.0319,GcCount=161,MemNonHeapMaxM=-0.0000009536743,LogWarn=0,GcNumWarnThresholdExceeded=0,ThreadsRunnable=25,LogFatal=0,MemHeapCommittedM=217.5,ThreadsWaiting=4,ThreadsNew=0,MemNonHeapUsedM=48.194725,ThreadsTerminated=0,MemHeapMaxM=3556,MemMaxM=3556 1493134976000000000
> hadoop_datanode_metricssystem,server=http://10.92.13.10:50075,type=data,bean_name=metricssystem NumAllSources=5,NumActiveSinks=0,NumAllSinks=0,NumActiveSources=5,PublishNumOps=0,SnapshotAvgTime=0,DroppedPubAll=0,SnapshotNumOps=0,PublishAvgTime=0 1493134976000000000
> hadoop_datanode_DataNodeInfo,bean_name=DataNodeInfo,server=http://10.92.13.10:50075,type=data XceiverCount=1 1493134976000000000
> hadoop_datanode_FSDatasetState-null,bean_name=FSDatasetState-null,server=http://10.92.13.10:50075,type=data LastVolumeFailureDate=0,NumBlocksFailedToCache=0,DfsUsed=327680,CacheUsed=0,CacheCapacity=0,NumBlocksFailedToUncache=0,Remaining=10260430602240,NumFailedVolumes=0,NumBlocksCached=0,EstimatedCapacityLostTotal=0,Capacity=10810613964800 1493134976000000000
> hadoop_datanode_ugi,server=http://10.92.13.10:50075,type=data,bean_name=ugi LoginFailureAvgTime=0,GetGroupsAvgTime=0,LoginSuccessNumOps=0,LoginFailureNumOps=0,GetGroupsNumOps=0,LoginSuccessAvgTime=0 1493134976000000000
> hadoop_datanode_rpc,server=http://10.92.13.10:50075,type=data,bean_name=rpc RpcAuthenticationFailures=0,RpcAuthorizationFailures=0,RpcQueueTimeNumOps=0,RpcProcessingTimeAvgTime=0,CallQueueLength=0,RpcQueueTimeAvgTime=0,NumOpenConnections=0,RpcAuthenticationSuccesses=0,RpcProcessingTimeNumOps=0,ReceivedBytes=0,SentBytes=0,RpcAuthorizationSuccesses=0 1493134976000000000
```

