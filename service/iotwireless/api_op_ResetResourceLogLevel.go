// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the log-level override, if any, for a specific resource-ID and
// resource-type. It can be used for a wireless device or a wireless gateway.
func (c *Client) ResetResourceLogLevel(ctx context.Context, params *ResetResourceLogLevelInput, optFns ...func(*Options)) (*ResetResourceLogLevelOutput, error) {
	if params == nil {
		params = &ResetResourceLogLevelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetResourceLogLevel", params, optFns, c.addOperationResetResourceLogLevelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetResourceLogLevelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetResourceLogLevelInput struct {

	// The identifier of the resource. For a Wireless Device, it is the wireless device
	// ID. For a wireless gateway, it is the wireless gateway ID.
	//
	// This member is required.
	ResourceIdentifier *string

	// The type of the resource, which can be WirelessDevice or WirelessGateway.
	//
	// This member is required.
	ResourceType *string

	noSmithyDocumentSerde
}

type ResetResourceLogLevelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetResourceLogLevelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpResetResourceLogLevel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpResetResourceLogLevel{}, middleware.After)
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
	if err = addOpResetResourceLogLevelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetResourceLogLevel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetResourceLogLevel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "ResetResourceLogLevel",
	}
}
