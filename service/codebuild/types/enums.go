// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ArtifactNamespace string

// Enum values for ArtifactNamespace
const (
	ArtifactNamespaceNone    ArtifactNamespace = "NONE"
	ArtifactNamespaceBuildId ArtifactNamespace = "BUILD_ID"
)

// Values returns all known values for ArtifactNamespace. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArtifactNamespace) Values() []ArtifactNamespace {
	return []ArtifactNamespace{
		"NONE",
		"BUILD_ID",
	}
}

type ArtifactPackaging string

// Enum values for ArtifactPackaging
const (
	ArtifactPackagingNone ArtifactPackaging = "NONE"
	ArtifactPackagingZip  ArtifactPackaging = "ZIP"
)

// Values returns all known values for ArtifactPackaging. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArtifactPackaging) Values() []ArtifactPackaging {
	return []ArtifactPackaging{
		"NONE",
		"ZIP",
	}
}

type ArtifactsType string

// Enum values for ArtifactsType
const (
	ArtifactsTypeCodepipeline ArtifactsType = "CODEPIPELINE"
	ArtifactsTypeS3           ArtifactsType = "S3"
	ArtifactsTypeNoArtifacts  ArtifactsType = "NO_ARTIFACTS"
)

// Values returns all known values for ArtifactsType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArtifactsType) Values() []ArtifactsType {
	return []ArtifactsType{
		"CODEPIPELINE",
		"S3",
		"NO_ARTIFACTS",
	}
}

type AuthType string

// Enum values for AuthType
const (
	AuthTypeOauth               AuthType = "OAUTH"
	AuthTypeBasicAuth           AuthType = "BASIC_AUTH"
	AuthTypePersonalAccessToken AuthType = "PERSONAL_ACCESS_TOKEN"
)

// Values returns all known values for AuthType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AuthType) Values() []AuthType {
	return []AuthType{
		"OAUTH",
		"BASIC_AUTH",
		"PERSONAL_ACCESS_TOKEN",
	}
}

type BucketOwnerAccess string

// Enum values for BucketOwnerAccess
const (
	BucketOwnerAccessNone     BucketOwnerAccess = "NONE"
	BucketOwnerAccessReadOnly BucketOwnerAccess = "READ_ONLY"
	BucketOwnerAccessFull     BucketOwnerAccess = "FULL"
)

// Values returns all known values for BucketOwnerAccess. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BucketOwnerAccess) Values() []BucketOwnerAccess {
	return []BucketOwnerAccess{
		"NONE",
		"READ_ONLY",
		"FULL",
	}
}

type BuildBatchPhaseType string

// Enum values for BuildBatchPhaseType
const (
	BuildBatchPhaseTypeSubmitted         BuildBatchPhaseType = "SUBMITTED"
	BuildBatchPhaseTypeDownloadBatchspec BuildBatchPhaseType = "DOWNLOAD_BATCHSPEC"
	BuildBatchPhaseTypeInProgress        BuildBatchPhaseType = "IN_PROGRESS"
	BuildBatchPhaseTypeCombineArtifacts  BuildBatchPhaseType = "COMBINE_ARTIFACTS"
	BuildBatchPhaseTypeSucceeded         BuildBatchPhaseType = "SUCCEEDED"
	BuildBatchPhaseTypeFailed            BuildBatchPhaseType = "FAILED"
	BuildBatchPhaseTypeStopped           BuildBatchPhaseType = "STOPPED"
)

// Values returns all known values for BuildBatchPhaseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BuildBatchPhaseType) Values() []BuildBatchPhaseType {
	return []BuildBatchPhaseType{
		"SUBMITTED",
		"DOWNLOAD_BATCHSPEC",
		"IN_PROGRESS",
		"COMBINE_ARTIFACTS",
		"SUCCEEDED",
		"FAILED",
		"STOPPED",
	}
}

type BuildPhaseType string

