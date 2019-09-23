// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTaskDefinitionFamiliesInput struct {
	_ struct{} `type:"structure"`

	// The familyPrefix is a string that is used to filter the results of ListTaskDefinitionFamilies.
	// If you specify a familyPrefix, only task definition family names that begin
	// with the familyPrefix string are returned.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition family results returned by ListTaskDefinitionFamilies
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitionFamilies request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListTaskDefinitionFamilies returns up to 100 results and a nextToken value
	// if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitionFamilies
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The task definition family status with which to filter the ListTaskDefinitionFamilies
	// results. By default, both ACTIVE and INACTIVE task definition families are
	// listed. If this parameter is set to ACTIVE, only task definition families
	// that have an ACTIVE task definition revision are returned. If this parameter
	// is set to INACTIVE, only task definition families that do not have any ACTIVE
	// task definition revisions are returned. If you paginate the resulting output,
	// be sure to keep the status value constant in each subsequent request.
	Status TaskDefinitionFamilyStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTaskDefinitionFamiliesInput) String() string {
	return awsutil.Prettify(s)
}

type ListTaskDefinitionFamiliesOutput struct {
	_ struct{} `type:"structure"`

	// The list of task definition family names that match the ListTaskDefinitionFamilies
	// request.
	Families []string `locationName:"families" type:"list"`

	// The nextToken value to include in a future ListTaskDefinitionFamilies request.
	// When the results of a ListTaskDefinitionFamilies request exceed maxResults,
	// this value can be used to retrieve the next page of results. This value is
	// null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListTaskDefinitionFamiliesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTaskDefinitionFamilies = "ListTaskDefinitionFamilies"

// ListTaskDefinitionFamiliesRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Returns a list of task definition families that are registered to your account
// (which may include task definition families that no longer have any ACTIVE
// task definition revisions).
//
// You can filter out task definition families that do not contain any ACTIVE
// task definition revisions by setting the status parameter to ACTIVE. You
// can also filter the results with the familyPrefix parameter.
//
//    // Example sending a request using ListTaskDefinitionFamiliesRequest.
//    req := client.ListTaskDefinitionFamiliesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/ListTaskDefinitionFamilies
func (c *Client) ListTaskDefinitionFamiliesRequest(input *ListTaskDefinitionFamiliesInput) ListTaskDefinitionFamiliesRequest {
	op := &aws.Operation{
		Name:       opListTaskDefinitionFamilies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTaskDefinitionFamiliesInput{}
	}

	req := c.newRequest(op, input, &ListTaskDefinitionFamiliesOutput{})
	return ListTaskDefinitionFamiliesRequest{Request: req, Input: input, Copy: c.ListTaskDefinitionFamiliesRequest}
}

// ListTaskDefinitionFamiliesRequest is the request type for the
// ListTaskDefinitionFamilies API operation.
type ListTaskDefinitionFamiliesRequest struct {
	*aws.Request
	Input *ListTaskDefinitionFamiliesInput
	Copy  func(*ListTaskDefinitionFamiliesInput) ListTaskDefinitionFamiliesRequest
}

// Send marshals and sends the ListTaskDefinitionFamilies API request.
func (r ListTaskDefinitionFamiliesRequest) Send(ctx context.Context) (*ListTaskDefinitionFamiliesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTaskDefinitionFamiliesResponse{
		ListTaskDefinitionFamiliesOutput: r.Request.Data.(*ListTaskDefinitionFamiliesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTaskDefinitionFamiliesRequestPaginator returns a paginator for ListTaskDefinitionFamilies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTaskDefinitionFamiliesRequest(input)
//   p := ecs.NewListTaskDefinitionFamiliesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTaskDefinitionFamiliesPaginator(req ListTaskDefinitionFamiliesRequest) ListTaskDefinitionFamiliesPaginator {
	return ListTaskDefinitionFamiliesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTaskDefinitionFamiliesInput
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

// ListTaskDefinitionFamiliesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTaskDefinitionFamiliesPaginator struct {
	aws.Pager
}

func (p *ListTaskDefinitionFamiliesPaginator) CurrentPage() *ListTaskDefinitionFamiliesOutput {
	return p.Pager.CurrentPage().(*ListTaskDefinitionFamiliesOutput)
}

// ListTaskDefinitionFamiliesResponse is the response type for the
// ListTaskDefinitionFamilies API operation.
type ListTaskDefinitionFamiliesResponse struct {
	*ListTaskDefinitionFamiliesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTaskDefinitionFamilies request.
func (r *ListTaskDefinitionFamiliesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}