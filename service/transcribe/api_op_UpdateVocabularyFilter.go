// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a vocabulary filter with a new list of filtered words.
func (c *Client) UpdateVocabularyFilter(ctx context.Context, params *UpdateVocabularyFilterInput, optFns ...func(*Options)) (*UpdateVocabularyFilterOutput, error) {
	if params == nil {
		params = &UpdateVocabularyFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVocabularyFilter", params, optFns, c.addOperationUpdateVocabularyFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVocabularyFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVocabularyFilterInput struct {

	// The name of the vocabulary filter to update. If you try to update a vocabulary
	// filter with the same name as another vocabulary filter, you get a
	// ConflictException error.
	//
	// This member is required.
	VocabularyFilterName *string

	// The Amazon S3 location of a text file used as input to create the vocabulary
	// filter. Only use characters from the character set defined for custom
	// vocabularies. For a list of character sets, see Character Sets for Custom
	// Vocabularies
	// (https://docs.aws.amazon.com/transcribe/latest/dg/how-vocabulary.html#charsets).
	// The specified file must be less than 50 KB of UTF-8 characters. If you provide
	// the location of a list of words in the VocabularyFilterFileUri parameter, you
	// can't use the Words parameter.
	VocabularyFilterFileUri *string

	// The words to use in the vocabulary filter. Only use characters from the
	// character set defined for custom vocabularies. For a list of character sets, see
	// Character Sets for Custom Vocabularies
	// (https://docs.aws.amazon.com/transcribe/latest/dg/how-vocabulary.html#charsets).
	// If you provide a list of words in the Words parameter, you can't use the
	// VocabularyFilterFileUri parameter.
	Words []string

	noSmithyDocumentSerde
}

type UpdateVocabularyFilterOutput struct {

	// The language code of the words in the vocabulary filter.
	LanguageCode types.LanguageCode

	// The date and time that the vocabulary filter was updated.
	LastModifiedTime *time.Time

	// The name of the updated vocabulary filter.
	VocabularyFilterName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVocabularyFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateVocabularyFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateVocabularyFilter{}, middleware.After)
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
	if err = addOpUpdateVocabularyFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVocabularyFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVocabularyFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "UpdateVocabularyFilter",
	}
}