// Enum values for BuildPhaseType
const (
	BuildPhaseTypeSubmitted       BuildPhaseType = "SUBMITTED"
	BuildPhaseTypeQueued          BuildPhaseType = "QUEUED"
	BuildPhaseTypeProvisioning    BuildPhaseType = "PROVISIONING"
	BuildPhaseTypeDownloadSource  BuildPhaseType = "DOWNLOAD_SOURCE"
	BuildPhaseTypeInstall         BuildPhaseType = "INSTALL"
	BuildPhaseTypePreBuild        BuildPhaseType = "PRE_BUILD"
	BuildPhaseTypeBuild           BuildPhaseType = "BUILD"
	BuildPhaseTypePostBuild       BuildPhaseType = "POST_BUILD"
	BuildPhaseTypeUploadArtifacts BuildPhaseType = "UPLOAD_ARTIFACTS"
	BuildPhaseTypeFinalizing      BuildPhaseType = "FINALIZING"
	BuildPhaseTypeCompleted       BuildPhaseType = "COMPLETED"
)

// Values returns all known values for BuildPhaseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BuildPhaseType) Values() []BuildPhaseType {
	return []BuildPhaseType{
		"SUBMITTED",
		"QUEUED",
		"PROVISIONING",
		"DOWNLOAD_SOURCE",
		"INSTALL",
		"PRE_BUILD",
		"BUILD",
		"POST_BUILD",
		"UPLOAD_ARTIFACTS",
		"FINALIZING",
		"COMPLETED",
	}
}

type CacheMode string

// Enum values for CacheMode
const (
	CacheModeLocalDockerLayerCache CacheMode = "LOCAL_DOCKER_LAYER_CACHE"
	CacheModeLocalSourceCache      CacheMode = "LOCAL_SOURCE_CACHE"
	CacheModeLocalCustomCache      CacheMode = "LOCAL_CUSTOM_CACHE"
)

// Values returns all known values for CacheMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (CacheMode) Values() []CacheMode {
	return []CacheMode{
		"LOCAL_DOCKER_LAYER_CACHE",
		"LOCAL_SOURCE_CACHE",
		"LOCAL_CUSTOM_CACHE",
	}
}

type CacheType string

// Enum values for CacheType
const (
	CacheTypeNoCache CacheType = "NO_CACHE"
	CacheTypeS3      CacheType = "S3"
	CacheTypeLocal   CacheType = "LOCAL"
)

// Values returns all known values for CacheType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (CacheType) Values() []CacheType {
	return []CacheType{
		"NO_CACHE",
		"S3",
		"LOCAL",
	}
}

type ComputeType string

// Enum values for ComputeType
const (
	ComputeTypeBuildGeneral1Small   ComputeType = "BUILD_GENERAL1_SMALL"
	ComputeTypeBuildGeneral1Medium  ComputeType = "BUILD_GENERAL1_MEDIUM"
	ComputeTypeBuildGeneral1Large   ComputeType = "BUILD_GENERAL1_LARGE"
	ComputeTypeBuildGeneral12xlarge ComputeType = "BUILD_GENERAL1_2XLARGE"
)

// Values returns all known values for ComputeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ComputeType) Values() []ComputeType {
	return []ComputeType{
		"BUILD_GENERAL1_SMALL",
		"BUILD_GENERAL1_MEDIUM",
		"BUILD_GENERAL1_LARGE",
		"BUILD_GENERAL1_2XLARGE",
	}
}

type CredentialProviderType string

// Enum values for CredentialProviderType
const (
	CredentialProviderTypeSecretsManager CredentialProviderType = "SECRETS_MANAGER"
)

// Values returns all known values for CredentialProviderType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CredentialProviderType) Values() []CredentialProviderType {
	return []CredentialProviderType{
		"SECRETS_MANAGER",
	}
}

type EnvironmentType string

