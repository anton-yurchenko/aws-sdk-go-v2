// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the readme file or descriptive text for a package version. For packages
// that do not contain a readme file, CodeArtifact extracts a description from a
// metadata file. For example, from the  element in the pom.xml file of a Maven
// package. The returned text might contain formatting. For example, it might
// contain formatting for Markdown or reStructuredText.
func (c *Client) GetPackageVersionReadme(ctx context.Context, params *GetPackageVersionReadmeInput, optFns ...func(*Options)) (*GetPackageVersionReadmeOutput, error) {
	stack := middleware.NewStack("GetPackageVersionReadme", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetPackageVersionReadmeMiddlewares(stack)
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
	addOpGetPackageVersionReadmeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPackageVersionReadme(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "GetPackageVersionReadme",
			Err:           err,
		}
	}
	out := result.(*GetPackageVersionReadmeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPackageVersionReadmeInput struct {

	// The name of the package version that contains the requested readme file.
	//
	// This member is required.
	Package *string

	// The repository that contains the package with the requested readme file.
	//
	// This member is required.
	Repository *string

	// A string that contains the package version (for example, 3.5.2).
	//
	// This member is required.
	PackageVersion *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string

	// A format that specifies the type of the package version with the requested
	// readme file. The valid values are:
	//
	//     * npm
	//
	//     * pypi
	//
	//     * maven
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the domain that contains the repository that contains the package
	// version with the requested readme file.
	//
	// This member is required.
	Domain *string
}

type GetPackageVersionReadmeOutput struct {

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string

	// The text of the returned readme file.
	Readme *string

	// The current revision associated with the package version.
	VersionRevision *string

	// The format of the package with the requested readme file. Valid format types
	// are:
	//
	//     * npm
	//
	//     * pypi
	//
	//     * maven
	Format types.PackageFormat

	// The name of the package that contains the returned readme file.
	Package *string

	// The version of the package with the requested readme file.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetPackageVersionReadmeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetPackageVersionReadme{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPackageVersionReadme{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPackageVersionReadme(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "GetPackageVersionReadme",
	}
}