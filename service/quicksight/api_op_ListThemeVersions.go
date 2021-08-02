// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the versions of the themes in the current Amazon Web Services
// account;.
func (c *Client) ListThemeVersions(ctx context.Context, params *ListThemeVersionsInput, optFns ...func(*Options)) (*ListThemeVersionsOutput, error) {
	if params == nil {
		params = &ListThemeVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThemeVersions", params, optFns, c.addOperationListThemeVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThemeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThemeVersionsInput struct {

	// The ID of the Amazon Web Services account; that contains the themes that you're
	// listing.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the theme.
	//
	// This member is required.
	ThemeId *string

	// The maximum number of results to be returned per request.
	MaxResults int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListThemeVersionsOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// A structure containing a list of all the versions of the specified theme.
	ThemeVersionSummaryList []types.ThemeVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThemeVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThemeVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThemeVersions{}, middleware.After)
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
	if err = addOpListThemeVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThemeVersions(options.Region), middleware.Before); err != nil {
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

// ListThemeVersionsAPIClient is a client that implements the ListThemeVersions
// operation.
type ListThemeVersionsAPIClient interface {
	ListThemeVersions(context.Context, *ListThemeVersionsInput, ...func(*Options)) (*ListThemeVersionsOutput, error)
}

var _ ListThemeVersionsAPIClient = (*Client)(nil)

// ListThemeVersionsPaginatorOptions is the paginator options for ListThemeVersions
type ListThemeVersionsPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThemeVersionsPaginator is a paginator for ListThemeVersions
type ListThemeVersionsPaginator struct {
	options   ListThemeVersionsPaginatorOptions
	client    ListThemeVersionsAPIClient
	params    *ListThemeVersionsInput
	nextToken *string
	firstPage bool
}

// NewListThemeVersionsPaginator returns a new ListThemeVersionsPaginator
func NewListThemeVersionsPaginator(client ListThemeVersionsAPIClient, params *ListThemeVersionsInput, optFns ...func(*ListThemeVersionsPaginatorOptions)) *ListThemeVersionsPaginator {
	if params == nil {
		params = &ListThemeVersionsInput{}
	}

	options := ListThemeVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThemeVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThemeVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListThemeVersions page.
func (p *ListThemeVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThemeVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListThemeVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListThemeVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListThemeVersions",
	}
}