// Enum values for EnvironmentType
const (
	EnvironmentTypeWindowsContainer           EnvironmentType = "WINDOWS_CONTAINER"
	EnvironmentTypeLinuxContainer             EnvironmentType = "LINUX_CONTAINER"
	EnvironmentTypeLinuxGpuContainer          EnvironmentType = "LINUX_GPU_CONTAINER"
	EnvironmentTypeArmContainer               EnvironmentType = "ARM_CONTAINER"
	EnvironmentTypeWindowsServer2019Container EnvironmentType = "WINDOWS_SERVER_2019_CONTAINER"
)

// Values returns all known values for EnvironmentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentType) Values() []EnvironmentType {
	return []EnvironmentType{
		"WINDOWS_CONTAINER",
		"LINUX_CONTAINER",
		"LINUX_GPU_CONTAINER",
		"ARM_CONTAINER",
		"WINDOWS_SERVER_2019_CONTAINER",
	}
}

type EnvironmentVariableType string

// Enum values for EnvironmentVariableType
const (
	EnvironmentVariableTypePlaintext      EnvironmentVariableType = "PLAINTEXT"
	EnvironmentVariableTypeParameterStore EnvironmentVariableType = "PARAMETER_STORE"
	EnvironmentVariableTypeSecretsManager EnvironmentVariableType = "SECRETS_MANAGER"
)

// Values returns all known values for EnvironmentVariableType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentVariableType) Values() []EnvironmentVariableType {
	return []EnvironmentVariableType{
		"PLAINTEXT",
		"PARAMETER_STORE",
		"SECRETS_MANAGER",
	}
}

type FileSystemType string

// Enum values for FileSystemType
const (
	FileSystemTypeEfs FileSystemType = "EFS"
)

// Values returns all known values for FileSystemType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileSystemType) Values() []FileSystemType {
	return []FileSystemType{
		"EFS",
	}
}

type ImagePullCredentialsType string

// Enum values for ImagePullCredentialsType
const (
	ImagePullCredentialsTypeCodebuild   ImagePullCredentialsType = "CODEBUILD"
	ImagePullCredentialsTypeServiceRole ImagePullCredentialsType = "SERVICE_ROLE"
)

// Values returns all known values for ImagePullCredentialsType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImagePullCredentialsType) Values() []ImagePullCredentialsType {
	return []ImagePullCredentialsType{
		"CODEBUILD",
		"SERVICE_ROLE",
	}
}

type LanguageType string

// Enum values for LanguageType
const (
	LanguageTypeJava    LanguageType = "JAVA"
	LanguageTypePython  LanguageType = "PYTHON"
	LanguageTypeNodeJs  LanguageType = "NODE_JS"
	LanguageTypeRuby    LanguageType = "RUBY"
	LanguageTypeGolang  LanguageType = "GOLANG"
	LanguageTypeDocker  LanguageType = "DOCKER"
	LanguageTypeAndroid LanguageType = "ANDROID"
	LanguageTypeDotnet  LanguageType = "DOTNET"
	LanguageTypeBase    LanguageType = "BASE"
	LanguageTypePhp     LanguageType = "PHP"
)

// Values returns all known values for LanguageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LanguageType) Values() []LanguageType {
	return []LanguageType{
		"JAVA",
		"PYTHON",
		"NODE_JS",
		"RUBY",
		"GOLANG",
		"DOCKER",
		"ANDROID",
		"DOTNET",
		"BASE",
		"PHP",
	}
}

type LogsConfigStatusType string

// Enum values for LogsConfigStatusType
const (
	LogsConfigStatusTypeEnabled  LogsConfigStatusType = "ENABLED"
	LogsConfigStatusTypeDisabled LogsConfigStatusType = "DISABLED"
)

// Values returns all known values for LogsConfigStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LogsConfigStatusType) Values() []LogsConfigStatusType {
	return []LogsConfigStatusType{
		"ENABLED",
		"DISABLED",
	}
}

type PlatformType string

