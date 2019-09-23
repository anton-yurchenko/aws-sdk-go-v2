// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StartMonitoringMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs of the GuardDuty member accounts whose findings you
	// want the master account to monitor.
	//
	// AccountIds is a required field
	AccountIds []string `locationName:"accountIds" min:"1" type:"list" required:"true"`

	// The unique ID of the detector of the GuardDuty account whom you want to re-enable
	// to monitor members' findings.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartMonitoringMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartMonitoringMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartMonitoringMembersInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}
	if s.AccountIds != nil && len(s.AccountIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountIds", 1))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartMonitoringMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccountIds != nil {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "accountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartMonitoringMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects containing the unprocessed account and a result string
	// explaining why it was unprocessed.
	//
	// UnprocessedAccounts is a required field
	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list" required:"true"`
}

// String returns the string representation
func (s StartMonitoringMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartMonitoringMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UnprocessedAccounts != nil {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "unprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opStartMonitoringMembers = "StartMonitoringMembers"

// StartMonitoringMembersRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Re-enables GuardDuty to monitor findings of the member accounts specified
// by the account IDs. A master GuardDuty account can run this command after
// disabling GuardDuty from monitoring these members' findings by running StopMonitoringMembers.
//
//    // Example sending a request using StartMonitoringMembersRequest.
//    req := client.StartMonitoringMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/StartMonitoringMembers
func (c *Client) StartMonitoringMembersRequest(input *StartMonitoringMembersInput) StartMonitoringMembersRequest {
	op := &aws.Operation{
		Name:       opStartMonitoringMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/member/start",
	}

	if input == nil {
		input = &StartMonitoringMembersInput{}
	}

	req := c.newRequest(op, input, &StartMonitoringMembersOutput{})
	return StartMonitoringMembersRequest{Request: req, Input: input, Copy: c.StartMonitoringMembersRequest}
}

// StartMonitoringMembersRequest is the request type for the
// StartMonitoringMembers API operation.
type StartMonitoringMembersRequest struct {
	*aws.Request
	Input *StartMonitoringMembersInput
	Copy  func(*StartMonitoringMembersInput) StartMonitoringMembersRequest
}

// Send marshals and sends the StartMonitoringMembers API request.
func (r StartMonitoringMembersRequest) Send(ctx context.Context) (*StartMonitoringMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartMonitoringMembersResponse{
		StartMonitoringMembersOutput: r.Request.Data.(*StartMonitoringMembersOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartMonitoringMembersResponse is the response type for the
// StartMonitoringMembers API operation.
type StartMonitoringMembersResponse struct {
	*StartMonitoringMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartMonitoringMembers request.
func (r *StartMonitoringMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}