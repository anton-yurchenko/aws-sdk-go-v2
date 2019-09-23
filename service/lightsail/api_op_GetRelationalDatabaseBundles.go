// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRelationalDatabaseBundlesInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to a specific page of results for your get relational
	// database bundles request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetRelationalDatabaseBundlesInput) String() string {
	return awsutil.Prettify(s)
}

type GetRelationalDatabaseBundlesOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your get relational database bundles request.
	Bundles []RelationalDatabaseBundle `locationName:"bundles" type:"list"`

	// A token used for advancing to the next page of results of your get relational
	// database bundles request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetRelationalDatabaseBundlesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRelationalDatabaseBundles = "GetRelationalDatabaseBundles"

// GetRelationalDatabaseBundlesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the list of bundles that are available in Amazon Lightsail. A bundle
// describes the performance specifications for a database.
//
// You can use a bundle ID to create a new database with explicit performance
// specifications.
//
//    // Example sending a request using GetRelationalDatabaseBundlesRequest.
//    req := client.GetRelationalDatabaseBundlesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseBundles
func (c *Client) GetRelationalDatabaseBundlesRequest(input *GetRelationalDatabaseBundlesInput) GetRelationalDatabaseBundlesRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseBundles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRelationalDatabaseBundlesInput{}
	}

	req := c.newRequest(op, input, &GetRelationalDatabaseBundlesOutput{})
	return GetRelationalDatabaseBundlesRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseBundlesRequest}
}

// GetRelationalDatabaseBundlesRequest is the request type for the
// GetRelationalDatabaseBundles API operation.
type GetRelationalDatabaseBundlesRequest struct {
	*aws.Request
	Input *GetRelationalDatabaseBundlesInput
	Copy  func(*GetRelationalDatabaseBundlesInput) GetRelationalDatabaseBundlesRequest
}

// Send marshals and sends the GetRelationalDatabaseBundles API request.
func (r GetRelationalDatabaseBundlesRequest) Send(ctx context.Context) (*GetRelationalDatabaseBundlesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseBundlesResponse{
		GetRelationalDatabaseBundlesOutput: r.Request.Data.(*GetRelationalDatabaseBundlesOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseBundlesResponse is the response type for the
// GetRelationalDatabaseBundles API operation.
type GetRelationalDatabaseBundlesResponse struct {
	*GetRelationalDatabaseBundlesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseBundles request.
func (r *GetRelationalDatabaseBundlesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}