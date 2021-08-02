// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This action deletes an Amazon S3 on Outposts bucket's tags. To delete an S3
// bucket tags, see DeleteBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html)
// in the Amazon S3 API Reference. Deletes the tags from the Outposts bucket. For
// more information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in
// Amazon S3 User Guide. To use this action, you must have permission to perform
// the PutBucketTagging action. By default, the bucket owner has this permission
// and can grant this permission to others. All Amazon S3 on Outposts REST API
// requests for this action require an additional parameter of x-amz-outpost-id to
// be passed with the request and an S3 on Outposts endpoint hostname prefix
// instead of s3-control. For an example of the request syntax for Amazon S3 on
// Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// x-amz-outpost-id derived using the access point ARN, see the Examples
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html#API_control_DeleteBucketTagging_Examples)
// section. The following actions are related to DeleteBucketTagging:
//
// *
// GetBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html)
//
// *
// PutBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html)
func (c *Client) DeleteBucketTagging(ctx context.Context, params *DeleteBucketTaggingInput, optFns ...func(*Options)) (*DeleteBucketTaggingOutput, error) {
	if params == nil {
		params = &DeleteBucketTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketTagging", params, optFns, c.addOperationDeleteBucketTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketTaggingInput struct {

	// The account ID of the Outposts bucket tag set to be removed.
	//
	// This member is required.
	AccountId *string

	// The bucket ARN that has the tag set to be removed. For using this parameter with
	// Amazon S3 on Outposts with the REST API, you must specify the name and the
	// x-amz-outpost-id as well. For using this parameter with S3 on Outposts with the
	// Amazon Web Services SDK and CLI, you must specify the ARN of the bucket accessed
	// in the format arn:aws:s3-outposts:::outpost//bucket/. For example, to access the
	// bucket reports through outpost my-outpost owned by account 123456789012 in
	// Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	noSmithyDocumentSerde
}

type DeleteBucketTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBucketTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketTagging{}, middleware.After)
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
	if err = addEndpointPrefix_opDeleteBucketTaggingMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteBucketTaggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketTagging(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addDeleteBucketTaggingUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opDeleteBucketTaggingMiddleware struct {
}

func (*endpointPrefix_opDeleteBucketTaggingMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteBucketTaggingMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteBucketTaggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteBucketTaggingMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteBucketTaggingMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBucketTagging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketTagging",
	}
}

func copyDeleteBucketTaggingInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteBucketTaggingInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteBucketTaggingInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getDeleteBucketTaggingARNMember(input interface{}) (*string, bool) {
	in := input.(*DeleteBucketTaggingInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setDeleteBucketTaggingARNMember(input interface{}, v string) error {
	in := input.(*DeleteBucketTaggingInput)
	in.Bucket = &v
	return nil
}
func backFillDeleteBucketTaggingAccountID(input interface{}, v string) error {
	in := input.(*DeleteBucketTaggingInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteBucketTaggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getDeleteBucketTaggingARNMember,
			BackfillAccountID: backFillDeleteBucketTaggingAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setDeleteBucketTaggingARNMember,
			CopyInput:         copyDeleteBucketTaggingInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
