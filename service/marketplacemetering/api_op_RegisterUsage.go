// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplacemetering

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterUsageInput struct {
	_ struct{} `type:"structure"`

	// (Optional) To scope down the registration to a specific running software
	// instance and guard against replay attacks.
	Nonce *string `type:"string"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// Public Key Version provided by AWS Marketplace
	//
	// PublicKeyVersion is a required field
	PublicKeyVersion *int64 `min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s RegisterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterUsageInput"}

	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductCode", 1))
	}

	if s.PublicKeyVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicKeyVersion"))
	}
	if s.PublicKeyVersion != nil && *s.PublicKeyVersion < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PublicKeyVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterUsageOutput struct {
	_ struct{} `type:"structure"`

	// (Optional) Only included when public key version has expired
	PublicKeyRotationTimestamp *time.Time `type:"timestamp"`

	// JWT Token
	Signature *string `type:"string"`
}

// String returns the string representation
func (s RegisterUsageOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterUsage = "RegisterUsage"

// RegisterUsageRequest returns a request value for making API operation for
// AWSMarketplace Metering.
//
// Paid container software products sold through AWS Marketplace must integrate
// with the AWS Marketplace Metering Service and call the RegisterUsage operation
// for software entitlement and metering. Calling RegisterUsage from containers
// running outside of ECS is not currently supported. Free and BYOL products
// for ECS aren't required to call RegisterUsage, but you may choose to do so
// if you would like to receive usage data in your seller reports. The sections
// below explain the behavior of RegisterUsage. RegisterUsage performs two primary
// functions: metering and entitlement.
//
//    * Entitlement: RegisterUsage allows you to verify that the customer running
//    your paid software is subscribed to your product on AWS Marketplace, enabling
//    you to guard against unauthorized use. Your container image that integrates
//    with RegisterUsage is only required to guard against unauthorized use
//    at container startup, as such a CustomerNotSubscribedException/PlatformNotSupportedException
//    will only be thrown on the initial call to RegisterUsage. Subsequent calls
//    from the same Amazon ECS task instance (e.g. task-id) will not throw a
//    CustomerNotSubscribedException, even if the customer unsubscribes while
//    the Amazon ECS task is still running.
//
//    * Metering: RegisterUsage meters software use per ECS task, per hour,
//    with usage prorated to the second. A minimum of 1 minute of usage applies
//    to tasks that are short lived. For example, if a customer has a 10 node
//    ECS cluster and creates an ECS service configured as a Daemon Set, then
//    ECS will launch a task on all 10 cluster nodes and the customer will be
//    charged: (10 * hourly_rate). Metering for software use is automatically
//    handled by the AWS Marketplace Metering Control Plane -- your software
//    is not required to perform any metering specific actions, other than call
//    RegisterUsage once for metering of software use to commence. The AWS Marketplace
//    Metering Control Plane will also continue to bill customers for running
//    ECS tasks, regardless of the customers subscription state, removing the
//    need for your software to perform entitlement checks at runtime.
//
//    // Example sending a request using RegisterUsageRequest.
//    req := client.RegisterUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/RegisterUsage
func (c *Client) RegisterUsageRequest(input *RegisterUsageInput) RegisterUsageRequest {
	op := &aws.Operation{
		Name:       opRegisterUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterUsageInput{}
	}

	req := c.newRequest(op, input, &RegisterUsageOutput{})
	return RegisterUsageRequest{Request: req, Input: input, Copy: c.RegisterUsageRequest}
}

// RegisterUsageRequest is the request type for the
// RegisterUsage API operation.
type RegisterUsageRequest struct {
	*aws.Request
	Input *RegisterUsageInput
	Copy  func(*RegisterUsageInput) RegisterUsageRequest
}

// Send marshals and sends the RegisterUsage API request.
func (r RegisterUsageRequest) Send(ctx context.Context) (*RegisterUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterUsageResponse{
		RegisterUsageOutput: r.Request.Data.(*RegisterUsageOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterUsageResponse is the response type for the
// RegisterUsage API operation.
type RegisterUsageResponse struct {
	*RegisterUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterUsage request.
func (r *RegisterUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}