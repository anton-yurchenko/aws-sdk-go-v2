// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the access control list (ACL) of an object. To use this operation, you
// must have READ_ACP access to the object.  <p> <b>Versioning</b> </p> <p>By
// default, GET returns ACL information about the current version of an object. To
// return ACL information about a different version, use the versionId
// subresource.</p> <p>The following operations are related to
// <code>GetObjectAcl</code>:</p> <ul> <li> <p> <a>GetObject</a> </p> </li> <li>
// <p> <a>DeleteObject</a> </p> </li> <li> <p> <a>PutObject</a> </p> </li> </ul>
func (c *Client) GetObjectAcl(ctx context.Context, params *GetObjectAclInput, optFns ...func(*Options)) (*GetObjectAclOutput, error) {
	stack := middleware.NewStack("GetObjectAcl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetObjectAclMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetObjectAclValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectAcl(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetObjectAcl",
			Err:           err,
		}
	}
	out := result.(*GetObjectAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectAclInput struct {
	// VersionId used to reference a specific version of the object.
	VersionId *string
	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
	// The key of the object for which to get the ACL information.
	Key *string
	// The bucket name that contains the object for which to get the ACL information.
	// When using this API with an access point, you must direct requests to the access
	// point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	Bucket *string
}

type GetObjectAclOutput struct {
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged
	// Container for the bucket owner's display name and ID.
	Owner *types.Owner
	// A list of grants.
	Grants []*types.Grant

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetObjectAclMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetObjectAcl{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetObjectAcl{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetObjectAcl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetObjectAcl",
	}
}
