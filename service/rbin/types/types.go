// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Information about a resource tag used to identify resources that are to be
// retained by a Recycle Bin retention rule.
type ResourceTag struct {

	// The tag key.
	//
	// This member is required.
	ResourceTagKey *string

	// The tag value.
	ResourceTagValue *string

	noSmithyDocumentSerde
}

// Information about the retention period for which a retention rule is to retain
// resources.
type RetentionPeriod struct {

	// The unit of time in which the retention period is measured. Currently, only DAYS
	// is supported.
	//
	// This member is required.
	RetentionPeriodUnit RetentionPeriodUnit

	// The period value for which the retention rule is to retain resources. The period
	// is measured using the unit specified for RetentionPeriodUnit.
	//
	// This member is required.
	RetentionPeriodValue *int32

	noSmithyDocumentSerde
}

// Information about a Recycle Bin retention rule.
type RuleSummary struct {

	// The description for the retention rule.
	Description *string

	// The unique ID of the retention rule.
	Identifier *string

	// Information about the retention period for which the retention rule retains
	// resources
	RetentionPeriod *RetentionPeriod

	noSmithyDocumentSerde
}

// Information about the tags assigned to a Recycle Bin retention rule.
type Tag struct {

	// The tag key.
	//
	// This member is required.
	Key *string

	// The tag value.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde