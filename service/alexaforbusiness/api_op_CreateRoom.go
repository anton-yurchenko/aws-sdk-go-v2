// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRoomInput struct {
	_ struct{} `type:"structure"`

	// A unique, user-specified identifier for this request that ensures idempotency.
	ClientRequestToken *string `min:"10" type:"string" idempotencyToken:"true"`

	// The description for the room.
	Description *string `min:"1" type:"string"`

	// The profile ARN for the room.
	ProfileArn *string `type:"string"`

	// The calendar ARN for the room.
	ProviderCalendarId *string `type:"string"`

	// The name for the room.
	//
	// RoomName is a required field
	RoomName *string `min:"1" type:"string" required:"true"`

	// The tags for the room.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateRoomInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRoomInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRoomInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.RoomName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomName"))
	}
	if s.RoomName != nil && len(*s.RoomName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoomName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRoomOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the newly created room in the response.
	RoomArn *string `type:"string"`
}

// String returns the string representation
func (s CreateRoomOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRoom = "CreateRoom"

// CreateRoomRequest returns a request value for making API operation for
// Alexa For Business.
//
// Creates a room with the specified details.
//
//    // Example sending a request using CreateRoomRequest.
//    req := client.CreateRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/CreateRoom
func (c *Client) CreateRoomRequest(input *CreateRoomInput) CreateRoomRequest {
	op := &aws.Operation{
		Name:       opCreateRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRoomInput{}
	}

	req := c.newRequest(op, input, &CreateRoomOutput{})
	return CreateRoomRequest{Request: req, Input: input, Copy: c.CreateRoomRequest}
}

// CreateRoomRequest is the request type for the
// CreateRoom API operation.
type CreateRoomRequest struct {
	*aws.Request
	Input *CreateRoomInput
	Copy  func(*CreateRoomInput) CreateRoomRequest
}

// Send marshals and sends the CreateRoom API request.
func (r CreateRoomRequest) Send(ctx context.Context) (*CreateRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRoomResponse{
		CreateRoomOutput: r.Request.Data.(*CreateRoomOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRoomResponse is the response type for the
// CreateRoom API operation.
type CreateRoomResponse struct {
	*CreateRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRoom request.
func (r *CreateRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}