// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amplify app.
func (c *Client) CreateApp(ctx context.Context, params *CreateAppInput, optFns ...func(*Options)) (*CreateAppOutput, error) {
	if params == nil {
		params = &CreateAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApp", params, optFns, c.addOperationCreateAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure used to create apps in Amplify.
type CreateAppInput struct {

	// The name for an Amplify app.
	//
	// This member is required.
	Name *string

	// The personal access token for a third-party source control system for an Amplify
	// app. The personal access token is used to create a webhook and a read-only
	// deploy key. The token is not stored.
	AccessToken *string

	// The automated branch creation configuration for an Amplify app.
	AutoBranchCreationConfig *types.AutoBranchCreationConfig

	// The automated branch creation glob patterns for an Amplify app.
	AutoBranchCreationPatterns []string

	// The credentials for basic authorization for an Amplify app.
	BasicAuthCredentials *string

	// The build specification (build spec) for an Amplify app.
	BuildSpec *string

	// The custom HTTP headers for an Amplify app.
	CustomHeaders *string

	// The custom rewrite and redirect rules for an Amplify app.
	CustomRules []types.CustomRule

	// The description for an Amplify app.
	Description *string

	// Enables automated branch creation for an Amplify app.
	EnableAutoBranchCreation *bool

	// Enables basic authorization for an Amplify app. This will apply to all branches
	// that are part of this app.
	EnableBasicAuth *bool

	// Enables the auto building of branches for an Amplify app.
	EnableBranchAutoBuild *bool

	// Automatically disconnects a branch in the Amplify Console when you delete a
	// branch from your Git repository.
	EnableBranchAutoDeletion *bool

	// The environment variables map for an Amplify app.
	EnvironmentVariables map[string]string

	// The AWS Identity and Access Management (IAM) service role for an Amplify app.
	IamServiceRoleArn *string

	// The OAuth token for a third-party source control system for an Amplify app. The
	// OAuth token is used to create a webhook and a read-only deploy key. The OAuth
	// token is not stored.
	OauthToken *string

	// The platform or framework for an Amplify app.
	Platform types.Platform

	// The repository for an Amplify app.
	Repository *string

	// The tag for an Amplify app.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAppOutput struct {

	// Represents the different branches of a repository for building, deploying, and
	// hosting an Amplify app.
	//
	// This member is required.
	App *types.App

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApp{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAppValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApp(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "CreateApp",
	}
}
