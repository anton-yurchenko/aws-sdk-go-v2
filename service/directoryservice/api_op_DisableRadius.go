// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the DisableRadius operation.
type DisableRadiusInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory for which to disable MFA.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisableRadiusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableRadiusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableRadiusInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the results of the DisableRadius operation.
type DisableRadiusOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableRadiusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisableRadius = "DisableRadius"

// DisableRadiusRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Disables multi-factor authentication (MFA) with the Remote Authentication
// Dial In User Service (RADIUS) server for an AD Connector or Microsoft AD
// directory.
//
//    // Example sending a request using DisableRadiusRequest.
//    req := client.DisableRadiusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DisableRadius
func (c *Client) DisableRadiusRequest(input *DisableRadiusInput) DisableRadiusRequest {
	op := &aws.Operation{
		Name:       opDisableRadius,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableRadiusInput{}
	}

	req := c.newRequest(op, input, &DisableRadiusOutput{})
	return DisableRadiusRequest{Request: req, Input: input, Copy: c.DisableRadiusRequest}
}

// DisableRadiusRequest is the request type for the
// DisableRadius API operation.
type DisableRadiusRequest struct {
	*aws.Request
	Input *DisableRadiusInput
	Copy  func(*DisableRadiusInput) DisableRadiusRequest
}

// Send marshals and sends the DisableRadius API request.
func (r DisableRadiusRequest) Send(ctx context.Context) (*DisableRadiusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableRadiusResponse{
		DisableRadiusOutput: r.Request.Data.(*DisableRadiusOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableRadiusResponse is the response type for the
// DisableRadius API operation.
type DisableRadiusResponse struct {
	*DisableRadiusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableRadius request.
func (r *DisableRadiusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}