// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of the reusable delegation sets that are associated with the
// current account.
func (c *Client) ListReusableDelegationSets(ctx context.Context, params *ListReusableDelegationSetsInput, optFns ...func(*Options)) (*ListReusableDelegationSetsOutput, error) {
	if params == nil {
		params = &ListReusableDelegationSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReusableDelegationSets", params, optFns, c.addOperationListReusableDelegationSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReusableDelegationSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to get a list of the reusable delegation sets that are associated with
// the current account.
type ListReusableDelegationSetsInput struct {

	// If the value of IsTruncated in the previous response was true, you have more
	// reusable delegation sets. To get another group, submit another
	// ListReusableDelegationSets request. For the value of marker, specify the value
	// of NextMarker from the previous response, which is the ID of the first reusable
	// delegation set that Amazon Route 53 will return if you submit another request.
	// If the value of IsTruncated in the previous response was false, there are no
	// more reusable delegation sets to get.
	Marker *string

	// The number of reusable delegation sets that you want Amazon Route 53 to return
	// in the response to this request. If you specify a value greater than 100, Route
	// 53 returns only the first 100 reusable delegation sets.
	MaxItems *int32

	noSmithyDocumentSerde
}

// A complex type that contains information about the reusable delegation sets that
// are associated with the current account.
type ListReusableDelegationSetsOutput struct {

	// A complex type that contains one DelegationSet element for each reusable
	// delegation set that was created by the current account.
	//
	// This member is required.
	DelegationSets []types.DelegationSet

	// A flag that indicates whether there are more reusable delegation sets to be
	// listed.
	//
	// This member is required.
	IsTruncated bool

	// For the second and subsequent calls to ListReusableDelegationSets, Marker is the
	// value that you specified for the marker parameter in the request that produced
	// the current response.
	//
	// This member is required.
	Marker *string

	// The value that you specified for the maxitems parameter in the call to
	// ListReusableDelegationSets that produced the current response.
	//
	// This member is required.
	MaxItems *int32

	// If IsTruncated is true, the value of NextMarker identifies the next reusable
	// delegation set that Amazon Route 53 will return if you submit another
	// ListReusableDelegationSets request and specify the value of NextMarker in the
	// marker parameter.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReusableDelegationSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListReusableDelegationSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListReusableDelegationSets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReusableDelegationSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListReusableDelegationSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListReusableDelegationSets",
	}
}
