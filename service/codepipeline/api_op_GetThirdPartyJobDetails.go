// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetThirdPartyJobDetails action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetThirdPartyJobDetailsInput
type GetThirdPartyJobDetailsInput struct {
	_ struct{} `type:"structure"`

	// The clientToken portion of the clientId and clientToken pair used to verify
	// that the calling entity is allowed access to the job and its details.
	//
	// ClientToken is a required field
	ClientToken *string `locationName:"clientToken" min:"1" type:"string" required:"true"`

	// The unique system-generated ID used for identifying the job.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetThirdPartyJobDetailsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetThirdPartyJobDetailsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetThirdPartyJobDetailsInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetThirdPartyJobDetails action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetThirdPartyJobDetailsOutput
type GetThirdPartyJobDetailsOutput struct {
	_ struct{} `type:"structure"`

	// The details of the job, including any protected values defined for the job.
	JobDetails *ThirdPartyJobDetails `locationName:"jobDetails" type:"structure"`
}

// String returns the string representation
func (s GetThirdPartyJobDetailsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetThirdPartyJobDetails = "GetThirdPartyJobDetails"

// GetThirdPartyJobDetailsRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Requests the details of a job for a third party action. Only used for partner
// actions.
//
// When this API is called, AWS CodePipeline returns temporary credentials for
// the Amazon S3 bucket used to store artifacts for the pipeline, if the action
// requires access to that Amazon S3 bucket for input or output artifacts. Additionally,
// this API returns any secret values defined for the action.
//
//    // Example sending a request using GetThirdPartyJobDetailsRequest.
//    req := client.GetThirdPartyJobDetailsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/GetThirdPartyJobDetails
func (c *Client) GetThirdPartyJobDetailsRequest(input *GetThirdPartyJobDetailsInput) GetThirdPartyJobDetailsRequest {
	op := &aws.Operation{
		Name:       opGetThirdPartyJobDetails,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetThirdPartyJobDetailsInput{}
	}

	req := c.newRequest(op, input, &GetThirdPartyJobDetailsOutput{})
	return GetThirdPartyJobDetailsRequest{Request: req, Input: input, Copy: c.GetThirdPartyJobDetailsRequest}
}

// GetThirdPartyJobDetailsRequest is the request type for the
// GetThirdPartyJobDetails API operation.
type GetThirdPartyJobDetailsRequest struct {
	*aws.Request
	Input *GetThirdPartyJobDetailsInput
	Copy  func(*GetThirdPartyJobDetailsInput) GetThirdPartyJobDetailsRequest
}

// Send marshals and sends the GetThirdPartyJobDetails API request.
func (r GetThirdPartyJobDetailsRequest) Send(ctx context.Context) (*GetThirdPartyJobDetailsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetThirdPartyJobDetailsResponse{
		GetThirdPartyJobDetailsOutput: r.Request.Data.(*GetThirdPartyJobDetailsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetThirdPartyJobDetailsResponse is the response type for the
// GetThirdPartyJobDetails API operation.
type GetThirdPartyJobDetailsResponse struct {
	*GetThirdPartyJobDetailsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetThirdPartyJobDetails request.
func (r *GetThirdPartyJobDetailsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}