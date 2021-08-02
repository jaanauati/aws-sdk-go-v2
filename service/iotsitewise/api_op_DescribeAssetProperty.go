// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about an asset property. When you call this operation for
// an attribute property, this response includes the default attribute value that
// you define in the asset model. If you update the default value in the model,
// this operation's response includes the new default value. This operation doesn't
// return the value of the asset property. To get the value of an asset property,
// use GetAssetPropertyValue
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_GetAssetPropertyValue.html).
func (c *Client) DescribeAssetProperty(ctx context.Context, params *DescribeAssetPropertyInput, optFns ...func(*Options)) (*DescribeAssetPropertyOutput, error) {
	if params == nil {
		params = &DescribeAssetPropertyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssetProperty", params, optFns, c.addOperationDescribeAssetPropertyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssetPropertyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssetPropertyInput struct {

	// The ID of the asset.
	//
	// This member is required.
	AssetId *string

	// The ID of the asset property.
	//
	// This member is required.
	PropertyId *string

	noSmithyDocumentSerde
}

type DescribeAssetPropertyOutput struct {

	// The ID of the asset.
	//
	// This member is required.
	AssetId *string

	// The ID of the asset model.
	//
	// This member is required.
	AssetModelId *string

	// The name of the asset.
	//
	// This member is required.
	AssetName *string

	// The asset property's definition, alias, and notification state. This response
	// includes this object for normal asset properties. If you describe an asset
	// property in a composite model, this response includes the asset property
	// information in compositeModel.
	AssetProperty *types.Property

	// The composite asset model that declares this asset property, if this asset
	// property exists in a composite model.
	CompositeModel *types.CompositeModelProperty

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAssetPropertyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAssetProperty{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAssetProperty{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeAssetPropertyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeAssetPropertyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssetProperty(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeAssetPropertyMiddleware struct {
}

func (*endpointPrefix_opDescribeAssetPropertyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeAssetPropertyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeAssetPropertyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeAssetPropertyMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAssetProperty(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeAssetProperty",
	}
}
