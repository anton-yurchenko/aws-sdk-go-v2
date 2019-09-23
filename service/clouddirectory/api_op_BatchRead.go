// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type BatchReadInput struct {
	_ struct{} `type:"structure"`

	// Represents the manner and timing in which the successful write or update
	// of an object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Directory. For
	// more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// A list of operations that are part of the batch.
	//
	// Operations is a required field
	Operations []BatchReadOperation `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchReadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchReadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchReadInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.Operations == nil {
		invalidParams.Add(aws.NewErrParamRequired("Operations"))
	}
	if s.Operations != nil {
		for i, v := range s.Operations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Operations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchReadInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Operations != nil {
		v := s.Operations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Operations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-consistency-level", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type BatchReadOutput struct {
	_ struct{} `type:"structure"`

	// A list of all the responses for each batch read.
	Responses []BatchReadOperationResponse `type:"list"`
}

// String returns the string representation
func (s BatchReadOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchReadOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Responses != nil {
		v := s.Responses

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Responses", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchRead = "BatchRead"

// BatchReadRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Performs all the read operations in a batch.
//
//    // Example sending a request using BatchReadRequest.
//    req := client.BatchReadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/BatchRead
func (c *Client) BatchReadRequest(input *BatchReadInput) BatchReadRequest {
	op := &aws.Operation{
		Name:       opBatchRead,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/batchread",
	}

	if input == nil {
		input = &BatchReadInput{}
	}

	req := c.newRequest(op, input, &BatchReadOutput{})
	return BatchReadRequest{Request: req, Input: input, Copy: c.BatchReadRequest}
}

// BatchReadRequest is the request type for the
// BatchRead API operation.
type BatchReadRequest struct {
	*aws.Request
	Input *BatchReadInput
	Copy  func(*BatchReadInput) BatchReadRequest
}

// Send marshals and sends the BatchRead API request.
func (r BatchReadRequest) Send(ctx context.Context) (*BatchReadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchReadResponse{
		BatchReadOutput: r.Request.Data.(*BatchReadOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchReadResponse is the response type for the
// BatchRead API operation.
type BatchReadResponse struct {
	*BatchReadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchRead request.
func (r *BatchReadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}