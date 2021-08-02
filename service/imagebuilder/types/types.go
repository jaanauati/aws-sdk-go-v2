// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// In addition to your infrastruction configuration, these settings provide an
// extra layer of control over your build instances. For instances where Image
// Builder installs the SSM agent, you can choose whether to keep it for the AMI
// that you create. You can also specify commands to run on launch for all of your
// build instances.
type AdditionalInstanceConfiguration struct {

	// Contains settings for the SSM agent on your build instance.
	SystemsManagerAgent *SystemsManagerAgent

	// Use this property to provide commands or a command script to run when you launch
	// your build instance. The userDataOverride property replaces any commands that
	// Image Builder might have added to ensure that SSM is installed on your Linux
	// build instance. If you override the user data, make sure that you add commands
	// to install SSM, if it is not pre-installed on your source image.
	UserDataOverride *string

	noSmithyDocumentSerde
}

// Details of an Amazon EC2 AMI.
type Ami struct {

	// The account ID of the owner of the AMI.
	AccountId *string

	// The description of the Amazon EC2 AMI. Minimum and maximum length are in
	// characters.
	Description *string

	// The AMI ID of the Amazon EC2 AMI.
	Image *string

	// The name of the Amazon EC2 AMI.
	Name *string

	// The Region of the Amazon EC2 AMI.
	Region *string

	// Image state shows the image status and the reason for that status.
	State *ImageState

	noSmithyDocumentSerde
}

// Define and configure the output AMIs of the pipeline.
type AmiDistributionConfiguration struct {

	// The tags to apply to AMIs distributed to this Region.
	AmiTags map[string]string

	// The description of the distribution configuration. Minimum and maximum length
	// are in characters.
	Description *string

	// The KMS key identifier used to encrypt the distributed image.
	KmsKeyId *string

	// Launch permissions can be used to configure which accounts can use the AMI to
	// launch instances.
	LaunchPermission *LaunchPermissionConfiguration

	// The name of the distribution configuration.
	Name *string

	// The ID of an account to which you want to distribute an image.
	TargetAccountIds []string

	noSmithyDocumentSerde
}

// A detailed view of a component.
type Component struct {

	// The Amazon Resource Name (ARN) of the component.
	Arn *string

	// The change description of the component.
	ChangeDescription *string

	// The data of the component.
	Data *string

	// The date that the component was created.
	DateCreated *string

	// The description of the component.
	Description *string

	// The encryption status of the component.
	Encrypted *bool

	// The KMS key identifier used to encrypt the component.
	KmsKeyId *string

	// The name of the component.
	Name *string

	// The owner of the component.
	Owner *string

	// Contains parameter details for each of the parameters that are defined for the
	// component.
	Parameters []ComponentParameterDetail

	// The platform of the component.
	Platform Platform

	// The operating system (OS) version supported by the component. If the OS
	// information is available, a prefix match is performed against the parent image
	// OS version during image recipe creation.
	SupportedOsVersions []string

	// The tags associated with the component.
	Tags map[string]string

	// The type of the component denotes whether the component is used to build the
	// image or only to test it.
	Type ComponentType

	// The version of the component.
	Version *string

	noSmithyDocumentSerde
}

// Configuration details of the component.
type ComponentConfiguration struct {

	// The Amazon Resource Name (ARN) of the component.
	//
	// This member is required.
	ComponentArn *string

	// A group of parameter settings that are used to configure the component for a
	// specific recipe.
	Parameters []ComponentParameter

	noSmithyDocumentSerde
}

// Contains a key/value pair that sets the named component parameter.
type ComponentParameter struct {

	// The name of the component parameter to set.
	//
	// This member is required.
	Name *string

	// Sets the value for the named component parameter.
	//
	// This member is required.
	Value []string

	noSmithyDocumentSerde
}

