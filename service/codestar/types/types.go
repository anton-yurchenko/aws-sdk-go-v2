// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Location and destination information about the source code files provided with
// the project request. The source code is uploaded to the new project source
// repository after project creation.
type Code struct {

	// The repository to be created in AWS CodeStar. Valid values are AWS CodeCommit or
	// GitHub. After AWS CodeStar provisions the new repository, the source code files
	// provided with the project request are placed in the repository.
	//
	// This member is required.
	Destination *CodeDestination

	// The location where the source code files provided with the project request are
	// stored. AWS CodeStar retrieves the files during project creation.
	//
	// This member is required.
	Source *CodeSource

	noSmithyDocumentSerde
}

// Information about the AWS CodeCommit repository to be created in AWS CodeStar.
// This is where the source code files provided with the project request will be
// uploaded after project creation.
type CodeCommitCodeDestination struct {

	// The name of the AWS CodeCommit repository to be created in AWS CodeStar.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// The repository to be created in AWS CodeStar. Valid values are AWS CodeCommit or
// GitHub. After AWS CodeStar provisions the new repository, the source code files
// provided with the project request are placed in the repository.
type CodeDestination struct {

	// Information about the AWS CodeCommit repository to be created in AWS CodeStar.
	// This is where the source code files provided with the project request will be
	// uploaded after project creation.
	CodeCommit *CodeCommitCodeDestination

	// Information about the GitHub repository to be created in AWS CodeStar. This is
	// where the source code files provided with the project request will be uploaded
	// after project creation.
	GitHub *GitHubCodeDestination

	noSmithyDocumentSerde
}

// The location where the source code files provided with the project request are
// stored. AWS CodeStar retrieves the files during project creation.
type CodeSource struct {

	// Information about the Amazon S3 location where the source code files provided
	// with the project request are stored.
	//
	// This member is required.
	S3 *S3Location

	noSmithyDocumentSerde
}

// Information about the GitHub repository to be created in AWS CodeStar. This is
// where the source code files provided with the project request will be uploaded
// after project creation.
type GitHubCodeDestination struct {

	// Whether to enable issues for the GitHub repository.
	//
	// This member is required.
	IssuesEnabled bool

	// Name of the GitHub repository to be created in AWS CodeStar.
	//
	// This member is required.
	Name *string

	// The GitHub username for the owner of the GitHub repository to be created in AWS
	// CodeStar. If this repository should be owned by a GitHub organization, provide
	// its name.
	//
	// This member is required.
	Owner *string

	// Whether the GitHub repository is to be a private repository.
	//
	// This member is required.
	PrivateRepository bool

	// The GitHub user's personal access token for the GitHub repository.
	//
	// This member is required.
	Token *string

	// The type of GitHub repository to be created in AWS CodeStar. Valid values are
	// User or Organization.
	//
	// This member is required.
	Type *string

	// Description for the GitHub repository to be created in AWS CodeStar. This
	// description displays in GitHub after the repository is created.
	Description *string

	noSmithyDocumentSerde
}

// An indication of whether a project creation or deletion is failed or successful.
type ProjectStatus struct {

	// The phase of completion for a project creation or deletion.
	//
	// This member is required.
	State *string

	// In the case of a project creation or deletion failure, a reason for the failure.
	Reason *string

	noSmithyDocumentSerde
}

// Information about the metadata for a project.
type ProjectSummary struct {

	// The Amazon Resource Name (ARN) of the project.
	ProjectArn *string

	// The ID of the project.
	ProjectId *string

	noSmithyDocumentSerde
}

// Information about a resource for a project.
type Resource struct {

	// The Amazon Resource Name (ARN) of the resource.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// The Amazon S3 location where the source code files provided with the project
// request are stored.
type S3Location struct {

	// The Amazon S3 object key where the source code files provided with the project
	// request are stored.
	BucketKey *string

	// The Amazon S3 bucket name where the source code files provided with the project
	// request are stored.
	BucketName *string

	noSmithyDocumentSerde
}

// Information about a team member in a project.
type TeamMember struct {

	// The role assigned to the user in the project. Project roles have different
	// levels of access. For more information, see Working with Teams
	// (http://docs.aws.amazon.com/codestar/latest/userguide/working-with-teams.html)
	// in the AWS CodeStar User Guide.
	//
	// This member is required.
	ProjectRole *string

	// The Amazon Resource Name (ARN) of the user in IAM.
	//
	// This member is required.
	UserArn *string

	// Whether the user is allowed to remotely access project resources using an SSH
	// public/private key pair.
	RemoteAccessAllowed *bool

	noSmithyDocumentSerde
}

// The toolchain template file provided with the project request. AWS CodeStar uses
// the template to provision the toolchain stack in AWS CloudFormation.
type Toolchain struct {

	// The Amazon S3 location where the toolchain template file provided with the
	// project request is stored. AWS CodeStar retrieves the file during project
	// creation.
	//
	// This member is required.
	Source *ToolchainSource

	// The service role ARN for AWS CodeStar to use for the toolchain template during
	// stack provisioning.
	RoleArn *string

	// The list of parameter overrides to be passed into the toolchain template during
	// stack provisioning, if any.
	StackParameters map[string]string

	noSmithyDocumentSerde
}

// The Amazon S3 location where the toolchain template file provided with the
// project request is stored. AWS CodeStar retrieves the file during project
// creation.
type ToolchainSource struct {

	// The Amazon S3 bucket where the toolchain template file provided with the project
	// request is stored.
	//
	// This member is required.
	S3 *S3Location

	noSmithyDocumentSerde
}

// Information about a user's profile in AWS CodeStar.
type UserProfileSummary struct {

	// The display name of a user in AWS CodeStar. For example, this could be set to
	// both first and last name ("Mary Major") or a single name ("Mary"). The display
	// name is also used to generate the initial icon associated with the user in AWS
	// CodeStar projects. If spaces are included in the display name, the first
	// character that appears after the space will be used as the second character in
	// the user initial icon. The initial icon displays a maximum of two characters, so
	// a display name with more than one space (for example "Mary Jane Major") would
	// generate an initial icon using the first character and the first character after
	// the space ("MJ", not "MM").
	DisplayName *string

	// The email address associated with the user.
	EmailAddress *string

	// The SSH public key associated with the user in AWS CodeStar. If a project owner
	// allows the user remote access to project resources, this public key will be used
	// along with the user's private key for SSH access.
	SshPublicKey *string

	// The Amazon Resource Name (ARN) of the user in IAM.
	UserArn *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
