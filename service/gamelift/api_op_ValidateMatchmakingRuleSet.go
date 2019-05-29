// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/ValidateMatchmakingRuleSetInput
type ValidateMatchmakingRuleSetInput struct {
	_ struct{} `type:"structure"`

	// Collection of matchmaking rules to validate, formatted as a JSON string.
	//
	// RuleSetBody is a required field
	RuleSetBody *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ValidateMatchmakingRuleSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ValidateMatchmakingRuleSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ValidateMatchmakingRuleSetInput"}

	if s.RuleSetBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetBody"))
	}
	if s.RuleSetBody != nil && len(*s.RuleSetBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleSetBody", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/ValidateMatchmakingRuleSetOutput
type ValidateMatchmakingRuleSetOutput struct {
	_ struct{} `type:"structure"`

	// Response indicating whether or not the rule set is valid.
	Valid *bool `type:"boolean"`
}

// String returns the string representation
func (s ValidateMatchmakingRuleSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opValidateMatchmakingRuleSet = "ValidateMatchmakingRuleSet"

// ValidateMatchmakingRuleSetRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Validates the syntax of a matchmaking rule or rule set. This operation checks
// that the rule set is using syntactically correct JSON and that it conforms
// to allowed property expressions. To validate syntax, provide a rule set JSON
// string.
//
// Learn more
//
//    * Build a Rule Set (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-rulesets.html)
//
// Related operations
//
//    * CreateMatchmakingConfiguration
//
//    * DescribeMatchmakingConfigurations
//
//    * UpdateMatchmakingConfiguration
//
//    * DeleteMatchmakingConfiguration
//
//    * CreateMatchmakingRuleSet
//
//    * DescribeMatchmakingRuleSets
//
//    * ValidateMatchmakingRuleSet
//
//    * DeleteMatchmakingRuleSet
//
//    // Example sending a request using ValidateMatchmakingRuleSetRequest.
//    req := client.ValidateMatchmakingRuleSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/ValidateMatchmakingRuleSet
func (c *Client) ValidateMatchmakingRuleSetRequest(input *ValidateMatchmakingRuleSetInput) ValidateMatchmakingRuleSetRequest {
	op := &aws.Operation{
		Name:       opValidateMatchmakingRuleSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ValidateMatchmakingRuleSetInput{}
	}

	req := c.newRequest(op, input, &ValidateMatchmakingRuleSetOutput{})
	return ValidateMatchmakingRuleSetRequest{Request: req, Input: input, Copy: c.ValidateMatchmakingRuleSetRequest}
}

// ValidateMatchmakingRuleSetRequest is the request type for the
// ValidateMatchmakingRuleSet API operation.
type ValidateMatchmakingRuleSetRequest struct {
	*aws.Request
	Input *ValidateMatchmakingRuleSetInput
	Copy  func(*ValidateMatchmakingRuleSetInput) ValidateMatchmakingRuleSetRequest
}

// Send marshals and sends the ValidateMatchmakingRuleSet API request.
func (r ValidateMatchmakingRuleSetRequest) Send(ctx context.Context) (*ValidateMatchmakingRuleSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ValidateMatchmakingRuleSetResponse{
		ValidateMatchmakingRuleSetOutput: r.Request.Data.(*ValidateMatchmakingRuleSetOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ValidateMatchmakingRuleSetResponse is the response type for the
// ValidateMatchmakingRuleSet API operation.
type ValidateMatchmakingRuleSetResponse struct {
	*ValidateMatchmakingRuleSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ValidateMatchmakingRuleSet request.
func (r *ValidateMatchmakingRuleSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}