// Defines a parameter that is used to provide configuration details for the
// component.
type ComponentParameterDetail struct {

	// The name of this input parameter.
	//
	// This member is required.
	Name *string

	// The type of input this parameter provides. The currently supported value is
	// "string".
	//
	// This member is required.
	Type *string

	// The default value of this parameter if no input is provided.
	DefaultValue []string

	// Describes this parameter.
	Description *string

	noSmithyDocumentSerde
}

// A high-level summary of a component.
type ComponentSummary struct {

	// The Amazon Resource Name (ARN) of the component.
	Arn *string

	// The change description of the component.
	ChangeDescription *string

	// The date that the component was created.
	DateCreated *string

	// The description of the component.
	Description *string

	// The name of the component.
	Name *string

	// The owner of the component.
	Owner *string

	// The platform of the component.
	Platform Platform

	// The operating system (OS) version supported by the component. If the OS
	// information is available, a prefix match is performed against the parent image
	// OS version during image recipe creation.
	SupportedOsVersions []string

	// The tags associated with the component.
	Tags map[string]string

	// The type of the component denotes whether the component is used to build the
	// image or only to test it.
	Type ComponentType

	// The version of the component.
	Version *string

	noSmithyDocumentSerde
}

// The defining characteristics of a specific version of an TOE component.
type ComponentVersion struct {

	// The Amazon Resource Name (ARN) of the component. Semantic versioning is included
	// in each object's Amazon Resource Name (ARN), at the level that applies to that
	// object as follows:
	//
	// * Versionless ARNs and Name ARNs do not include specific
	// values in any of the nodes. The nodes are either left off entirely, or they are
	// specified as wildcards, for example: x.x.x.
	//
	// * Version ARNs have only the first
	// three nodes: ..
	//
	// * Build version ARNs have all four nodes, and point to a
	// specific build for a specific version of an object.
	Arn *string

	// The date that the component was created.
	DateCreated *string

	// The description of the component.
	Description *string

	// The name of the component.
	Name *string

	// The owner of the component.
	Owner *string

	// The platform of the component.
	Platform Platform

	// he operating system (OS) version supported by the component. If the OS
	// information is available, a prefix match is performed against the parent image
	// OS version during image recipe creation.
	SupportedOsVersions []string

	// The type of the component denotes whether the component is used to build the
	// image or only to test it.
	Type ComponentType

	// The semantic version of the component. The semantic version has four nodes: ../.
	// You can assign values for the first three, and can filter on all of them.
	// Assignment: For the first three nodes you can assign any positive integer value,
	// including zero, with an upper limit of 2^30-1, or 1073741823 for each node.
	// Image Builder automatically assigns the build number, and that is not open for
	// updates. Patterns: You can use any numeric pattern that adheres to the
	// assignment requirements for the nodes that you can assign. For example, you
	// might choose a software version pattern, such as 1.0.0, or a date, such as
	// 2021.01.01. Filtering: When you retrieve or reference a resource with a semantic
	// version, you can use wildcards (x) to filter your results. When you use a
	// wildcard in any node, all nodes to the right of the first wildcard must also be
	// wildcards. For example, specifying "1.2.x", or "1.x.x" works to filter list
	// results, but neither "1.x.2", nor "x.2.x" will work. You do not have to specify
	// the build - Image Builder automatically uses a wildcard for that, if applicable.
	Version *string

	noSmithyDocumentSerde
}

// A container encapsulates the runtime environment for an application.
type Container struct {

	// A list of URIs for containers created in the context Region.
	ImageUris []string

	// Containers and container images are Region-specific. This is the Region context
	// for the container.
	Region *string

	noSmithyDocumentSerde
}

// Container distribution settings for encryption, licensing, and sharing in a
// specific Region.
type ContainerDistributionConfiguration struct {

	// The destination repository for the container distribution configuration.
	//
	// This member is required.
	TargetRepository *TargetContainerRepository

	// Tags that are attached to the container distribution configuration.
	ContainerTags []string

	// The description of the container distribution configuration.
	Description *string

	noSmithyDocumentSerde
}

