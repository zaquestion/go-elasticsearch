// Code generated from specification version 8.0.0 (2f9f41ee7dc): DO NOT EDIT

package esapi

import (
	"github.com/elastic/go-elasticsearch/v8/esapi/xpack"
)

// API contains the Elasticsearch APIs
//
type API struct {
	Cat      *Cat
	Cluster  *Cluster
	Indices  *Indices
	Ingest   *Ingest
	Nodes    *Nodes
	Remote   *Remote
	Snapshot *Snapshot
	Tasks    *Tasks

	Bulk                    Bulk
	ClearScroll             ClearScroll
	Count                   Count
	Create                  Create
	Delete                  Delete
	DeleteByQuery           DeleteByQuery
	DeleteByQueryRethrottle DeleteByQueryRethrottle
	DeleteScript            DeleteScript
	Exists                  Exists
	ExistsSource            ExistsSource
	Explain                 Explain
	FieldCaps               FieldCaps
	Get                     Get
	GetScript               GetScript
	GetSource               GetSource
	Index                   Index
	Info                    Info
	Mget                    Mget
	Msearch                 Msearch
	MsearchTemplate         MsearchTemplate
	Mtermvectors            Mtermvectors
	Ping                    Ping
	PutScript               PutScript
	RankEval                RankEval
	Reindex                 Reindex
	ReindexRethrottle       ReindexRethrottle
	RenderSearchTemplate    RenderSearchTemplate
	ScriptsPainlessExecute  ScriptsPainlessExecute
	Scroll                  Scroll
	Search                  Search
	SearchShards            SearchShards
	SearchTemplate          SearchTemplate
	Termvectors             Termvectors
	Update                  Update
	UpdateByQuery           UpdateByQuery
	UpdateByQueryRethrottle UpdateByQueryRethrottle

	// X-Pack
	CCR        *CCR
	ILM        *ILM
	License    *License
	Migration  *Migration
	ML         *ML
	Monitoring *Monitoring
	Rollup     *Rollup
	Security   *Security
	SQL        *SQL
	SSL        *SSL
	Watcher    *Watcher
	XPack      *XPack

	DataFrameDeleteDataFrameTransform   xpack.DataFrameDeleteDataFrameTransform
	DataFrameGetDataFrameTransform      xpack.DataFrameGetDataFrameTransform
	DataFrameGetDataFrameTransformStats xpack.DataFrameGetDataFrameTransformStats
	DataFramePreviewDataFrameTransform  xpack.DataFramePreviewDataFrameTransform
	DataFramePutDataFrameTransform      xpack.DataFramePutDataFrameTransform
	DataFrameStartDataFrameTransform    xpack.DataFrameStartDataFrameTransform
	DataFrameStopDataFrameTransform     xpack.DataFrameStopDataFrameTransform
	GraphExplore                        xpack.GraphExplore
}

// Cat contains the Cat APIs
type Cat struct {
	Aliases      CatAliases
	Allocation   CatAllocation
	Count        CatCount
	Fielddata    CatFielddata
	Health       CatHealth
	Help         CatHelp
	Indices      CatIndices
	Master       CatMaster
	Nodeattrs    CatNodeattrs
	Nodes        CatNodes
	PendingTasks CatPendingTasks
	Plugins      CatPlugins
	Recovery     CatRecovery
	Repositories CatRepositories
	Segments     CatSegments
	Shards       CatShards
	Snapshots    CatSnapshots
	Tasks        CatTasks
	Templates    CatTemplates
	ThreadPool   CatThreadPool
}

// Cluster contains the Cluster APIs
type Cluster struct {
	AllocationExplain ClusterAllocationExplain
	GetSettings       ClusterGetSettings
	Health            ClusterHealth
	PendingTasks      ClusterPendingTasks
	PutSettings       ClusterPutSettings
	RemoteInfo        ClusterRemoteInfo
	Reroute           ClusterReroute
	State             ClusterState
	Stats             ClusterStats
}

