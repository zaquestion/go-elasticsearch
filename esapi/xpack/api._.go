// Code generated from specification version 8-0-0-SNAPSHOT (e343ed1633b): DO NOT EDIT

package xpack

// XAPI contains the Elasticsearch APIs
//
type XAPI struct {
	CCR        *CCR
	ILM        *ILM
	Indices    *Indices
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

	DataFrameDeleteDataFrameTransform   DataFrameDeleteDataFrameTransform
	DataFrameGetDataFrameTransform      DataFrameGetDataFrameTransform
	DataFrameGetDataFrameTransformStats DataFrameGetDataFrameTransformStats
	DataFramePreviewDataFrameTransform  DataFramePreviewDataFrameTransform
	DataFramePutDataFrameTransform      DataFramePutDataFrameTransform
	DataFrameStartDataFrameTransform    DataFrameStartDataFrameTransform
	DataFrameStopDataFrameTransform     DataFrameStopDataFrameTransform
	GraphExplore                        GraphExplore
}

// CCR contains the CCR APIs
type CCR struct {
	DeleteAutoFollowPattern CCRDeleteAutoFollowPattern
	Follow                  CCRFollow
	FollowInfo              CCRFollowInfo
	FollowStats             CCRFollowStats
	ForgetFollower          CCRForgetFollower
	GetAutoFollowPattern    CCRGetAutoFollowPattern
	PauseFollow             CCRPauseFollow
	PutAutoFollowPattern    CCRPutAutoFollowPattern
	ResumeFollow            CCRResumeFollow
	Stats                   CCRStats
	Unfollow                CCRUnfollow
}

// ILM contains the ILM APIs
type ILM struct {
	DeleteLifecycle  ILMDeleteLifecycle
	ExplainLifecycle ILMExplainLifecycle
	GetLifecycle     ILMGetLifecycle
	GetStatus        ILMGetStatus
	MoveToStep       ILMMoveToStep
	PutLifecycle     ILMPutLifecycle
	RemovePolicy     ILMRemovePolicy
	Retry            ILMRetry
	Start            ILMStart
	Stop             ILMStop
}

// Indices contains the Indices APIs
type Indices struct {
	Freeze   IndicesFreeze
	Unfreeze IndicesUnfreeze
}

// License contains the License APIs
type License struct {
	Delete         LicenseDelete
	Get            LicenseGet
	GetBasicStatus LicenseGetBasicStatus
	GetTrialStatus LicenseGetTrialStatus
	Post           LicensePost
	PostStartBasic LicensePostStartBasic
	PostStartTrial LicensePostStartTrial
}

// Migration contains the Migration APIs
type Migration struct {
	Deprecations MigrationDeprecations
}

// ML contains the ML APIs
type ML struct {
	CloseJob            MLCloseJob
	DeleteCalendar      MLDeleteCalendar
	DeleteCalendarEvent MLDeleteCalendarEvent
	DeleteCalendarJob   MLDeleteCalendarJob
	DeleteDatafeed      MLDeleteDatafeed
	DeleteExpiredData   MLDeleteExpiredData
	DeleteFilter        MLDeleteFilter
	DeleteForecast      MLDeleteForecast
	DeleteJob           MLDeleteJob
	DeleteModelSnapshot MLDeleteModelSnapshot
	FindFileStructure   MLFindFileStructure
	FlushJob            MLFlushJob
	Forecast            MLForecast
	GetBuckets          MLGetBuckets
	GetCalendarEvents   MLGetCalendarEvents
	GetCalendars        MLGetCalendars
	GetCategories       MLGetCategories
	GetDatafeedStats    MLGetDatafeedStats
	GetDatafeeds        MLGetDatafeeds
	GetFilters          MLGetFilters
	GetInfluencers      MLGetInfluencers
	GetJobStats         MLGetJobStats
	GetJobs             MLGetJobs
	GetModelSnapshots   MLGetModelSnapshots
	GetOverallBuckets   MLGetOverallBuckets
	GetRecords          MLGetRecords
	Info                MLInfo
	OpenJob             MLOpenJob
	PostCalendarEvents  MLPostCalendarEvents
	PostData            MLPostData
	PreviewDatafeed     MLPreviewDatafeed
	PutCalendar         MLPutCalendar
	PutCalendarJob      MLPutCalendarJob
	PutDatafeed         MLPutDatafeed
	PutFilter           MLPutFilter
	PutJob              MLPutJob
	RevertModelSnapshot MLRevertModelSnapshot
	SetUpgradeMode      MLSetUpgradeMode
	StartDatafeed       MLStartDatafeed
	StopDatafeed        MLStopDatafeed
	UpdateDatafeed      MLUpdateDatafeed
	UpdateFilter        MLUpdateFilter
	UpdateJob           MLUpdateJob
	UpdateModelSnapshot MLUpdateModelSnapshot
	Validate            MLValidate
	ValidateDetector    MLValidateDetector
}