// A container recipe.
type ContainerRecipe struct {

	// The Amazon Resource Name (ARN) of the container recipe. Semantic versioning is
	// included in each object's Amazon Resource Name (ARN), at the level that applies
	// to that object as follows:
	//
	// * Versionless ARNs and Name ARNs do not include
	// specific values in any of the nodes. The nodes are either left off entirely, or
	// they are specified as wildcards, for example: x.x.x.
	//
	// * Version ARNs have only
	// the first three nodes: ..
	//
	// * Build version ARNs have all four nodes, and point
	// to a specific build for a specific version of an object.
	Arn *string

	// Components for build and test that are included in the container recipe.
	Components []ComponentConfiguration

	// Specifies the type of container, such as Docker.
	ContainerType ContainerType

	// The date when this container recipe was created.
	DateCreated *string

	// The description of the container recipe.
	Description *string

	// Dockerfiles are text documents that are used to build Docker containers, and
	// ensure that they contain all of the elements required by the application running
	// inside. The template data consists of contextual variables where Image Builder
	// places build information or scripts, based on your container image recipe.
	DockerfileTemplateData *string

	// A flag that indicates if the target container is encrypted.
	Encrypted *bool

	// A group of options that can be used to configure an instance for building and
	// testing container images.
	InstanceConfiguration *InstanceConfiguration

	// Identifies which KMS key is used to encrypt the container image for distribution
	// to the target Region.
	KmsKeyId *string

	// The name of the container recipe.
	Name *string

	// The owner of the container recipe.
	Owner *string

	// The source image for the container recipe.
	ParentImage *string

	// The system platform for the container, such as Windows or Linux.
	Platform Platform

	// Tags that are attached to the container recipe.
	Tags map[string]string

	// The destination repository for the container image.
	TargetRepository *TargetContainerRepository

	// The semantic version of the container recipe. The semantic version has four
	// nodes: ../. You can assign values for the first three, and can filter on all of
	// them. Assignment: For the first three nodes you can assign any positive integer
	// value, including zero, with an upper limit of 2^30-1, or 1073741823 for each
	// node. Image Builder automatically assigns the build number, and that is not open
	// for updates. Patterns: You can use any numeric pattern that adheres to the
	// assignment requirements for the nodes that you can assign. For example, you
	// might choose a software version pattern, such as 1.0.0, or a date, such as
	// 2021.01.01. Filtering: When you retrieve or reference a resource with a semantic
	// version, you can use wildcards (x) to filter your results. When you use a
	// wildcard in any node, all nodes to the right of the first wildcard must also be
	// wildcards. For example, specifying "1.2.x", or "1.x.x" works to filter list
	// results, but neither "1.x.2", nor "x.2.x" will work. You do not have to specify
	// the build - Image Builder automatically uses a wildcard for that, if applicable.
	Version *string

	// The working directory for use during build and test workflows.
	WorkingDirectory *string

	noSmithyDocumentSerde
}