// Indices contains the Indices APIs
type Indices struct {
	Analyze         IndicesAnalyze
	ClearCache      IndicesClearCache
	Close           IndicesClose
	Create          IndicesCreate
	Delete          IndicesDelete
	DeleteAlias     IndicesDeleteAlias
	DeleteTemplate  IndicesDeleteTemplate
	Exists          IndicesExists
	ExistsAlias     IndicesExistsAlias
	ExistsTemplate  IndicesExistsTemplate
	ExistsType      IndicesExistsType
	Flush           IndicesFlush
	FlushSynced     IndicesFlushSynced
	Forcemerge      IndicesForcemerge
	Get             IndicesGet
	GetAlias        IndicesGetAlias
	GetFieldMapping IndicesGetFieldMapping
	GetMapping      IndicesGetMapping
	GetSettings     IndicesGetSettings
	GetTemplate     IndicesGetTemplate
	GetUpgrade      IndicesGetUpgrade
	Open            IndicesOpen
	PutAlias        IndicesPutAlias
	PutMapping      IndicesPutMapping
	PutSettings     IndicesPutSettings
	PutTemplate     IndicesPutTemplate
	Recovery        IndicesRecovery
	Refresh         IndicesRefresh
	Rollover        IndicesRollover
	Segments        IndicesSegments
	ShardStores     IndicesShardStores
	Shrink          IndicesShrink
	Split           IndicesSplit
	Stats           IndicesStats
	UpdateAliases   IndicesUpdateAliases
	Upgrade         IndicesUpgrade
	ValidateQuery   IndicesValidateQuery

	Freeze   xpack.IndicesFreeze
	Unfreeze xpack.IndicesUnfreeze
}

// Ingest contains the Ingest APIs
type Ingest struct {
	DeletePipeline IngestDeletePipeline
	GetPipeline    IngestGetPipeline
	ProcessorGrok  IngestProcessorGrok
	PutPipeline    IngestPutPipeline
	Simulate       IngestSimulate
}

// Nodes contains the Nodes APIs
type Nodes struct {
	HotThreads           NodesHotThreads
	Info                 NodesInfo
	ReloadSecureSettings NodesReloadSecureSettings
	Stats                NodesStats
	Usage                NodesUsage
}

// Remote contains the Remote APIs
type Remote struct {
}

// Snapshot contains the Snapshot APIs
type Snapshot struct {
	Create           SnapshotCreate
	CreateRepository SnapshotCreateRepository
	Delete           SnapshotDelete
	DeleteRepository SnapshotDeleteRepository
	Get              SnapshotGet
	GetRepository    SnapshotGetRepository
	Restore          SnapshotRestore
	Status           SnapshotStatus
	VerifyRepository SnapshotVerifyRepository
}

// Tasks contains the Tasks APIs
type Tasks struct {
	Cancel TasksCancel
	Get    TasksGet
	List   TasksList
}

// CCR contains the CCR APIs
type CCR struct {
	DeleteAutoFollowPattern xpack.CCRDeleteAutoFollowPattern
	Follow                  xpack.CCRFollow
	FollowInfo              xpack.CCRFollowInfo
	FollowStats             xpack.CCRFollowStats
	ForgetFollower          xpack.CCRForgetFollower
	GetAutoFollowPattern    xpack.CCRGetAutoFollowPattern
	PauseFollow             xpack.CCRPauseFollow
	PutAutoFollowPattern    xpack.CCRPutAutoFollowPattern
	ResumeFollow            xpack.CCRResumeFollow
	Stats                   xpack.CCRStats
	Unfollow                xpack.CCRUnfollow
}

// ILM contains the ILM APIs
type ILM struct {
	DeleteLifecycle  xpack.ILMDeleteLifecycle
	ExplainLifecycle xpack.ILMExplainLifecycle
	GetLifecycle     xpack.ILMGetLifecycle
	GetStatus        xpack.ILMGetStatus
	MoveToStep       xpack.ILMMoveToStep
	PutLifecycle     xpack.ILMPutLifecycle
	RemovePolicy     xpack.ILMRemovePolicy
	Retry            xpack.ILMRetry
	Start            xpack.ILMStart
	Stop             xpack.ILMStop
}

// License contains the License APIs
type License struct {
	Delete         xpack.LicenseDelete
	Get            xpack.LicenseGet
	GetBasicStatus xpack.LicenseGetBasicStatus
	GetTrialStatus xpack.LicenseGetTrialStatus
	Post           xpack.LicensePost
	PostStartBasic xpack.LicensePostStartBasic
	PostStartTrial xpack.LicensePostStartTrial
}

// Migration contains the Migration APIs
type Migration struct {
	Deprecations xpack.MigrationDeprecations
}

