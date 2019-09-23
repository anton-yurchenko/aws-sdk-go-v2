// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure used to request a project be created.
type CreateProjectInput struct {
	_ struct{} `type:"structure" payload:"Contents"`

	// ZIP or YAML file which contains configuration settings to be used when creating
	// the project. This may be the contents of the file downloaded from the URL
	// provided in an export project operation.
	Contents []byte `locationName:"contents" type:"blob"`

	// Name of the project.
	Name *string `location:"querystring" locationName:"name" type:"string"`

	// Default region where project resources should be created.
	Region *string `location:"querystring" locationName:"region" type:"string"`

	// Unique identifier for an exported snapshot of project configuration. This
	// snapshot identifier is included in the share URL when a project is exported.
	SnapshotId *string `location:"querystring" locationName:"snapshotId" type:"string"`
}

// String returns the string representation
func (s CreateProjectInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Contents != nil {
		v := s.Contents

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "contents", protocol.BytesStream(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Region != nil {
		v := *s.Region

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "region", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SnapshotId != nil {
		v := *s.SnapshotId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "snapshotId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure used in response to a request to create a project.
type CreateProjectOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about the created AWS Mobile Hub project.
	Details *ProjectDetails `locationName:"details" type:"structure"`
}

// String returns the string representation
func (s CreateProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Details != nil {
		v := s.Details

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "details", v, metadata)
	}
	return nil
}

const opCreateProject = "CreateProject"

// CreateProjectRequest returns a request value for making API operation for
// AWS Mobile.
//
// Creates an AWS Mobile Hub project.
//
//    // Example sending a request using CreateProjectRequest.
//    req := client.CreateProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/CreateProject
func (c *Client) CreateProjectRequest(input *CreateProjectInput) CreateProjectRequest {
	op := &aws.Operation{
		Name:       opCreateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/projects",
	}

	if input == nil {
		input = &CreateProjectInput{}
	}

	req := c.newRequest(op, input, &CreateProjectOutput{})
	return CreateProjectRequest{Request: req, Input: input, Copy: c.CreateProjectRequest}
}

// CreateProjectRequest is the request type for the
// CreateProject API operation.
type CreateProjectRequest struct {
	*aws.Request
	Input *CreateProjectInput
	Copy  func(*CreateProjectInput) CreateProjectRequest
}

// Send marshals and sends the CreateProject API request.
func (r CreateProjectRequest) Send(ctx context.Context) (*CreateProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProjectResponse{
		CreateProjectOutput: r.Request.Data.(*CreateProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProjectResponse is the response type for the
// CreateProject API operation.
type CreateProjectResponse struct {
	*CreateProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProject request.
func (r *CreateProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}