// A summary of a container recipe
type ContainerRecipeSummary struct {

	// The Amazon Resource Name (ARN) of the container recipe.
	Arn *string

	// Specifies the type of container, such as "Docker".
	ContainerType ContainerType

	// The date when this container recipe was created.
	DateCreated *string

	// The name of the container recipe.
	Name *string

	// The owner of the container recipe.
	Owner *string

	// The source image for the container recipe.
	ParentImage *string

	// The system platform for the container, such as Windows or Linux.
	Platform Platform

	// Tags that are attached to the container recipe.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Defines the settings for a specific Region.
type Distribution struct {

	// The target Region.
	//
	// This member is required.
	Region *string

	// The specific AMI settings; for example, launch permissions or AMI tags.
	AmiDistributionConfiguration *AmiDistributionConfiguration

	// Container distribution settings for encryption, licensing, and sharing in a
	// specific Region.
	ContainerDistributionConfiguration *ContainerDistributionConfiguration

	// A group of launchTemplateConfiguration settings that apply to image distribution
	// for specified accounts.
	LaunchTemplateConfigurations []LaunchTemplateConfiguration

	// The License Manager Configuration to associate with the AMI in the specified
	// Region.
	LicenseConfigurationArns []string

	noSmithyDocumentSerde
}

// A distribution configuration.
type DistributionConfiguration struct {

	// The maximum duration in minutes for this distribution configuration.
	//
	// This member is required.
	TimeoutMinutes *int32

	// The Amazon Resource Name (ARN) of the distribution configuration.
	Arn *string

	// The date on which this distribution configuration was created.
	DateCreated *string

	// The date on which this distribution configuration was last updated.
	DateUpdated *string

	// The description of the distribution configuration.
	Description *string

	// The distribution objects that apply Region-specific settings for the deployment
	// of the image to targeted Regions.
	Distributions []Distribution

	// The name of the distribution configuration.
	Name *string

	// The tags of the distribution configuration.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A high-level overview of a distribution configuration.
type DistributionConfigurationSummary struct {

	// The Amazon Resource Name (ARN) of the distribution configuration.
	Arn *string

	// The date on which the distribution configuration was created.
	DateCreated *string

	// The date on which the distribution configuration was updated.
	DateUpdated *string

	// The description of the distribution configuration.
	Description *string

	// The name of the distribution configuration.
	Name *string

	// A list of Regions where the container image is distributed to.
	Regions []string

	// The tags associated with the distribution configuration.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Amazon EBS-specific block device mapping specifications.
type EbsInstanceBlockDeviceSpecification struct {

	// Use to configure delete on termination of the associated device.
	DeleteOnTermination *bool

	// Use to configure device encryption.
	Encrypted *bool

	// Use to configure device IOPS.
	Iops *int32

	// Use to configure the KMS key to use when encrypting the device.
	KmsKeyId *string

	// The snapshot that defines the device contents.
	SnapshotId *string

	// Use to override the device's volume size.
	VolumeSize *int32

	// Use to override the device's volume type.
	VolumeType EbsVolumeType

	noSmithyDocumentSerde
}

// A filter name and value pair that is used to return a more specific list of
// results from a list operation. Filters can be used to match a set of resources
// by specific criteria, such as tags, attributes, or IDs.
type Filter struct {

	// The name of the filter. Filter names are case-sensitive.
	Name *string

	// The filter values. Filter values are case-sensitive.
	Values []string

	noSmithyDocumentSerde
}

// An Image Builder image. You must specify exactly one recipe for the image –
// either a container recipe (containerRecipe), which creates a container image, or
// an image recipe (imageRecipe), which creates an AMI.
type Image struct {

	// The Amazon Resource Name (ARN) of the image. Semantic versioning is included in
	// each object's Amazon Resource Name (ARN), at the level that applies to that
	// object as follows:
	//
	// * Versionless ARNs and Name ARNs do not include specific
	// values in any of the nodes. The nodes are either left off entirely, or they are
	// specified as wildcards, for example: x.x.x.
	//
	// * Version ARNs have only the first
	// three nodes: ..
	//
	// * Build version ARNs have all four nodes, and point to a
	// specific build for a specific version of an object.
	Arn *string

	// The recipe that is used to create an Image Builder container image.
	ContainerRecipe *ContainerRecipe

	// The date on which this image was created.
	DateCreated *string

	// The distribution configuration used when creating this image.
	DistributionConfiguration *DistributionConfiguration

	// Collects additional information about the image being created, including the
	// operating system (OS) version and package list. This information is used to
	// enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled *bool

	// The image recipe used when creating the image.
	ImageRecipe *ImageRecipe

	// The image tests configuration used when creating this image.
	ImageTestsConfiguration *ImageTestsConfiguration

	// The infrastructure used when creating this image.
	InfrastructureConfiguration *InfrastructureConfiguration

	// The name of the image.
	Name *string

	// The operating system version of the instance. For example, Amazon Linux 2,
	// Ubuntu 18, or Microsoft Windows Server 2019.
	OsVersion *string

	// The output resources produced when creating this image.
	OutputResources *OutputResources

	// The platform of the image.
	Platform Platform

	// The Amazon Resource Name (ARN) of the image pipeline that created this image.
	SourcePipelineArn *string

	// The name of the image pipeline that created this image.
	SourcePipelineName *string

	// The state of the image.
	State *ImageState

	// The tags of the image.
	Tags map[string]string

	// Specifies whether this is an AMI or container image.
	Type ImageType

	// The semantic version of the image. The semantic version has four nodes: ../. You
	// can assign values for the first three, and can filter on all of them.
	// Assignment: For the first three nodes you can assign any positive integer value,
	// including zero, with an upper limit of 2^30-1, or 1073741823 for each node.
	// Image Builder automatically assigns the build number, and that is not open for
	// updates. Patterns: You can use any numeric pattern that adheres to the
	// assignment requirements for the nodes that you can assign. For example, you
	// might choose a software version pattern, such as 1.0.0, or a date, such as
	// 2021.01.01. Filtering: When you retrieve or reference a resource with a semantic
	// version, you can use wildcards (x) to filter your results. When you use a
	// wildcard in any node, all nodes to the right of the first wildcard must also be
	// wildcards. For example, specifying "1.2.x", or "1.x.x" works to filter list
	// results, but neither "1.x.2", nor "x.2.x" will work. You do not have to specify
	// the build - Image Builder automatically uses a wildcard for that, if applicable.
	Version *string

	noSmithyDocumentSerde
}

// Represents a package installed on an Image Builder image.
type ImagePackage struct {

	// The name of the package as reported to the operating system package manager.
	PackageName *string

	// The version of the package as reported to the operating system package manager.
	PackageVersion *string

	noSmithyDocumentSerde
}

// Details of an image pipeline.
type ImagePipeline struct {

	// The Amazon Resource Name (ARN) of the image pipeline.
	Arn *string

	// The Amazon Resource Name (ARN) of the container recipe that is used for this
	// pipeline.
	ContainerRecipeArn *string

	// The date on which this image pipeline was created.
	DateCreated *string

	// The date on which this image pipeline was last run.
	DateLastRun *string

	// The date on which this image pipeline will next be run.
	DateNextRun *string

	// The date on which this image pipeline was last updated.
	DateUpdated *string

	// The description of the image pipeline.
	Description *string

	// The Amazon Resource Name (ARN) of the distribution configuration associated with
	// this image pipeline.
	DistributionConfigurationArn *string

	// Collects additional information about the image being created, including the
	// operating system (OS) version and package list. This information is used to
	// enhance the overall experience of using EC2 Image Builder. Enabled by default.
	EnhancedImageMetadataEnabled *bool

	// The Amazon Resource Name (ARN) of the image recipe associated with this image
	// pipeline.
	ImageRecipeArn *string

	// The image tests configuration of the image pipeline.
	ImageTestsConfiguration *ImageTestsConfiguration

	// The Amazon Resource Name (ARN) of the infrastructure configuration associated
	// with this image pipeline.
	InfrastructureConfigurationArn *string

	// The name of the image pipeline.
	Name *string

	// The platform of the image pipeline.
	Platform Platform

	// The schedule of the image pipeline.
	Schedule *Schedule

	// The status of the image pipeline.
	Status PipelineStatus

	// The tags of this image pipeline.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An image recipe.
type ImageRecipe struct {

	// Before you create a new AMI, Image Builder launches temporary Amazon EC2
	// instances to build and test your image configuration. Instance configuration
	// adds a layer of control over those instances. You can define settings and add
	// scripts to run when an instance is launched from your AMI.
	AdditionalInstanceConfiguration *AdditionalInstanceConfiguration

	// The Amazon Resource Name (ARN) of the image recipe.
	Arn *string

	// The block device mappings to apply when creating images from this recipe.
	BlockDeviceMappings []InstanceBlockDeviceMapping

	// The components of the image recipe.
	Components []ComponentConfiguration

	// The date on which this image recipe was created.
	DateCreated *string

	// The description of the image recipe.
	Description *string

	// The name of the image recipe.
	Name *string

	// The owner of the image recipe.
	Owner *string

	// The parent image of the image recipe.
	ParentImage *string

	// The platform of the image recipe.
	Platform Platform

	// The tags of the image recipe.
	Tags map[string]string

	// Specifies which type of image is created by the recipe - an AMI or a container
	// image.
	Type ImageType

	// The version of the image recipe.
	Version *string

	// The working directory to be used during build and test workflows.
	WorkingDirectory *string

	noSmithyDocumentSerde
}

// A summary of an image recipe.
type ImageRecipeSummary struct {

	// The Amazon Resource Name (ARN) of the image recipe.
	Arn *string

	// The date on which this image recipe was created.
	DateCreated *string

	// The name of the image recipe.
	Name *string

	// The owner of the image recipe.
	Owner *string

	// The parent image of the image recipe.
	ParentImage *string

	// The platform of the image recipe.
	Platform Platform

	// The tags of the image recipe.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Image state shows the image status and the reason for that status.
type ImageState struct {

	// The reason for the image's status.
	Reason *string

	// The status of the image.
	Status ImageStatus

	noSmithyDocumentSerde
}

// An image summary.
type ImageSummary struct {

	// The Amazon Resource Name (ARN) of the image.
	Arn *string

	// The date on which this image was created.
	DateCreated *string

	// The name of the image.
	Name *string

	// The operating system version of the instance. For example, Amazon Linux 2,
	// Ubuntu 18, or Microsoft Windows Server 2019.
	OsVersion *string

	// The output resources produced when creating this image.
	OutputResources *OutputResources

	// The owner of the image.
	Owner *string

	// The platform of the image.
	Platform Platform

	// The state of the image.
	State *ImageState

	// The tags of the image.
	Tags map[string]string

	// Specifies whether this is an AMI or container image.
	Type ImageType

	// The version of the image.
	Version *string

	noSmithyDocumentSerde
}

// Image tests configuration.
type ImageTestsConfiguration struct {

	// Defines if tests should be executed when building this image.
	ImageTestsEnabled *bool

	// The maximum time in minutes that tests are permitted to run.
	TimeoutMinutes *int32

	noSmithyDocumentSerde
}

// The defining characteristics of a specific version of an Image Builder image.
type ImageVersion struct {

	// The Amazon Resource Name (ARN) of a specific version of an Image Builder image.
	// Semantic versioning is included in each object's Amazon Resource Name (ARN), at
	// the level that applies to that object as follows:
	//
	// * Versionless ARNs and Name
	// ARNs do not include specific values in any of the nodes. The nodes are either
	// left off entirely, or they are specified as wildcards, for example: x.x.x.
	//
	// *
	// Version ARNs have only the first three nodes: ..
	//
	// * Build version ARNs have all
	// four nodes, and point to a specific build for a specific version of an object.
	Arn *string

	// The date on which this specific version of the Image Builder image was created.
	DateCreated *string

	// The name of this specific version of an Image Builder image.
	Name *string

	// The operating system version of the Amazon EC2 build instance. For example,
	// Amazon Linux 2, Ubuntu 18, or Microsoft Windows Server 2019.
	OsVersion *string

	// The owner of the image version.
	Owner *string

	// The platform of the image version, for example "Windows" or "Linux".
	Platform Platform

	// Specifies whether this image is an AMI or a container image.
	Type ImageType

	// Details for a specific version of an Image Builder image. This version follows
	// the semantic version syntax. The semantic version has four nodes: ../. You can
	// assign values for the first three, and can filter on all of them. Assignment:
	// For the first three nodes you can assign any positive integer value, including
	// zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder
	// automatically assigns the build number, and that is not open for updates.
	// Patterns: You can use any numeric pattern that adheres to the assignment
	// requirements for the nodes that you can assign. For example, you might choose a
	// software version pattern, such as 1.0.0, or a date, such as 2021.01.01.
	// Filtering: When you retrieve or reference a resource with a semantic version,
	// you can use wildcards (x) to filter your results. When you use a wildcard in any
	// node, all nodes to the right of the first wildcard must also be wildcards. For
	// example, specifying "1.2.x", or "1.x.x" works to filter list results, but
	// neither "1.x.2", nor "x.2.x" will work. You do not have to specify the build -
	// Image Builder automatically uses a wildcard for that, if applicable.
	Version *string

	noSmithyDocumentSerde
}

// Details of the infrastructure configuration.
type InfrastructureConfiguration struct {

	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn *string

	// The date on which the infrastructure configuration was created.
	DateCreated *string

	// The date on which the infrastructure configuration was last updated.
	DateUpdated *string

	// The description of the infrastructure configuration.
	Description *string

	// The instance profile of the infrastructure configuration.
	InstanceProfileName *string

	// The instance types of the infrastructure configuration.
	InstanceTypes []string

	// The Amazon EC2 key pair of the infrastructure configuration.
	KeyPair *string

	// The logging configuration of the infrastructure configuration.
	Logging *Logging

	// The name of the infrastructure configuration.
	Name *string

	// The tags attached to the resource created by Image Builder.
	ResourceTags map[string]string

	// The security group IDs of the infrastructure configuration.
	SecurityGroupIds []string

	// The SNS topic Amazon Resource Name (ARN) of the infrastructure configuration.
	SnsTopicArn *string

	// The subnet ID of the infrastructure configuration.
	SubnetId *string

	// The tags of the infrastructure configuration.
	Tags map[string]string

	// The terminate instance on failure configuration of the infrastructure
	// configuration.
	TerminateInstanceOnFailure *bool

	noSmithyDocumentSerde
}

// The infrastructure used when building Amazon EC2 AMIs.
type InfrastructureConfigurationSummary struct {

	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn *string

	// The date on which the infrastructure configuration was created.
	DateCreated *string

	// The date on which the infrastructure configuration was last updated.
	DateUpdated *string

	// The description of the infrastructure configuration.
	Description *string

	// The instance profile of the infrastructure configuration.
	InstanceProfileName *string

	// The instance types of the infrastructure configuration.
	InstanceTypes []string

	// The name of the infrastructure configuration.
	Name *string

	// The tags attached to the image created by Image Builder.
	ResourceTags map[string]string

	// The tags of the infrastructure configuration.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Defines block device mappings for the instance used to configure your image.
type InstanceBlockDeviceMapping struct {

	// The device to which these mappings apply.
	DeviceName *string

	// Use to manage Amazon EBS-specific configuration for this mapping.
	Ebs *EbsInstanceBlockDeviceSpecification

	// Use to remove a mapping from the parent image.
	NoDevice *string

	// Use to manage instance ephemeral devices.
	VirtualName *string

	noSmithyDocumentSerde
}

// Defines a custom source AMI and block device mapping configurations of an
// instance used for building and testing container images.
type InstanceConfiguration struct {

	// Defines the block devices to attach for building an instance from this Image
	// Builder AMI.
	BlockDeviceMappings []InstanceBlockDeviceMapping

	// The AMI ID to use as the base image for a container build and test instance. If
	// not specified, Image Builder will use the appropriate ECS-optimized AMI as a
	// base image.
	Image *string

	noSmithyDocumentSerde
}

// Describes the configuration for a launch permission. The launch permission
// modification request is sent to the Amazon EC2 ModifyImageAttribute
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyImageAttribute.html)
// API on behalf of the user for each Region they have selected to distribute the
// AMI. To make an AMI public, set the launch permission authorized accounts to
// all. See the examples for making an AMI public at Amazon EC2
// ModifyImageAttribute
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyImageAttribute.html).
type LaunchPermissionConfiguration struct {

	// The name of the group.
	UserGroups []string

	// The account ID.
	UserIds []string

	noSmithyDocumentSerde
}

// Identifies an Amazon EC2 launch template to use for a specific account.
type LaunchTemplateConfiguration struct {

	// Identifies the Amazon EC2 launch template to use.
	//
	// This member is required.
	LaunchTemplateId *string

	// The account ID that this configuration applies to.
	AccountId *string

	// Set the specified Amazon EC2 launch template as the default launch template for
	// the specified account.
	SetDefaultVersion bool

	noSmithyDocumentSerde
}

// Logging configuration defines where Image Builder uploads your logs.
type Logging struct {

	// The Amazon S3 logging configuration.
	S3Logs *S3Logs

	noSmithyDocumentSerde
}

// The resources produced by this image.
type OutputResources struct {

	// The Amazon EC2 AMIs created by this image.
	Amis []Ami

	// Container images that the pipeline has generated and stored in the output
	// repository.
	Containers []Container

	noSmithyDocumentSerde
}

// Amazon S3 logging configuration.
type S3Logs struct {

	// The Amazon S3 bucket in which to store the logs.
	S3BucketName *string

	// The Amazon S3 path in which to store the logs.
	S3KeyPrefix *string

	noSmithyDocumentSerde
}

// A schedule configures how often and when a pipeline will automatically create a
// new image.
type Schedule struct {

	// The condition configures when the pipeline should trigger a new image build.
	// When the pipelineExecutionStartCondition is set to
	// EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE, and you use semantic version
	// filters on the source image or components in your image recipe, EC2 Image
	// Builder will build a new image only when there are new versions of the image or
	// components in your recipe that match the semantic version filter. When it is set
	// to EXPRESSION_MATCH_ONLY, it will build a new image every time the CRON
	// expression matches the current time. For semantic version syntax, see
	// CreateComponent
	// (https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_CreateComponent.html)
	// in the EC2 Image Builder API Reference.
	PipelineExecutionStartCondition PipelineExecutionStartCondition

	// The cron expression determines how often EC2 Image Builder evaluates your
	// pipelineExecutionStartCondition. For information on how to format a cron
	// expression in Image Builder, see Use cron expressions in EC2 Image Builder
	// (https://docs.aws.amazon.com/imagebuilder/latest/userguide/image-builder-cron.html).
	ScheduleExpression *string

	// The timezone that applies to the scheduling expression. For example, "Etc/UTC",
	// "America/Los_Angeles" in the IANA timezone format
	// (https://www.joda.org/joda-time/timezones.html). If not specified this defaults
	// to UTC.
	Timezone *string

	noSmithyDocumentSerde
}

// Contains settings for the SSM agent on your build instance.
type SystemsManagerAgent struct {

	// Controls whether the SSM agent is removed from your final build image, prior to
	// creating the new AMI. If this is set to true, then the agent is removed from the
	// final image. If it's set to false, then the agent is left in, so that it is
	// included in the new AMI. The default value is false.
	UninstallAfterBuild *bool

	noSmithyDocumentSerde
}

// The container repository where the output container image is stored.
type TargetContainerRepository struct {

	// The name of the container repository where the output container image is stored.
	// This name is prefixed by the repository location.
	//
	// This member is required.
	RepositoryName *string

	// Specifies the service in which this image was registered.
	//
	// This member is required.
	Service ContainerRepositoryService

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