// Enum values for PlatformType
const (
	PlatformTypeDebian        PlatformType = "DEBIAN"
	PlatformTypeAmazonLinux   PlatformType = "AMAZON_LINUX"
	PlatformTypeUbuntu        PlatformType = "UBUNTU"
	PlatformTypeWindowsServer PlatformType = "WINDOWS_SERVER"
)

// Values returns all known values for PlatformType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PlatformType) Values() []PlatformType {
	return []PlatformType{
		"DEBIAN",
		"AMAZON_LINUX",
		"UBUNTU",
		"WINDOWS_SERVER",
	}
}

type ProjectSortByType string

// Enum values for ProjectSortByType
const (
	ProjectSortByTypeName             ProjectSortByType = "NAME"
	ProjectSortByTypeCreatedTime      ProjectSortByType = "CREATED_TIME"
	ProjectSortByTypeLastModifiedTime ProjectSortByType = "LAST_MODIFIED_TIME"
)

// Values returns all known values for ProjectSortByType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProjectSortByType) Values() []ProjectSortByType {
	return []ProjectSortByType{
		"NAME",
		"CREATED_TIME",
		"LAST_MODIFIED_TIME",
	}
}

type ReportCodeCoverageSortByType string

// Enum values for ReportCodeCoverageSortByType
const (
	ReportCodeCoverageSortByTypeLineCoveragePercentage ReportCodeCoverageSortByType = "LINE_COVERAGE_PERCENTAGE"
	ReportCodeCoverageSortByTypeFilePath               ReportCodeCoverageSortByType = "FILE_PATH"
)

// Values returns all known values for ReportCodeCoverageSortByType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReportCodeCoverageSortByType) Values() []ReportCodeCoverageSortByType {
	return []ReportCodeCoverageSortByType{
		"LINE_COVERAGE_PERCENTAGE",
		"FILE_PATH",
	}
}

type ReportExportConfigType string

// Enum values for ReportExportConfigType
const (
	ReportExportConfigTypeS3       ReportExportConfigType = "S3"
	ReportExportConfigTypeNoExport ReportExportConfigType = "NO_EXPORT"
)

// Values returns all known values for ReportExportConfigType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReportExportConfigType) Values() []ReportExportConfigType {
	return []ReportExportConfigType{
		"S3",
		"NO_EXPORT",
	}
}

type ReportGroupSortByType string

// Enum values for ReportGroupSortByType
const (
	ReportGroupSortByTypeName             ReportGroupSortByType = "NAME"
	ReportGroupSortByTypeCreatedTime      ReportGroupSortByType = "CREATED_TIME"
	ReportGroupSortByTypeLastModifiedTime ReportGroupSortByType = "LAST_MODIFIED_TIME"
)

// Values returns all known values for ReportGroupSortByType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReportGroupSortByType) Values() []ReportGroupSortByType {
	return []ReportGroupSortByType{
		"NAME",
		"CREATED_TIME",
		"LAST_MODIFIED_TIME",
	}
}

type ReportGroupStatusType string

// Enum values for ReportGroupStatusType
const (
	ReportGroupStatusTypeActive   ReportGroupStatusType = "ACTIVE"
	ReportGroupStatusTypeDeleting ReportGroupStatusType = "DELETING"
)

// Values returns all known values for ReportGroupStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReportGroupStatusType) Values() []ReportGroupStatusType {
	return []ReportGroupStatusType{
		"ACTIVE",
		"DELETING",
	}
}

type ReportGroupTrendFieldType string

