// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified properties of a given tracker resource.
func (c *Client) UpdateTracker(ctx context.Context, params *UpdateTrackerInput, optFns ...func(*Options)) (*UpdateTrackerOutput, error) {
	if params == nil {
		params = &UpdateTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTracker", params, optFns, c.addOperationUpdateTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTrackerInput struct {

	// The name of the tracker resource to update.
	//
	// This member is required.
	TrackerName *string

	// Updates the description for the tracker resource.
	Description *string

	// Updates the pricing plan for the tracker resource. For more information about
	// each pricing plan option restrictions, see Amazon Location Service pricing
	// (https://aws.amazon.com/location/pricing/).
	PricingPlan types.PricingPlan

	// Updates the data provider for the tracker resource. A required value for the
	// following pricing plans: MobileAssetTracking| MobileAssetManagement For more
	// information about data providers
	// (https://aws.amazon.com/location/data-providers/) and pricing plans
	// (https://aws.amazon.com/location/pricing/), see the Amazon Location Service
	// product page This can only be updated when updating the PricingPlan in the same
	// request. Amazon Location Service uses PricingPlanDataSource to calculate billing
	// for your tracker resource. Your data won't be shared with the data provider, and
	// will remain in your AWS account and Region unless you move it.
	PricingPlanDataSource *string

	noSmithyDocumentSerde
}

type UpdateTrackerOutput struct {

	// The Amazon Resource Name (ARN) of the updated tracker resource. Used to specify
	// a resource across AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:tracker/ExampleTracker
	//
	// This member is required.
	TrackerArn *string

	// The name of the updated tracker resource.
	//
	// This member is required.
	TrackerName *string

	// The timestamp for when the tracker resource was last updated in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTracker{}, middleware.After)
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
	if err = addOpUpdateTrackerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTracker(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTracker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "UpdateTracker",
	}
}
