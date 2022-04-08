// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches an IAM policy to the specified resource. Use this to share a rule group
// across accounts. You must be the owner of the rule group to perform this
// operation. This action is subject to the following restrictions:
//
// * You can
// attach only one policy with each PutPermissionPolicy request.
//
// * The ARN in the
// request must be a valid WAF RuleGroup ARN and the rule group must exist in the
// same Region.
//
// * The user making the request must be the owner of the rule group.
func (c *Client) PutPermissionPolicy(ctx context.Context, params *PutPermissionPolicyInput, optFns ...func(*Options)) (*PutPermissionPolicyOutput, error) {
	if params == nil {
		params = &PutPermissionPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPermissionPolicy", params, optFns, c.addOperationPutPermissionPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPermissionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPermissionPolicyInput struct {

	// The policy to attach to the specified rule group. The policy specifications must
	// conform to the following:
	//
	// * The policy must be composed using IAM Policy
	// version 2012-10-17 or version 2015-01-01.
	//
	// * The policy must include
	// specifications for Effect, Action, and Principal.
	//
	// * Effect must specify
	// Allow.
	//
	// * Action must specify wafv2:CreateWebACL, wafv2:UpdateWebACL, and
	// wafv2:PutFirewallManagerRuleGroups and may optionally specify
	// wafv2:GetRuleGroup. WAF rejects any extra actions or wildcard actions in the
	// policy.
	//
	// * The policy must not include a Resource parameter.
	//
	// For more
	// information, see IAM Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html).
	//
	// This member is required.
	Policy *string

	// The Amazon Resource Name (ARN) of the RuleGroup to which you want to attach the
	// policy.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type PutPermissionPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPermissionPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutPermissionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPermissionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutPermissionPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPermissionPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutPermissionPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "PutPermissionPolicy",
	}
}