// Enum values for ReportGroupTrendFieldType
const (
	ReportGroupTrendFieldTypePassRate        ReportGroupTrendFieldType = "PASS_RATE"
	ReportGroupTrendFieldTypeDuration        ReportGroupTrendFieldType = "DURATION"
	ReportGroupTrendFieldTypeTotal           ReportGroupTrendFieldType = "TOTAL"
	ReportGroupTrendFieldTypeLineCoverage    ReportGroupTrendFieldType = "LINE_COVERAGE"
	ReportGroupTrendFieldTypeLinesCovered    ReportGroupTrendFieldType = "LINES_COVERED"
	ReportGroupTrendFieldTypeLinesMissed     ReportGroupTrendFieldType = "LINES_MISSED"
	ReportGroupTrendFieldTypeBranchCoverage  ReportGroupTrendFieldType = "BRANCH_COVERAGE"
	ReportGroupTrendFieldTypeBranchesCovered ReportGroupTrendFieldType = "BRANCHES_COVERED"
	ReportGroupTrendFieldTypeBranchesMissed  ReportGroupTrendFieldType = "BRANCHES_MISSED"
)

// Values returns all known values for ReportGroupTrendFieldType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ReportGroupTrendFieldType) Values() []ReportGroupTrendFieldType {
	return []ReportGroupTrendFieldType{
		"PASS_RATE",
		"DURATION",
		"TOTAL",
		"LINE_COVERAGE",
		"LINES_COVERED",
		"LINES_MISSED",
		"BRANCH_COVERAGE",
		"BRANCHES_COVERED",
		"BRANCHES_MISSED",
	}
}

type ReportPackagingType string

// Enum values for ReportPackagingType
const (
	ReportPackagingTypeZip  ReportPackagingType = "ZIP"
	ReportPackagingTypeNone ReportPackagingType = "NONE"
)

// Values returns all known values for ReportPackagingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReportPackagingType) Values() []ReportPackagingType {
	return []ReportPackagingType{
		"ZIP",
		"NONE",
	}
}

type ReportStatusType string

// Enum values for ReportStatusType
const (
	ReportStatusTypeGenerating ReportStatusType = "GENERATING"
	ReportStatusTypeSucceeded  ReportStatusType = "SUCCEEDED"
	ReportStatusTypeFailed     ReportStatusType = "FAILED"
	ReportStatusTypeIncomplete ReportStatusType = "INCOMPLETE"
	ReportStatusTypeDeleting   ReportStatusType = "DELETING"
)

// Values returns all known values for ReportStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReportStatusType) Values() []ReportStatusType {
	return []ReportStatusType{
		"GENERATING",
		"SUCCEEDED",
		"FAILED",
		"INCOMPLETE",
		"DELETING",
	}
}

type ReportType string

// Enum values for ReportType
const (
	ReportTypeTest         ReportType = "TEST"
	ReportTypeCodeCoverage ReportType = "CODE_COVERAGE"
)

// Values returns all known values for ReportType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReportType) Values() []ReportType {
	return []ReportType{
		"TEST",
		"CODE_COVERAGE",
	}
}

type RetryBuildBatchType string

// Enum values for RetryBuildBatchType
const (
	RetryBuildBatchTypeRetryAllBuilds    RetryBuildBatchType = "RETRY_ALL_BUILDS"
	RetryBuildBatchTypeRetryFailedBuilds RetryBuildBatchType = "RETRY_FAILED_BUILDS"
)

// Values returns all known values for RetryBuildBatchType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RetryBuildBatchType) Values() []RetryBuildBatchType {
	return []RetryBuildBatchType{
		"RETRY_ALL_BUILDS",
		"RETRY_FAILED_BUILDS",
	}
}

type ServerType string

// Enum values for ServerType
const (
	ServerTypeGithub           ServerType = "GITHUB"
	ServerTypeBitbucket        ServerType = "BITBUCKET"
	ServerTypeGithubEnterprise ServerType = "GITHUB_ENTERPRISE"
)

// Values returns all known values for ServerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ServerType) Values() []ServerType {
	return []ServerType{
		"GITHUB",
		"BITBUCKET",
		"GITHUB_ENTERPRISE",
	}
}

type SharedResourceSortByType string

