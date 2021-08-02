// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the permission sets that are provisioned to a specified Amazon Web
// Services account.
func (c *Client) ListPermissionSetsProvisionedToAccount(ctx context.Context, params *ListPermissionSetsProvisionedToAccountInput, optFns ...func(*Options)) (*ListPermissionSetsProvisionedToAccountOutput, error) {
	if params == nil {
		params = &ListPermissionSetsProvisionedToAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissionSetsProvisionedToAccount", params, optFns, c.addOperationListPermissionSetsProvisionedToAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionSetsProvisionedToAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionSetsProvisionedToAccountInput struct {

	// The identifier of the Amazon Web Services account from which to list the
	// assignments.
	//
	// This member is required.
	AccountId *string

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services
	// Service Namespaces in the Amazon Web Services General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The maximum number of results to display for the assignment.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// The status object for the permission set provisioning operation.
	ProvisioningStatus types.ProvisioningStatus

	noSmithyDocumentSerde
}

type ListPermissionSetsProvisionedToAccountOutput struct {

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// Defines the level of access that an Amazon Web Services account has.
	PermissionSets []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPermissionSetsProvisionedToAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPermissionSetsProvisionedToAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPermissionSetsProvisionedToAccount{}, middleware.After)
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
	if err = addOpListPermissionSetsProvisionedToAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissionSetsProvisionedToAccount(options.Region), middleware.Before); err != nil {
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

// ListPermissionSetsProvisionedToAccountAPIClient is a client that implements the
// ListPermissionSetsProvisionedToAccount operation.
type ListPermissionSetsProvisionedToAccountAPIClient interface {
	ListPermissionSetsProvisionedToAccount(context.Context, *ListPermissionSetsProvisionedToAccountInput, ...func(*Options)) (*ListPermissionSetsProvisionedToAccountOutput, error)
}

var _ ListPermissionSetsProvisionedToAccountAPIClient = (*Client)(nil)

// ListPermissionSetsProvisionedToAccountPaginatorOptions is the paginator options
// for ListPermissionSetsProvisionedToAccount
type ListPermissionSetsProvisionedToAccountPaginatorOptions struct {
	// The maximum number of results to display for the assignment.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPermissionSetsProvisionedToAccountPaginator is a paginator for
// ListPermissionSetsProvisionedToAccount
type ListPermissionSetsProvisionedToAccountPaginator struct {
	options   ListPermissionSetsProvisionedToAccountPaginatorOptions
	client    ListPermissionSetsProvisionedToAccountAPIClient
	params    *ListPermissionSetsProvisionedToAccountInput
	nextToken *string
	firstPage bool
}

// NewListPermissionSetsProvisionedToAccountPaginator returns a new
// ListPermissionSetsProvisionedToAccountPaginator
func NewListPermissionSetsProvisionedToAccountPaginator(client ListPermissionSetsProvisionedToAccountAPIClient, params *ListPermissionSetsProvisionedToAccountInput, optFns ...func(*ListPermissionSetsProvisionedToAccountPaginatorOptions)) *ListPermissionSetsProvisionedToAccountPaginator {
	if params == nil {
		params = &ListPermissionSetsProvisionedToAccountInput{}
	}

	options := ListPermissionSetsProvisionedToAccountPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPermissionSetsProvisionedToAccountPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPermissionSetsProvisionedToAccountPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPermissionSetsProvisionedToAccount page.
func (p *ListPermissionSetsProvisionedToAccountPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPermissionSetsProvisionedToAccountOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListPermissionSetsProvisionedToAccount(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPermissionSetsProvisionedToAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "ListPermissionSetsProvisionedToAccount",
	}
}