// Monitoring contains the Monitoring APIs
type Monitoring struct {
	Bulk MonitoringBulk
}

// Rollup contains the Rollup APIs
type Rollup struct {
	DeleteJob          RollupDeleteJob
	GetJobs            RollupGetJobs
	GetRollupCaps      RollupGetRollupCaps
	GetRollupIndexCaps RollupGetRollupIndexCaps
	PutJob             RollupPutJob
	RollupSearch       RollupRollupSearch
	StartJob           RollupStartJob
	StopJob            RollupStopJob
}

// Security contains the Security APIs
type Security struct {
	Authenticate      SecurityAuthenticate
	ChangePassword    SecurityChangePassword
	ClearCachedRealms SecurityClearCachedRealms
	ClearCachedRoles  SecurityClearCachedRoles
	CreateApiKey      SecurityCreateApiKey
	DeletePrivileges  SecurityDeletePrivileges
	DeleteRole        SecurityDeleteRole
	DeleteRoleMapping SecurityDeleteRoleMapping
	DeleteUser        SecurityDeleteUser
	DisableUser       SecurityDisableUser
	EnableUser        SecurityEnableUser
	GetApiKey         SecurityGetApiKey
	GetPrivileges     SecurityGetPrivileges
	GetRole           SecurityGetRole
	GetRoleMapping    SecurityGetRoleMapping
	GetToken          SecurityGetToken
	GetUser           SecurityGetUser
	GetUserPrivileges SecurityGetUserPrivileges
	HasPrivileges     SecurityHasPrivileges
	InvalidateApiKey  SecurityInvalidateApiKey
	InvalidateToken   SecurityInvalidateToken
	PutPrivileges     SecurityPutPrivileges
	PutRole           SecurityPutRole
	PutRoleMapping    SecurityPutRoleMapping
	PutUser           SecurityPutUser
}

// SQL contains the SQL APIs
type SQL struct {
	ClearCursor SQLClearCursor
	Query       SQLQuery
	Translate   SQLTranslate
}

// SSL contains the SSL APIs
type SSL struct {
	Certificates SSLCertificates
}

// Watcher contains the Watcher APIs
type Watcher struct {
	AckWatch        WatcherAckWatch
	ActivateWatch   WatcherActivateWatch
	DeactivateWatch WatcherDeactivateWatch
	DeleteWatch     WatcherDeleteWatch
	ExecuteWatch    WatcherExecuteWatch
	GetWatch        WatcherGetWatch
	PutWatch        WatcherPutWatch
	Start           WatcherStart
	Stats           WatcherStats
	Stop            WatcherStop
}

// XPack contains the XPack APIs
type XPack struct {
	Info  XPackInfo
	Usage XPackUsage
}

