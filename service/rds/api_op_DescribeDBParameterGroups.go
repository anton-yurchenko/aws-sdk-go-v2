// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBParameterGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific DB parameter group to return details for.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBClusterParameterGroup.
	DBParameterGroupName *string `type:"string"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBParameterGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDBParameterGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBParameterGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBParameterGroupsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the result of a successful invocation of the DescribeDBParameterGroups
// action.
type DescribeDBParameterGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of DBParameterGroup instances.
	DBParameterGroups []DBParameterGroup `locationNameList:"DBParameterGroup" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBParameterGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBParameterGroups = "DescribeDBParameterGroups"

// DescribeDBParameterGroupsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns a list of DBParameterGroup descriptions. If a DBParameterGroupName
// is specified, the list will contain only the description of the specified
// DB parameter group.
//
//    // Example sending a request using DescribeDBParameterGroupsRequest.
//    req := client.DescribeDBParameterGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBParameterGroups
func (c *Client) DescribeDBParameterGroupsRequest(input *DescribeDBParameterGroupsInput) DescribeDBParameterGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBParameterGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBParameterGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeDBParameterGroupsOutput{})
	return DescribeDBParameterGroupsRequest{Request: req, Input: input, Copy: c.DescribeDBParameterGroupsRequest}
}

// DescribeDBParameterGroupsRequest is the request type for the
// DescribeDBParameterGroups API operation.
type DescribeDBParameterGroupsRequest struct {
	*aws.Request
	Input *DescribeDBParameterGroupsInput
	Copy  func(*DescribeDBParameterGroupsInput) DescribeDBParameterGroupsRequest
}

// Send marshals and sends the DescribeDBParameterGroups API request.
func (r DescribeDBParameterGroupsRequest) Send(ctx context.Context) (*DescribeDBParameterGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBParameterGroupsResponse{
		DescribeDBParameterGroupsOutput: r.Request.Data.(*DescribeDBParameterGroupsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBParameterGroupsRequestPaginator returns a paginator for DescribeDBParameterGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBParameterGroupsRequest(input)
//   p := rds.NewDescribeDBParameterGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBParameterGroupsPaginator(req DescribeDBParameterGroupsRequest) DescribeDBParameterGroupsPaginator {
	return DescribeDBParameterGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBParameterGroupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeDBParameterGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBParameterGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeDBParameterGroupsPaginator) CurrentPage() *DescribeDBParameterGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeDBParameterGroupsOutput)
}

// DescribeDBParameterGroupsResponse is the response type for the
// DescribeDBParameterGroups API operation.
type DescribeDBParameterGroupsResponse struct {
	*DescribeDBParameterGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBParameterGroups request.
func (r *DescribeDBParameterGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}