// ML contains the ML APIs
type ML struct {
	CloseJob            xpack.MLCloseJob
	DeleteCalendar      xpack.MLDeleteCalendar
	DeleteCalendarEvent xpack.MLDeleteCalendarEvent
	DeleteCalendarJob   xpack.MLDeleteCalendarJob
	DeleteDatafeed      xpack.MLDeleteDatafeed
	DeleteExpiredData   xpack.MLDeleteExpiredData
	DeleteFilter        xpack.MLDeleteFilter
	DeleteForecast      xpack.MLDeleteForecast
	DeleteJob           xpack.MLDeleteJob
	DeleteModelSnapshot xpack.MLDeleteModelSnapshot
	FindFileStructure   xpack.MLFindFileStructure
	FlushJob            xpack.MLFlushJob
	Forecast            xpack.MLForecast
	GetBuckets          xpack.MLGetBuckets
	GetCalendarEvents   xpack.MLGetCalendarEvents
	GetCalendars        xpack.MLGetCalendars
	GetCategories       xpack.MLGetCategories
	GetDatafeedStats    xpack.MLGetDatafeedStats
	GetDatafeeds        xpack.MLGetDatafeeds
	GetFilters          xpack.MLGetFilters
	GetInfluencers      xpack.MLGetInfluencers
	GetJobStats         xpack.MLGetJobStats
	GetJobs             xpack.MLGetJobs
	GetModelSnapshots   xpack.MLGetModelSnapshots
	GetOverallBuckets   xpack.MLGetOverallBuckets
	GetRecords          xpack.MLGetRecords
	Info                xpack.MLInfo
	OpenJob             xpack.MLOpenJob
	PostCalendarEvents  xpack.MLPostCalendarEvents
	PostData            xpack.MLPostData
	PreviewDatafeed     xpack.MLPreviewDatafeed
	PutCalendar         xpack.MLPutCalendar
	PutCalendarJob      xpack.MLPutCalendarJob
	PutDatafeed         xpack.MLPutDatafeed
	PutFilter           xpack.MLPutFilter
	PutJob              xpack.MLPutJob
	RevertModelSnapshot xpack.MLRevertModelSnapshot
	SetUpgradeMode      xpack.MLSetUpgradeMode
	StartDatafeed       xpack.MLStartDatafeed
	StopDatafeed        xpack.MLStopDatafeed
	UpdateDatafeed      xpack.MLUpdateDatafeed
	UpdateFilter        xpack.MLUpdateFilter
	UpdateJob           xpack.MLUpdateJob
	UpdateModelSnapshot xpack.MLUpdateModelSnapshot
	Validate            xpack.MLValidate
	ValidateDetector    xpack.MLValidateDetector
}

// Monitoring contains the Monitoring APIs
type Monitoring struct {
	Bulk xpack.MonitoringBulk
}

// Rollup contains the Rollup APIs
type Rollup struct {
	DeleteJob          xpack.RollupDeleteJob
	GetJobs            xpack.RollupGetJobs
	GetRollupCaps      xpack.RollupGetRollupCaps
	GetRollupIndexCaps xpack.RollupGetRollupIndexCaps
	PutJob             xpack.RollupPutJob
	RollupSearch       xpack.RollupRollupSearch
	StartJob           xpack.RollupStartJob
	StopJob            xpack.RollupStopJob
}

// Security contains the Security APIs
type Security struct {
	Authenticate      xpack.SecurityAuthenticate
	ChangePassword    xpack.SecurityChangePassword
	ClearCachedRealms xpack.SecurityClearCachedRealms
	ClearCachedRoles  xpack.SecurityClearCachedRoles
	CreateApiKey      xpack.SecurityCreateApiKey
	DeletePrivileges  xpack.SecurityDeletePrivileges
	DeleteRole        xpack.SecurityDeleteRole
	DeleteRoleMapping xpack.SecurityDeleteRoleMapping
	DeleteUser        xpack.SecurityDeleteUser
	DisableUser       xpack.SecurityDisableUser
	EnableUser        xpack.SecurityEnableUser
	GetApiKey         xpack.SecurityGetApiKey
	GetPrivileges     xpack.SecurityGetPrivileges
	GetRole           xpack.SecurityGetRole
	GetRoleMapping    xpack.SecurityGetRoleMapping
	GetToken          xpack.SecurityGetToken
	GetUser           xpack.SecurityGetUser
	GetUserPrivileges xpack.SecurityGetUserPrivileges
	HasPrivileges     xpack.SecurityHasPrivileges
	InvalidateApiKey  xpack.SecurityInvalidateApiKey
	InvalidateToken   xpack.SecurityInvalidateToken
	PutPrivileges     xpack.SecurityPutPrivileges
	PutRole           xpack.SecurityPutRole
	PutRoleMapping    xpack.SecurityPutRoleMapping
	PutUser           xpack.SecurityPutUser
}

// SQL contains the SQL APIs
type SQL struct {
	ClearCursor xpack.SQLClearCursor
	Query       xpack.SQLQuery
	Translate   xpack.SQLTranslate
}

