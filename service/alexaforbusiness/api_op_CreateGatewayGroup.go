// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateGatewayGroupInput struct {
	_ struct{} `type:"structure"`

	// A unique, user-specified identifier for the request that ensures idempotency.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `min:"10" type:"string" required:"true" idempotencyToken:"true"`

	// The description of the gateway group.
	Description *string `type:"string"`

	// The name of the gateway group.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGatewayGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGatewayGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGatewayGroupInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGatewayGroupOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the created gateway group.
	GatewayGroupArn *string `type:"string"`
}

// String returns the string representation
func (s CreateGatewayGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateGatewayGroup = "CreateGatewayGroup"

// CreateGatewayGroupRequest returns a request value for making API operation for
// Alexa For Business.
//
// Creates a gateway group with the specified details.
//
//    // Example sending a request using CreateGatewayGroupRequest.
//    req := client.CreateGatewayGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateGatewayGroup
func (c *Client) CreateGatewayGroupRequest(input *CreateGatewayGroupInput) CreateGatewayGroupRequest {
	op := &aws.Operation{
		Name:       opCreateGatewayGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateGatewayGroupInput{}
	}

	req := c.newRequest(op, input, &CreateGatewayGroupOutput{})
	return CreateGatewayGroupRequest{Request: req, Input: input, Copy: c.CreateGatewayGroupRequest}
}

// CreateGatewayGroupRequest is the request type for the
// CreateGatewayGroup API operation.
type CreateGatewayGroupRequest struct {
	*aws.Request
	Input *CreateGatewayGroupInput
	Copy  func(*CreateGatewayGroupInput) CreateGatewayGroupRequest
}

// Send marshals and sends the CreateGatewayGroup API request.
func (r CreateGatewayGroupRequest) Send(ctx context.Context) (*CreateGatewayGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGatewayGroupResponse{
		CreateGatewayGroupOutput: r.Request.Data.(*CreateGatewayGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGatewayGroupResponse is the response type for the
// CreateGatewayGroup API operation.
type CreateGatewayGroupResponse struct {
	*CreateGatewayGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGatewayGroup request.
func (r *CreateGatewayGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}