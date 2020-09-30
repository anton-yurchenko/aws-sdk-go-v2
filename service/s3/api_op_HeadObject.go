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
	"time"
)

// The HEAD operation retrieves metadata from an object without returning the
// object itself. This operation is useful if you're only interested in an object's
// metadata. To use HEAD, you must have READ access to the object.  <p>A
// <code>HEAD</code> request has the same options as a <code>GET</code> operation
// on an object. The response is identical to the <code>GET</code> response except
// that there is no response body.</p> <p>If you encrypt an object by using
// server-side encryption with customer-provided encryption keys (SSE-C) when you
// store the object in Amazon S3, then when you retrieve the metadata from the
// object, you must use the following headers:</p> <ul> <li>
// <p>x-amz-server-side-encryption-customer-algorithm</p> </li> <li>
// <p>x-amz-server-side-encryption-customer-key</p> </li> <li>
// <p>x-amz-server-side-encryption-customer-key-MD5</p> </li> </ul> <p>For more
// information about SSE-C, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html">Server-Side
// Encryption (Using Customer-Provided Encryption Keys)</a>.</p> <note>
// <p>Encryption request headers, like <code>x-amz-server-side-encryption</code>,
// should not be sent for GET requests if your object uses server-side encryption
// with CMKs stored in AWS KMS (SSE-KMS) or server-side encryption with Amazon
// S3–managed encryption keys (SSE-S3). If your object does use these types of
// keys, you’ll get an HTTP 400 BadRequest error.</p> </note> <p>Request headers
// are limited to 8 KB in size. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonRequestHeaders.html">Common
// Request Headers</a>.</p> <p>Consider the following when using request
// headers:</p> <ul> <li> <p> Consideration 1 – If both of the
// <code>If-Match</code> and <code>If-Unmodified-Since</code> headers are present
// in the request as follows:</p> <ul> <li> <p> <code>If-Match</code> condition
// evaluates to <code>true</code>, and;</p> </li> <li> <p>
// <code>If-Unmodified-Since</code> condition evaluates to <code>false</code>;</p>
// </li> </ul> <p>Then Amazon S3 returns <code>200 OK</code> and the data
// requested.</p> </li> <li> <p> Consideration 2 – If both of the
// <code>If-None-Match</code> and <code>If-Modified-Since</code> headers are
// present in the request as follows:</p> <ul> <li> <p> <code>If-None-Match</code>
// condition evaluates to <code>false</code>, and;</p> </li> <li> <p>
// <code>If-Modified-Since</code> condition evaluates to <code>true</code>;</p>
// </li> </ul> <p>Then Amazon S3 returns the <code>304 Not Modified</code> response
// code.</p> </li> </ul> <p>For more information about conditional requests, see <a
// href="https://tools.ietf.org/html/rfc7232">RFC 7232</a>.</p> <p>
// <b>Permissions</b> </p> <p>You need the <code>s3:GetObject</code> permission for
// this operation. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html">Specifying
// Permissions in a Policy</a>. If the object you request does not exist, the error
// Amazon S3 returns depends on whether you also have the s3:ListBucket
// permission.</p> <ul> <li> <p>If you have the <code>s3:ListBucket</code>
// permission on the bucket, Amazon S3 returns an HTTP status code 404 ("no such
// key") error.</p> </li> <li> <p>If you don’t have the <code>s3:ListBucket</code>
// permission, Amazon S3 returns an HTTP status code 403 ("access denied")
// error.</p> </li> </ul> <p>The following operation is related to
// <code>HeadObject</code>:</p> <ul> <li> <p> <a>GetObject</a> </p> </li> </ul>
func (c *Client) HeadObject(ctx context.Context, params *HeadObjectInput, optFns ...func(*Options)) (*HeadObjectOutput, error) {
	stack := middleware.NewStack("HeadObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpHeadObjectMiddlewares(stack)
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
	addOpHeadObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHeadObject(options.Region), middleware.Before)
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
			OperationName: "HeadObject",
			Err:           err,
		}
	}
	out := result.(*HeadObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HeadObjectInput struct {
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string
	// The name of the bucket containing the object.
	Bucket *string
	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
	// Downloads the specified range bytes of an object. For more information about the
	// HTTP Range header, see
	// http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35 (). Amazon S3
	// doesn't support retrieving multiple ranges of data per GET request.
	Range *string
	// Return the object only if it has not been modified since the specified time,
	// otherwise return a 412 (precondition failed).
	IfUnmodifiedSince *time.Time
	// Part number of the object being read. This is a positive integer between 1 and
	// 10,000. Effectively performs a 'ranged' HEAD request for the part specified.
	// Useful querying about the size of the part and the number of parts in this
	// object.
	PartNumber *int32
	// Return the object only if its entity tag (ETag) is different from the one
	// specified, otherwise return a 304 (not modified).
	IfNoneMatch *string
	// The object key.
	Key *string
	// Return the object only if it has been modified since the specified time,
	// otherwise return a 304 (not modified).
	IfModifiedSince *time.Time
	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string
	// VersionId used to reference a specific version of the object.
	VersionId *string
	// Return the object only if its entity tag (ETag) is the same as the one
	// specified, otherwise return a 412 (precondition failed).
	IfMatch *string
	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string
}

type HeadObjectOutput struct {
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string
	// Last modified date of the object
	LastModified *time.Time
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string
	// The count of parts this object has.
	PartsCount *int32
	// Amazon S3 can return this header if your request involves a bucket that is
	// either a source or destination in a replication rule.  <p>In replication, you
	// have a source bucket on which you configure replication and destination bucket
	// where Amazon S3 stores object replicas. When you request an object
	// (<code>GetObject</code>) or object metadata (<code>HeadObject</code>) from these
	// buckets, Amazon S3 will return the <code>x-amz-replication-status</code> header
	// in the response as follows:</p> <ul> <li> <p>If requesting an object from the
	// source bucket — Amazon S3 will return the <code>x-amz-replication-status</code>
	// header if the object in your request is eligible for replication.</p> <p> For
	// example, suppose that in your replication configuration, you specify object
	// prefix <code>TaxDocs</code> requesting Amazon S3 to replicate objects with key
	// prefix <code>TaxDocs</code>. Any objects you upload with this key name prefix,
	// for example <code>TaxDocs/document1.pdf</code>, are eligible for replication.
	// For any object request with this key name prefix, Amazon S3 will return the
	// <code>x-amz-replication-status</code> header with value PENDING, COMPLETED or
	// FAILED indicating object replication status.</p> </li> <li> <p>If requesting an
	// object from the destination bucket — Amazon S3 will return the
	// <code>x-amz-replication-status</code> header with value REPLICA if the object in
	// your request is a replica that Amazon S3 created.</p> </li> </ul> <p>For more
	// information, see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html">Replication</a>.</p>
	ReplicationStatus types.ReplicationStatus
	// A map of metadata to store with the object in S3.
	Metadata map[string]*string
	// If the object is an archived object (an object whose storage class is GLACIER),
	// the response includes this header if either the archive restoration is in
	// progress (see RestoreObject () or an archive copy is already restored.  <p> If
	// an archive copy is already restored, the header value indicates when Amazon S3
	// is scheduled to delete the object copy. For example:</p> <p>
	// <code>x-amz-restore: ongoing-request="false", expiry-date="Fri, 23 Dec 2012
	// 00:00:00 GMT"</code> </p> <p>If the object restoration is in progress, the
	// header returns the value <code>ongoing-request="true"</code>.</p> <p>For more
	// information about archiving objects, see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-transition-general-considerations">Transitioning
	// Objects: General Considerations</a>.</p>
	Restore *string
	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged
	// Version of the object.
	VersionId *string
	// Specifies caching behavior along the request/reply chain.
	CacheControl *string
	// The date and time when the Object Lock retention period expires. This header is
	// only returned if the requester has the s3:GetObjectRetention permission.
	ObjectLockRetainUntilDate *time.Time
	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata.
	WebsiteRedirectLocation *string
	// Size of the body in bytes.
	ContentLength *int64
	// The date and time at which the object is no longer cacheable.
	Expires *time.Time
	// Specifies whether a legal hold is in effect for this object. This header is only
	// returned if the requester has the s3:GetObjectLegalHold permission. This header
	// is not returned if the specified version of this object has never had a legal
	// hold applied. For more information about S3 Object Lock, see Object Lock
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus
	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP, you
	// can create metadata whose values are not legal HTTP headers.
	MissingMeta *int32
	// Specifies presentational information for the object.
	ContentDisposition *string
	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool
	// The language the content is in.
	ContentLanguage *string
	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL.
	ETag *string
	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key-value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string
	// Indicates that a range of bytes was specified.
	AcceptRanges *string
	// The Object Lock mode, if any, that's in effect for this object. This header is
	// only returned if the requester has the s3:GetObjectRetention permission. For
	// more information about S3 Object Lock, see Object Lock
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockMode types.ObjectLockMode
	// A standard MIME type describing the format of the object data.
	ContentType *string
	// Provides storage class information of the object. Amazon S3 returns this header
	// for all objects except for S3 Standard storage class objects.  <p>For more
	// information, see <a
	// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html">Storage
	// Classes</a>.</p>
	StorageClass types.StorageClass
	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string
	// If the object is stored using server-side encryption either with an AWS KMS
	// customer master key (CMK) or an Amazon S3-managed encryption key, the response
	// includes this header with the value of the server-side encryption algorithm used
	// when storing this object in Amazon S3 (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpHeadObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpHeadObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpHeadObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opHeadObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "HeadObject",
	}
}