// Enum values for SharedResourceSortByType
const (
	SharedResourceSortByTypeArn          SharedResourceSortByType = "ARN"
	SharedResourceSortByTypeModifiedTime SharedResourceSortByType = "MODIFIED_TIME"
)

// Values returns all known values for SharedResourceSortByType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SharedResourceSortByType) Values() []SharedResourceSortByType {
	return []SharedResourceSortByType{
		"ARN",
		"MODIFIED_TIME",
	}
}

type SortOrderType string

// Enum values for SortOrderType
const (
	SortOrderTypeAscending  SortOrderType = "ASCENDING"
	SortOrderTypeDescending SortOrderType = "DESCENDING"
)

// Values returns all known values for SortOrderType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SortOrderType) Values() []SortOrderType {
	return []SortOrderType{
		"ASCENDING",
		"DESCENDING",
	}
}

type SourceAuthType string

// Enum values for SourceAuthType
const (
	SourceAuthTypeOauth SourceAuthType = "OAUTH"
)

// Values returns all known values for SourceAuthType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SourceAuthType) Values() []SourceAuthType {
	return []SourceAuthType{
		"OAUTH",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeCodecommit       SourceType = "CODECOMMIT"
	SourceTypeCodepipeline     SourceType = "CODEPIPELINE"
	SourceTypeGithub           SourceType = "GITHUB"
	SourceTypeS3               SourceType = "S3"
	SourceTypeBitbucket        SourceType = "BITBUCKET"
	SourceTypeGithubEnterprise SourceType = "GITHUB_ENTERPRISE"
	SourceTypeNoSource         SourceType = "NO_SOURCE"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"CODECOMMIT",
		"CODEPIPELINE",
		"GITHUB",
		"S3",
		"BITBUCKET",
		"GITHUB_ENTERPRISE",
		"NO_SOURCE",
	}
}

type StatusType string

// Enum values for StatusType
const (
	StatusTypeSucceeded  StatusType = "SUCCEEDED"
	StatusTypeFailed     StatusType = "FAILED"
	StatusTypeFault      StatusType = "FAULT"
	StatusTypeTimedOut   StatusType = "TIMED_OUT"
	StatusTypeInProgress StatusType = "IN_PROGRESS"
	StatusTypeStopped    StatusType = "STOPPED"
)

// Values returns all known values for StatusType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StatusType) Values() []StatusType {
	return []StatusType{
		"SUCCEEDED",
		"FAILED",
		"FAULT",
		"TIMED_OUT",
		"IN_PROGRESS",
		"STOPPED",
	}
}

type WebhookBuildType string

// Enum values for WebhookBuildType
const (
	WebhookBuildTypeBuild      WebhookBuildType = "BUILD"
	WebhookBuildTypeBuildBatch WebhookBuildType = "BUILD_BATCH"
)

// Values returns all known values for WebhookBuildType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WebhookBuildType) Values() []WebhookBuildType {
	return []WebhookBuildType{
		"BUILD",
		"BUILD_BATCH",
	}
}

type WebhookFilterType string

// Enum values for WebhookFilterType
const (
	WebhookFilterTypeEvent          WebhookFilterType = "EVENT"
	WebhookFilterTypeBaseRef        WebhookFilterType = "BASE_REF"
	WebhookFilterTypeHeadRef        WebhookFilterType = "HEAD_REF"
	WebhookFilterTypeActorAccountId WebhookFilterType = "ACTOR_ACCOUNT_ID"
	WebhookFilterTypeFilePath       WebhookFilterType = "FILE_PATH"
	WebhookFilterTypeCommitMessage  WebhookFilterType = "COMMIT_MESSAGE"
)

// Values returns all known values for WebhookFilterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WebhookFilterType) Values() []WebhookFilterType {
	return []WebhookFilterType{
		"EVENT",
		"BASE_REF",
		"HEAD_REF",
		"ACTOR_ACCOUNT_ID",
		"FILE_PATH",
		"COMMIT_MESSAGE",
	}
}