// SSL contains the SSL APIs
type SSL struct {
	Certificates xpack.SSLCertificates
}

// Watcher contains the Watcher APIs
type Watcher struct {
	AckWatch        xpack.WatcherAckWatch
	ActivateWatch   xpack.WatcherActivateWatch
	DeactivateWatch xpack.WatcherDeactivateWatch
	DeleteWatch     xpack.WatcherDeleteWatch
	ExecuteWatch    xpack.WatcherExecuteWatch
	GetWatch        xpack.WatcherGetWatch
	PutWatch        xpack.WatcherPutWatch
	Start           xpack.WatcherStart
	Stats           xpack.WatcherStats
	Stop            xpack.WatcherStop
}

// XPack contains the XPack APIs
type XPack struct {
	Info  xpack.XPackInfo
	Usage xpack.XPackUsage
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
		Bulk:                    newBulkFunc(t),
		ClearScroll:             newClearScrollFunc(t),
		Count:                   newCountFunc(t),
		Create:                  newCreateFunc(t),
		Delete:                  newDeleteFunc(t),
		DeleteByQuery:           newDeleteByQueryFunc(t),
		DeleteByQueryRethrottle: newDeleteByQueryRethrottleFunc(t),
		DeleteScript:            newDeleteScriptFunc(t),
		Exists:                  newExistsFunc(t),
		ExistsSource:            newExistsSourceFunc(t),
		Explain:                 newExplainFunc(t),
		FieldCaps:               newFieldCapsFunc(t),
		Get:                     newGetFunc(t),
		GetScript:               newGetScriptFunc(t),
		GetSource:               newGetSourceFunc(t),
		Index:                   newIndexFunc(t),
		Info:                    newInfoFunc(t),
		Mget:                    newMgetFunc(t),
		Msearch:                 newMsearchFunc(t),
		MsearchTemplate:         newMsearchTemplateFunc(t),
		Mtermvectors:            newMtermvectorsFunc(t),
		Ping:                    newPingFunc(t),
		PutScript:               newPutScriptFunc(t),
		RankEval:                newRankEvalFunc(t),
		Reindex:                 newReindexFunc(t),
		ReindexRethrottle:       newReindexRethrottleFunc(t),
		RenderSearchTemplate:    newRenderSearchTemplateFunc(t),
		ScriptsPainlessExecute:  newScriptsPainlessExecuteFunc(t),
		Scroll:                  newScrollFunc(t),
		Search:                  newSearchFunc(t),
		SearchShards:            newSearchShardsFunc(t),
		SearchTemplate:          newSearchTemplateFunc(t),
		Termvectors:             newTermvectorsFunc(t),
		Update:                  newUpdateFunc(t),
		UpdateByQuery:           newUpdateByQueryFunc(t),
		UpdateByQueryRethrottle: newUpdateByQueryRethrottleFunc(t),
		Cat: &Cat{
			Aliases:      newCatAliasesFunc(t),
			Allocation:   newCatAllocationFunc(t),
			Count:        newCatCountFunc(t),
			Fielddata:    newCatFielddataFunc(t),
			Health:       newCatHealthFunc(t),
			Help:         newCatHelpFunc(t),
			Indices:      newCatIndicesFunc(t),
			Master:       newCatMasterFunc(t),
			Nodeattrs:    newCatNodeattrsFunc(t),
			Nodes:        newCatNodesFunc(t),
			PendingTasks: newCatPendingTasksFunc(t),
			Plugins:      newCatPluginsFunc(t),
			Recovery:     newCatRecoveryFunc(t),
			Repositories: newCatRepositoriesFunc(t),
			Segments:     newCatSegmentsFunc(t),
			Shards:       newCatShardsFunc(t),
			Snapshots:    newCatSnapshotsFunc(t),
			Tasks:        newCatTasksFunc(t),
			Templates:    newCatTemplatesFunc(t),
			ThreadPool:   newCatThreadPoolFunc(t),
		},
		Cluster: &Cluster{
			AllocationExplain: newClusterAllocationExplainFunc(t),
			GetSettings:       newClusterGetSettingsFunc(t),
			Health:            newClusterHealthFunc(t),
			PendingTasks:      newClusterPendingTasksFunc(t),
			PutSettings:       newClusterPutSettingsFunc(t),
			RemoteInfo:        newClusterRemoteInfoFunc(t),
			Reroute:           newClusterRerouteFunc(t),
			State:             newClusterStateFunc(t),
			Stats:             newClusterStatsFunc(t),
		},
		Indices: &Indices{
			Analyze:         newIndicesAnalyzeFunc(t),
			ClearCache:      newIndicesClearCacheFunc(t),
			Close:           newIndicesCloseFunc(t),
			Create:          newIndicesCreateFunc(t),
			Delete:          newIndicesDeleteFunc(t),
			DeleteAlias:     newIndicesDeleteAliasFunc(t),
			DeleteTemplate:  newIndicesDeleteTemplateFunc(t),
			Exists:          newIndicesExistsFunc(t),
			ExistsAlias:     newIndicesExistsAliasFunc(t),
			ExistsTemplate:  newIndicesExistsTemplateFunc(t),
			ExistsType:      newIndicesExistsTypeFunc(t),
			Flush:           newIndicesFlushFunc(t),
			FlushSynced:     newIndicesFlushSyncedFunc(t),
			Forcemerge:      newIndicesForcemergeFunc(t),
			Get:             newIndicesGetFunc(t),
			GetAlias:        newIndicesGetAliasFunc(t),
			GetFieldMapping: newIndicesGetFieldMappingFunc(t),
			GetMapping:      newIndicesGetMappingFunc(t),
			GetSettings:     newIndicesGetSettingsFunc(t),
			GetTemplate:     newIndicesGetTemplateFunc(t),
			GetUpgrade:      newIndicesGetUpgradeFunc(t),
			Open:            newIndicesOpenFunc(t),
			PutAlias:        newIndicesPutAliasFunc(t),
			PutMapping:      newIndicesPutMappingFunc(t),
			PutSettings:     newIndicesPutSettingsFunc(t),
			PutTemplate:     newIndicesPutTemplateFunc(t),
			Recovery:        newIndicesRecoveryFunc(t),
			Refresh:         newIndicesRefreshFunc(t),
			Rollover:        newIndicesRolloverFunc(t),
			Segments:        newIndicesSegmentsFunc(t),
			ShardStores:     newIndicesShardStoresFunc(t),
			Shrink:          newIndicesShrinkFunc(t),
			Split:           newIndicesSplitFunc(t),
			Stats:           newIndicesStatsFunc(t),
			UpdateAliases:   newIndicesUpdateAliasesFunc(t),
			Upgrade:         newIndicesUpgradeFunc(t),
			ValidateQuery:   newIndicesValidateQueryFunc(t),

			Freeze:   xpack.NewIndicesFreezeFunc(t),
			Unfreeze: xpack.NewIndicesUnfreezeFunc(t),
		},
		Ingest: &Ingest{
			DeletePipeline: newIngestDeletePipelineFunc(t),
			GetPipeline:    newIngestGetPipelineFunc(t),
			ProcessorGrok:  newIngestProcessorGrokFunc(t),
			PutPipeline:    newIngestPutPipelineFunc(t),
			Simulate:       newIngestSimulateFunc(t),
		},
		Nodes: &Nodes{
			HotThreads:           newNodesHotThreadsFunc(t),
			Info:                 newNodesInfoFunc(t),
			ReloadSecureSettings: newNodesReloadSecureSettingsFunc(t),
			Stats:                newNodesStatsFunc(t),
			Usage:                newNodesUsageFunc(t),
		},
		Remote: &Remote{},
		Snapshot: &Snapshot{
			Create:           newSnapshotCreateFunc(t),
			CreateRepository: newSnapshotCreateRepositoryFunc(t),
			Delete:           newSnapshotDeleteFunc(t),
			DeleteRepository: newSnapshotDeleteRepositoryFunc(t),
			Get:              newSnapshotGetFunc(t),
			GetRepository:    newSnapshotGetRepositoryFunc(t),
			Restore:          newSnapshotRestoreFunc(t),
			Status:           newSnapshotStatusFunc(t),
			VerifyRepository: newSnapshotVerifyRepositoryFunc(t),
		},
		Tasks: &Tasks{
			Cancel: newTasksCancelFunc(t),
			Get:    newTasksGetFunc(t),
			List:   newTasksListFunc(t),
		},

		License: &License{
			Delete:         xpack.NewLicenseDeleteFunc(t),
			Get:            xpack.NewLicenseGetFunc(t),
			GetBasicStatus: xpack.NewLicenseGetBasicStatusFunc(t),
			GetTrialStatus: xpack.NewLicenseGetTrialStatusFunc(t),
			Post:           xpack.NewLicensePostFunc(t),
			PostStartBasic: xpack.NewLicensePostStartBasicFunc(t),
			PostStartTrial: xpack.NewLicensePostStartTrialFunc(t),
		},
		XPack: &XPack{
			Info:  xpack.NewXPackInfoFunc(t),
			Usage: xpack.NewXPackUsageFunc(t),
		},
	}
}
