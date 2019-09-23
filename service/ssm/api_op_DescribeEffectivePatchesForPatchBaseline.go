// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEffectivePatchesForPatchBaselineInput struct {
	_ struct{} `type:"structure"`

	// The ID of the patch baseline to retrieve the effective patches for.
	//
	// BaselineId is a required field
	BaselineId *string `min:"20" type:"string" required:"true"`

	// The maximum number of patches to return (per page).
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEffectivePatchesForPatchBaselineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEffectivePatchesForPatchBaselineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEffectivePatchesForPatchBaselineInput"}

	if s.BaselineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BaselineId"))
	}
	if s.BaselineId != nil && len(*s.BaselineId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("BaselineId", 20))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEffectivePatchesForPatchBaselineOutput struct {
	_ struct{} `type:"structure"`

	// An array of patches and patch status.
	EffectivePatches []EffectivePatch `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEffectivePatchesForPatchBaselineOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEffectivePatchesForPatchBaseline = "DescribeEffectivePatchesForPatchBaseline"

// DescribeEffectivePatchesForPatchBaselineRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves the current effective patches (the patch and the approval state)
// for the specified patch baseline. Note that this API applies only to Windows
// patch baselines.
//
//    // Example sending a request using DescribeEffectivePatchesForPatchBaselineRequest.
//    req := client.DescribeEffectivePatchesForPatchBaselineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeEffectivePatchesForPatchBaseline
func (c *Client) DescribeEffectivePatchesForPatchBaselineRequest(input *DescribeEffectivePatchesForPatchBaselineInput) DescribeEffectivePatchesForPatchBaselineRequest {
	op := &aws.Operation{
		Name:       opDescribeEffectivePatchesForPatchBaseline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEffectivePatchesForPatchBaselineInput{}
	}

	req := c.newRequest(op, input, &DescribeEffectivePatchesForPatchBaselineOutput{})
	return DescribeEffectivePatchesForPatchBaselineRequest{Request: req, Input: input, Copy: c.DescribeEffectivePatchesForPatchBaselineRequest}
}

// DescribeEffectivePatchesForPatchBaselineRequest is the request type for the
// DescribeEffectivePatchesForPatchBaseline API operation.
type DescribeEffectivePatchesForPatchBaselineRequest struct {
	*aws.Request
	Input *DescribeEffectivePatchesForPatchBaselineInput
	Copy  func(*DescribeEffectivePatchesForPatchBaselineInput) DescribeEffectivePatchesForPatchBaselineRequest
}

// Send marshals and sends the DescribeEffectivePatchesForPatchBaseline API request.
func (r DescribeEffectivePatchesForPatchBaselineRequest) Send(ctx context.Context) (*DescribeEffectivePatchesForPatchBaselineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEffectivePatchesForPatchBaselineResponse{
		DescribeEffectivePatchesForPatchBaselineOutput: r.Request.Data.(*DescribeEffectivePatchesForPatchBaselineOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEffectivePatchesForPatchBaselineResponse is the response type for the
// DescribeEffectivePatchesForPatchBaseline API operation.
type DescribeEffectivePatchesForPatchBaselineResponse struct {
	*DescribeEffectivePatchesForPatchBaselineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEffectivePatchesForPatchBaseline request.
func (r *DescribeEffectivePatchesForPatchBaselineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}