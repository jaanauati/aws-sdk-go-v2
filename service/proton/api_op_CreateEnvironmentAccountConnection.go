// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an environment account connection in an environment account so that
// environment infrastructure resources can be provisioned in the environment
// account from a management account. An environment account connection is a secure
// bi-directional connection between a management account and an environment
// account that maintains authorization and permissions. For more information, see
// Environment account connections
// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-env-account-connections.html)
// in the AWS Proton Administrator guide.
func (c *Client) CreateEnvironmentAccountConnection(ctx context.Context, params *CreateEnvironmentAccountConnectionInput, optFns ...func(*Options)) (*CreateEnvironmentAccountConnectionOutput, error) {
	if params == nil {
		params = &CreateEnvironmentAccountConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentAccountConnection", params, optFns, c.addOperationCreateEnvironmentAccountConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentAccountConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentAccountConnectionInput struct {

	// The name of the AWS Proton environment that's created in the associated
	// management account.
	//
	// This member is required.
	EnvironmentName *string

	// The ID of the management account that accepts or rejects the environment account
	// connection. You create an manage the AWS Proton environment in this account. If
	// the management account accepts the environment account connection, AWS Proton
	// can use the associated IAM role to provision environment infrastructure
	// resources in the associated environment account.
	//
	// This member is required.
	ManagementAccountId *string

	// The Amazon Resource Name (ARN) of the IAM service role that's created in the
	// environment account. AWS Proton uses this role to provision infrastructure
	// resources in the associated environment account.
	//
	// This member is required.
	RoleArn *string

	// When included, if two identicial requests are made with the same client token,
	// AWS Proton returns the environment account connection that the first request
	// created.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateEnvironmentAccountConnectionOutput struct {

	// The environment account connection detail data that's returned by AWS Proton.
	//
	// This member is required.
	EnvironmentAccountConnection *types.EnvironmentAccountConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentAccountConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateEnvironmentAccountConnection{}, middleware.After)
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
	if err = addOpCreateEnvironmentAccountConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentAccountConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentAccountConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "CreateEnvironmentAccountConnection",
	}
}
