// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A configuration for CloudWatch monitoring. You can configure your jobs to send
// log information to CloudWatch Logs.
type CloudWatchMonitoringConfiguration struct {

	// The name of the log group for log publishing.
	//
	// This member is required.
	LogGroupName *string

	// The specified name prefix for log streams.
	LogStreamNamePrefix *string

	noSmithyDocumentSerde
}

// A configuration specification to be used when provisioning virtual clusters,
// which can include configurations for applications and software bundled with
// Amazon EMR on EKS. A configuration consists of a classification, properties, and
// optional nested configurations. A classification refers to an
// application-specific configuration file. Properties are the settings you want to
// change in that file.
type Configuration struct {

	// The classification within a configuration.
	//
	// This member is required.
	Classification *string

	// A list of additional configurations to apply within a configuration object.
	Configurations []Configuration

	// A set of properties specified within a configuration classification.
	Properties map[string]string

	noSmithyDocumentSerde
}

// A configuration specification to be used to override existing configurations.
type ConfigurationOverrides struct {

	// The configurations for the application running by the job run.
	ApplicationConfiguration []Configuration

	// The configurations for monitoring.
	MonitoringConfiguration *MonitoringConfiguration

	noSmithyDocumentSerde
}

// The information about the container used for a job run or a managed endpoint.
//
// The following types satisfy this interface:
//  ContainerInfoMemberEksInfo
type ContainerInfo interface {
	isContainerInfo()
}

// The information about the EKS cluster.
type ContainerInfoMemberEksInfo struct {
	Value EksInfo

	noSmithyDocumentSerde
}

func (*ContainerInfoMemberEksInfo) isContainerInfo() {}

// The information about the container provider.
type ContainerProvider struct {

	// The ID of the container cluster.
	//
	// This member is required.
	Id *string

	// The type of the container provider. EKS is the only supported type as of now.
	//
	// This member is required.
	Type ContainerProviderType

	// The information about the container cluster.
	Info ContainerInfo

	noSmithyDocumentSerde
}

// The information about the EKS cluster.
type EksInfo struct {

	// The namespaces of the EKS cluster.
	Namespace *string

	noSmithyDocumentSerde
}

// This entity represents the endpoint that is managed by Amazon EMR on EKS.
type Endpoint struct {

	// The ARN of the endpoint.
	Arn *string

	// The certificate ARN of the endpoint.
	CertificateArn *string

	// The configuration settings that are used to override existing configurations for
	// endpoints.
	ConfigurationOverrides *ConfigurationOverrides

	// The date and time when the endpoint was created.
	CreatedAt *time.Time

	// The execution role ARN of the endpoint.
	ExecutionRoleArn *string

	// The reasons why the endpoint has failed.
	FailureReason FailureReason

	// The ID of the endpoint.
	Id *string

	// The name of the endpoint.
	Name *string

	// The EMR release version to be used for the endpoint.
	ReleaseLabel *string

	// The security group configuration of the endpoint.
	SecurityGroup *string

	// The server URL of the endpoint.
	ServerUrl *string

	// The state of the endpoint.
	State EndpointState

	// Additional details of the endpoint state.
	StateDetails *string

	// The subnet IDs of the endpoint.
	SubnetIds []string

	// The tags of the endpoint.
	Tags map[string]string

	// The type of the endpoint.
	Type *string

	// The ID of the endpoint's virtual cluster.
	VirtualClusterId *string

	noSmithyDocumentSerde
}

// Specify the driver that the job runs on.
type JobDriver struct {

	// The job driver parameters specified for spark submit.
	SparkSubmitJobDriver *SparkSubmitJobDriver

	noSmithyDocumentSerde
}

// This entity describes a job run. A job run is a unit of work, such as a Spark
// jar, PySpark script, or SparkSQL query, that you submit to Amazon EMR on EKS.
type JobRun struct {

	// The ARN of job run.
	Arn *string

	// The client token used to start a job run.
	ClientToken *string

	// The configuration settings that are used to override default configuration.
	ConfigurationOverrides *ConfigurationOverrides

	// The date and time when the job run was created.
	CreatedAt *time.Time

	// The user who created the job run.
	CreatedBy *string

	// The execution role ARN of the job run.
	ExecutionRoleArn *string

	// The reasons why the job run has failed.
	FailureReason FailureReason

	// The date and time when the job run has finished.
	FinishedAt *time.Time

	// The ID of the job run.
	Id *string

	// Parameters of job driver for the job run.
	JobDriver *JobDriver

	// The name of the job run.
	Name *string

	// The release version of Amazon EMR.
	ReleaseLabel *string

	// The state of the job run.
	State JobRunState

	// Additional details of the job run state.
	StateDetails *string

	// The assigned tags of the job run.
	Tags map[string]string

	// The ID of the job run's virtual cluster.
	VirtualClusterId *string

	noSmithyDocumentSerde
}

// Configuration setting for monitoring.
type MonitoringConfiguration struct {

	// Monitoring configurations for CloudWatch.
	CloudWatchMonitoringConfiguration *CloudWatchMonitoringConfiguration

	// Monitoring configurations for the persistent application UI.
	PersistentAppUI PersistentAppUI

	// Amazon S3 configuration for monitoring log publishing.
	S3MonitoringConfiguration *S3MonitoringConfiguration

	noSmithyDocumentSerde
}

// Amazon S3 configuration for monitoring log publishing. You can configure your
// jobs to send log information to Amazon S3.
type S3MonitoringConfiguration struct {

	// Amazon S3 destination URI for log publishing.
	//
	// This member is required.
	LogUri *string

	noSmithyDocumentSerde
}

// The information about job driver for Spark submit.
type SparkSubmitJobDriver struct {

	// The entry point of job application.
	//
	// This member is required.
	EntryPoint *string

	// The arguments for job application.
	EntryPointArguments []string

	// The Spark submit parameters that are used for job runs.
	SparkSubmitParameters *string

	noSmithyDocumentSerde
}

// This entity describes a virtual cluster. A virtual cluster is a Kubernetes
// namespace that Amazon EMR is registered with. Amazon EMR uses virtual clusters
// to run jobs and host endpoints. Multiple virtual clusters can be backed by the
// same physical cluster. However, each virtual cluster maps to one namespace on an
// EKS cluster. Virtual clusters do not create any active resources that contribute
// to your bill or that require lifecycle management outside the service.
type VirtualCluster struct {

	// The ARN of the virtual cluster.
	Arn *string

	// The container provider of the virtual cluster.
	ContainerProvider *ContainerProvider

	// The date and time when the virtual cluster is created.
	CreatedAt *time.Time

	// The ID of the virtual cluster.
	Id *string

	// The name of the virtual cluster.
	Name *string

	// The state of the virtual cluster.
	State VirtualClusterState

	// The assigned tags of the virtual cluster.
	Tags map[string]string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isContainerInfo() {}
