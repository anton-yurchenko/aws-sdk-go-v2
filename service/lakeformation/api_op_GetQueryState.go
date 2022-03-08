// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the state of a query previously submitted. Clients are expected to poll
// GetQueryState to monitor the current state of the planning before retrieving the
// work units. A query state is only visible to the principal that made the initial
// call to StartQueryPlanning.
func (c *Client) GetQueryState(ctx context.Context, params *GetQueryStateInput, optFns ...func(*Options)) (*GetQueryStateOutput, error) {
	if params == nil {
		params = &GetQueryStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueryState", params, optFns, c.addOperationGetQueryStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueryStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryStateInput struct {

	// The ID of the plan query operation.
	//
	// This member is required.
	QueryId *string

	noSmithyDocumentSerde
}

// A structure for the output.
type GetQueryStateOutput struct {

	// The state of a query previously submitted. The possible states are:
	//
	// * PENDING:
	// the query is pending.
	//
	// * WORKUNITS_AVAILABLE: some work units are ready for
	// retrieval and execution.
	//
	// * FINISHED: the query planning finished successfully,
	// and all work units are ready for retrieval and execution.
	//
	// * ERROR: an error
	// occurred with the query, such as an invalid query ID or a backend error.
	//
	// This member is required.
	State types.QueryStateString

	// An error message when the operation fails.
	Error *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueryStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetQueryState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetQueryState{}, middleware.After)
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
	if err = addEndpointPrefix_opGetQueryStateMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetQueryStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryState(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetQueryStateMiddleware struct {
}

func (*endpointPrefix_opGetQueryStateMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetQueryStateMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "query-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetQueryStateMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetQueryStateMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetQueryState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GetQueryState",
	}
}