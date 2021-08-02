// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Displays details about a routing control. A routing control has one of two
// states: ON and OFF. You can map the routing control state to the state of an
// Amazon Route 53 health check, which can be used to control routing. To get or
// update the routing control state, see the Recovery Cluster (data plane) API
// actions for Amazon Route 53 Application Recovery Controller.
func (c *Client) DescribeRoutingControl(ctx context.Context, params *DescribeRoutingControlInput, optFns ...func(*Options)) (*DescribeRoutingControlOutput, error) {
	if params == nil {
		params = &DescribeRoutingControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRoutingControl", params, optFns, c.addOperationDescribeRoutingControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRoutingControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRoutingControlInput struct {

	// The Amazon Resource Name (ARN) of the routing control that you're getting
	// details for.
	//
	// This member is required.
	RoutingControlArn *string

	noSmithyDocumentSerde
}

type DescribeRoutingControlOutput struct {

	// Information about the routing control.
	RoutingControl *types.RoutingControl

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRoutingControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRoutingControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRoutingControl{}, middleware.After)
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
	if err = addOpDescribeRoutingControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRoutingControl(options.Region), middleware.Before); err != nil {
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

// DescribeRoutingControlAPIClient is a client that implements the
// DescribeRoutingControl operation.
type DescribeRoutingControlAPIClient interface {
	DescribeRoutingControl(context.Context, *DescribeRoutingControlInput, ...func(*Options)) (*DescribeRoutingControlOutput, error)
}

var _ DescribeRoutingControlAPIClient = (*Client)(nil)

// RoutingControlCreatedWaiterOptions are waiter options for
// RoutingControlCreatedWaiter
type RoutingControlCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// RoutingControlCreatedWaiter will use default minimum delay of 5 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, RoutingControlCreatedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeRoutingControlInput, *DescribeRoutingControlOutput, error) (bool, error)
}

// RoutingControlCreatedWaiter defines the waiters for RoutingControlCreated
type RoutingControlCreatedWaiter struct {
	client DescribeRoutingControlAPIClient

	options RoutingControlCreatedWaiterOptions
}

// NewRoutingControlCreatedWaiter constructs a RoutingControlCreatedWaiter.
func NewRoutingControlCreatedWaiter(client DescribeRoutingControlAPIClient, optFns ...func(*RoutingControlCreatedWaiterOptions)) *RoutingControlCreatedWaiter {
	options := RoutingControlCreatedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = routingControlCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &RoutingControlCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for RoutingControlCreated waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *RoutingControlCreatedWaiter) Wait(ctx context.Context, params *DescribeRoutingControlInput, maxWaitDur time.Duration, optFns ...func(*RoutingControlCreatedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeRoutingControl(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for RoutingControlCreated waiter")
}

func routingControlCreatedStateRetryable(ctx context.Context, input *DescribeRoutingControlInput, output *DescribeRoutingControlOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("RoutingControl.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DEPLOYED"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("RoutingControl.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "PENDING"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// RoutingControlDeletedWaiterOptions are waiter options for
// RoutingControlDeletedWaiter
type RoutingControlDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// RoutingControlDeletedWaiter will use default minimum delay of 5 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, RoutingControlDeletedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeRoutingControlInput, *DescribeRoutingControlOutput, error) (bool, error)
}

// RoutingControlDeletedWaiter defines the waiters for RoutingControlDeleted
type RoutingControlDeletedWaiter struct {
	client DescribeRoutingControlAPIClient

	options RoutingControlDeletedWaiterOptions
}

// NewRoutingControlDeletedWaiter constructs a RoutingControlDeletedWaiter.
func NewRoutingControlDeletedWaiter(client DescribeRoutingControlAPIClient, optFns ...func(*RoutingControlDeletedWaiterOptions)) *RoutingControlDeletedWaiter {
	options := RoutingControlDeletedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = routingControlDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &RoutingControlDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for RoutingControlDeleted waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *RoutingControlDeletedWaiter) Wait(ctx context.Context, params *DescribeRoutingControlInput, maxWaitDur time.Duration, optFns ...func(*RoutingControlDeletedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeRoutingControl(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for RoutingControlDeleted waiter")
}

func routingControlDeletedStateRetryable(ctx context.Context, input *DescribeRoutingControlInput, output *DescribeRoutingControlOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("RoutingControl.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "PENDING_DELETION"
		value, ok := pathValue.(types.Status)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeRoutingControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-control-config",
		OperationName: "DescribeRoutingControl",
	}
}
