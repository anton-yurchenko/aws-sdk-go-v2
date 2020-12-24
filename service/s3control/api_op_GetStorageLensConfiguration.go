// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Gets the Amazon S3 Storage Lens configuration. For more information, see Working
// with Amazon S3 Storage Lens
// (https://docs.aws.amazon.com/https:/docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html)
// in the Amazon Simple Storage Service Developer Guide. To use this action, you
// must have permission to perform the s3:GetStorageLensConfiguration action. For
// more information, see Setting permissions to use Amazon S3 Storage Lens
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html#storage_lens_IAM)
// in the Amazon Simple Storage Service Developer Guide.
func (c *Client) GetStorageLensConfiguration(ctx context.Context, params *GetStorageLensConfigurationInput, optFns ...func(*Options)) (*GetStorageLensConfigurationOutput, error) {
	if params == nil {
		params = &GetStorageLensConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStorageLensConfiguration", params, optFns, addOperationGetStorageLensConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStorageLensConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStorageLensConfigurationInput struct {

	// The account ID of the requester.
	//
	// This member is required.
	AccountId *string

	// The ID of the Amazon S3 Storage Lens configuration.
	//
	// This member is required.
	ConfigId *string
}

type GetStorageLensConfigurationOutput struct {

	// The S3 Storage Lens configuration requested.
	StorageLensConfiguration *types.StorageLensConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetStorageLensConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetStorageLensConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetStorageLensConfiguration{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addEndpointPrefix_opGetStorageLensConfigurationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetStorageLensConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStorageLensConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addGetStorageLensConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetStorageLensConfigurationMiddleware struct {
}

func (*endpointPrefix_opGetStorageLensConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetStorageLensConfigurationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetStorageLensConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetStorageLensConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetStorageLensConfigurationMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetStorageLensConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetStorageLensConfiguration",
	}
}

func copyGetStorageLensConfigurationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetStorageLensConfigurationInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetStorageLensConfigurationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillGetStorageLensConfigurationAccountID(input interface{}, v string) error {
	in := input.(*GetStorageLensConfigurationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetStorageLensConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetStorageLensConfigurationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}