// New creates new XAPI
//
func New(t Transport) *XAPI {
	return &XAPI{
		DataFrameDeleteDataFrameTransform:   newDataFrameDeleteDataFrameTransformFunc(t),
		DataFrameGetDataFrameTransform:      newDataFrameGetDataFrameTransformFunc(t),
		DataFrameGetDataFrameTransformStats: newDataFrameGetDataFrameTransformStatsFunc(t),
		DataFramePreviewDataFrameTransform:  newDataFramePreviewDataFrameTransformFunc(t),
		DataFramePutDataFrameTransform:      newDataFramePutDataFrameTransformFunc(t),
		DataFrameStartDataFrameTransform:    newDataFrameStartDataFrameTransformFunc(t),
		DataFrameStopDataFrameTransform:     newDataFrameStopDataFrameTransformFunc(t),
		GraphExplore:                        newGraphExploreFunc(t),
		CCR: &CCR{
			DeleteAutoFollowPattern: newCCRDeleteAutoFollowPatternFunc(t),
			Follow:                  newCCRFollowFunc(t),
			FollowInfo:              newCCRFollowInfoFunc(t),
			FollowStats:             newCCRFollowStatsFunc(t),
			ForgetFollower:          newCCRForgetFollowerFunc(t),
			GetAutoFollowPattern:    newCCRGetAutoFollowPatternFunc(t),
			PauseFollow:             newCCRPauseFollowFunc(t),
			PutAutoFollowPattern:    newCCRPutAutoFollowPatternFunc(t),
			ResumeFollow:            newCCRResumeFollowFunc(t),
			Stats:                   newCCRStatsFunc(t),
			Unfollow:                newCCRUnfollowFunc(t),
		},
		ILM: &ILM{
			DeleteLifecycle:  newILMDeleteLifecycleFunc(t),
			ExplainLifecycle: newILMExplainLifecycleFunc(t),
			GetLifecycle:     newILMGetLifecycleFunc(t),
			GetStatus:        newILMGetStatusFunc(t),
			MoveToStep:       newILMMoveToStepFunc(t),
			PutLifecycle:     newILMPutLifecycleFunc(t),
			RemovePolicy:     newILMRemovePolicyFunc(t),
			Retry:            newILMRetryFunc(t),
			Start:            newILMStartFunc(t),
			Stop:             newILMStopFunc(t),
		},
		Indices: &Indices{
			Freeze:   newIndicesFreezeFunc(t),
			Unfreeze: newIndicesUnfreezeFunc(t),
		},
		License: &License{
			Delete:         newLicenseDeleteFunc(t),
			Get:            newLicenseGetFunc(t),
			GetBasicStatus: newLicenseGetBasicStatusFunc(t),
			GetTrialStatus: newLicenseGetTrialStatusFunc(t),
			Post:           newLicensePostFunc(t),
			PostStartBasic: newLicensePostStartBasicFunc(t),
			PostStartTrial: newLicensePostStartTrialFunc(t),
		},
		Migration: &Migration{
			Deprecations: newMigrationDeprecationsFunc(t),
		},
		ML: &ML{
			CloseJob:            newMLCloseJobFunc(t),
			DeleteCalendar:      newMLDeleteCalendarFunc(t),
			DeleteCalendarEvent: newMLDeleteCalendarEventFunc(t),
			DeleteCalendarJob:   newMLDeleteCalendarJobFunc(t),
			DeleteDatafeed:      newMLDeleteDatafeedFunc(t),
			DeleteExpiredData:   newMLDeleteExpiredDataFunc(t),
			DeleteFilter:        newMLDeleteFilterFunc(t),
			DeleteForecast:      newMLDeleteForecastFunc(t),
			DeleteJob:           newMLDeleteJobFunc(t),
			DeleteModelSnapshot: newMLDeleteModelSnapshotFunc(t),
			FindFileStructure:   newMLFindFileStructureFunc(t),
			FlushJob:            newMLFlushJobFunc(t),
			Forecast:            newMLForecastFunc(t),
			GetBuckets:          newMLGetBucketsFunc(t),
			GetCalendarEvents:   newMLGetCalendarEventsFunc(t),
			GetCalendars:        newMLGetCalendarsFunc(t),
			GetCategories:       newMLGetCategoriesFunc(t),
			GetDatafeedStats:    newMLGetDatafeedStatsFunc(t),
			GetDatafeeds:        newMLGetDatafeedsFunc(t),
			GetFilters:          newMLGetFiltersFunc(t),
			GetInfluencers:      newMLGetInfluencersFunc(t),
			GetJobStats:         newMLGetJobStatsFunc(t),
			GetJobs:             newMLGetJobsFunc(t),
			GetModelSnapshots:   newMLGetModelSnapshotsFunc(t),
			GetOverallBuckets:   newMLGetOverallBucketsFunc(t),
			GetRecords:          newMLGetRecordsFunc(t),
			Info:                newMLInfoFunc(t),
			OpenJob:             newMLOpenJobFunc(t),
			PostCalendarEvents:  newMLPostCalendarEventsFunc(t),
			PostData:            newMLPostDataFunc(t),
			PreviewDatafeed:     newMLPreviewDatafeedFunc(t),
			PutCalendar:         newMLPutCalendarFunc(t),
			PutCalendarJob:      newMLPutCalendarJobFunc(t),
			PutDatafeed:         newMLPutDatafeedFunc(t),
			PutFilter:           newMLPutFilterFunc(t),
			PutJob:              newMLPutJobFunc(t),
			RevertModelSnapshot: newMLRevertModelSnapshotFunc(t),
			SetUpgradeMode:      newMLSetUpgradeModeFunc(t),
			StartDatafeed:       newMLStartDatafeedFunc(t),
			StopDatafeed:        newMLStopDatafeedFunc(t),
			UpdateDatafeed:      newMLUpdateDatafeedFunc(t),
			UpdateFilter:        newMLUpdateFilterFunc(t),
			UpdateJob:           newMLUpdateJobFunc(t),
			UpdateModelSnapshot: newMLUpdateModelSnapshotFunc(t),
			Validate:            newMLValidateFunc(t),
			ValidateDetector:    newMLValidateDetectorFunc(t),
		},
		Monitoring: &Monitoring{
			Bulk: newMonitoringBulkFunc(t),
		},
		Rollup: &Rollup{
			DeleteJob:          newRollupDeleteJobFunc(t),
			GetJobs:            newRollupGetJobsFunc(t),
			GetRollupCaps:      newRollupGetRollupCapsFunc(t),
			GetRollupIndexCaps: newRollupGetRollupIndexCapsFunc(t),
			PutJob:             newRollupPutJobFunc(t),
			RollupSearch:       newRollupRollupSearchFunc(t),
			StartJob:           newRollupStartJobFunc(t),
			StopJob:            newRollupStopJobFunc(t),
		},
		Security: &Security{
			Authenticate:      newSecurityAuthenticateFunc(t),
			ChangePassword:    newSecurityChangePasswordFunc(t),
			ClearCachedRealms: newSecurityClearCachedRealmsFunc(t),
			ClearCachedRoles:  newSecurityClearCachedRolesFunc(t),
			CreateApiKey:      newSecurityCreateApiKeyFunc(t),
			DeletePrivileges:  newSecurityDeletePrivilegesFunc(t),
			DeleteRole:        newSecurityDeleteRoleFunc(t),
			DeleteRoleMapping: newSecurityDeleteRoleMappingFunc(t),
			DeleteUser:        newSecurityDeleteUserFunc(t),
			DisableUser:       newSecurityDisableUserFunc(t),
			EnableUser:        newSecurityEnableUserFunc(t),
			GetApiKey:         newSecurityGetApiKeyFunc(t),
			GetPrivileges:     newSecurityGetPrivilegesFunc(t),
			GetRole:           newSecurityGetRoleFunc(t),
			GetRoleMapping:    newSecurityGetRoleMappingFunc(t),
			GetToken:          newSecurityGetTokenFunc(t),
			GetUser:           newSecurityGetUserFunc(t),
			GetUserPrivileges: newSecurityGetUserPrivilegesFunc(t),
			HasPrivileges:     newSecurityHasPrivilegesFunc(t),
			InvalidateApiKey:  newSecurityInvalidateApiKeyFunc(t),
			InvalidateToken:   newSecurityInvalidateTokenFunc(t),
			PutPrivileges:     newSecurityPutPrivilegesFunc(t),
			PutRole:           newSecurityPutRoleFunc(t),
			PutRoleMapping:    newSecurityPutRoleMappingFunc(t),
			PutUser:           newSecurityPutUserFunc(t),
		},
		SQL: &SQL{
			ClearCursor: newSQLClearCursorFunc(t),
			Query:       newSQLQueryFunc(t),
			Translate:   newSQLTranslateFunc(t),
		},
		SSL: &SSL{
			Certificates: newSSLCertificatesFunc(t),
		},
		Watcher: &Watcher{
			AckWatch:        newWatcherAckWatchFunc(t),
			ActivateWatch:   newWatcherActivateWatchFunc(t),
			DeactivateWatch: newWatcherDeactivateWatchFunc(t),
			DeleteWatch:     newWatcherDeleteWatchFunc(t),
			ExecuteWatch:    newWatcherExecuteWatchFunc(t),
			GetWatch:        newWatcherGetWatchFunc(t),
			PutWatch:        newWatcherPutWatchFunc(t),
			Start:           newWatcherStartFunc(t),
			Stats:           newWatcherStatsFunc(t),
			Stop:            newWatcherStopFunc(t),
		},
		XPack: &XPack{
			Info:  newXPackInfoFunc(t),
			Usage: newXPackUsageFunc(t),
		},
	}
}
