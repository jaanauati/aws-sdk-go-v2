// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

//
type BadRequestException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested Outposts bucket name is not available. The bucket namespace is
// shared by all users of the Outposts in this Region. Select a different name and
// try again.
type BucketAlreadyExists struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BucketAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BucketAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BucketAlreadyExists) ErrorCode() string             { return "BucketAlreadyExists" }
func (e *BucketAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Outposts bucket you tried to create already exists, and you own it.
type BucketAlreadyOwnedByYou struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *BucketAlreadyOwnedByYou) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BucketAlreadyOwnedByYou) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BucketAlreadyOwnedByYou) ErrorCode() string             { return "BucketAlreadyOwnedByYou" }
func (e *BucketAlreadyOwnedByYou) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//
type IdempotencyException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *IdempotencyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotencyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotencyException) ErrorCode() string             { return "IdempotencyException" }
func (e *IdempotencyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//
type InternalServiceException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string             { return "InternalServiceException" }
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

//
type InvalidNextTokenException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//
type InvalidRequestException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//
type JobStatusException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *JobStatusException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *JobStatusException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *JobStatusException) ErrorCode() string             { return "JobStatusException" }
func (e *JobStatusException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon S3 throws this exception if you make a GetPublicAccessBlock request
// against an account that doesn't have a PublicAccessBlockConfiguration set.
type NoSuchPublicAccessBlockConfiguration struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NoSuchPublicAccessBlockConfiguration) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchPublicAccessBlockConfiguration) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchPublicAccessBlockConfiguration) ErrorCode() string {
	return "NoSuchPublicAccessBlockConfiguration"
}
func (e *NoSuchPublicAccessBlockConfiguration) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

//
type NotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//
type TooManyRequestsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string             { return "TooManyRequestsException" }
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon S3 throws this exception if you have too many tags in your tag set.
type TooManyTagsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
