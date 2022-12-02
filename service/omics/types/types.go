// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A read set activation job filter.
type ActivateReadSetFilter struct {

	// The filter's start date.
	CreatedAfter *time.Time

	// The filter's end date.
	CreatedBefore *time.Time

	// The filter's status.
	Status ReadSetActivationJobStatus

	noSmithyDocumentSerde
}

// A read set activation job.
type ActivateReadSetJobItem struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The job's status.
	//
	// This member is required.
	Status ReadSetActivationJobStatus

	// When the job completed.
	CompletionTime *time.Time

	noSmithyDocumentSerde
}

// A source for a read set activation job.
type ActivateReadSetSourceItem struct {

	// The source's read set ID.
	//
	// This member is required.
	ReadSetId *string

	// The source's status.
	//
	// This member is required.
	Status ReadSetActivationJobItemStatus

	// The source's status message.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Details about an imported annotation item.
type AnnotationImportItemDetail struct {

	// The item's job status.
	//
	// This member is required.
	JobStatus JobStatus

	// The source file's location in Amazon S3.
	//
	// This member is required.
	Source *string

	noSmithyDocumentSerde
}

// A source for an annotation import job.
type AnnotationImportItemSource struct {

	// The source file's location in Amazon S3.
	//
	// This member is required.
	Source *string

	noSmithyDocumentSerde
}

// An annotation import job.
type AnnotationImportJobItem struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's destination annotation store.
	//
	// This member is required.
	DestinationName *string

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's service role ARN.
	//
	// This member is required.
	RoleArn *string

	// The job's status.
	//
	// This member is required.
	Status JobStatus

	// When the job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// When the job completed.
	CompletionTime *time.Time

	// The job's left normalization setting.
	RunLeftNormalization bool

	noSmithyDocumentSerde
}

// An annotation store.
type AnnotationStoreItem struct {

	// The store's creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The store's description.
	//
	// This member is required.
	Description *string

	// The store's ID.
	//
	// This member is required.
	Id *string

	// The store's name.
	//
	// This member is required.
	Name *string

	// The store's genome reference.
	//
	// This member is required.
	Reference ReferenceItem

	// The store's server-side encryption (SSE) settings.
	//
	// This member is required.
	SseConfig *SseConfig

	// The store's status.
	//
	// This member is required.
	Status StoreStatus

	// The store's status message.
	//
	// This member is required.
	StatusMessage *string

	// The store's ARN.
	//
	// This member is required.
	StoreArn *string

	// The store's file format.
	//
	// This member is required.
	StoreFormat StoreFormat

	// The store's size in bytes.
	//
	// This member is required.
	StoreSizeBytes *int64

	// When the store was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	noSmithyDocumentSerde
}

// A read set.
type ExportReadSet struct {

	// The set's ID.
	//
	// This member is required.
	ReadSetId *string

	noSmithyDocumentSerde
}

// Details about a read set.
type ExportReadSetDetail struct {

	// The set's ID.
	//
	// This member is required.
	Id *string

	// The set's status.
	//
	// This member is required.
	Status ReadSetExportJobItemStatus

	// The set's status message.
	StatusMessage *string

	noSmithyDocumentSerde
}

// An read set export job filter.
type ExportReadSetFilter struct {

	// The filter's start date.
	CreatedAfter *time.Time

	// The filter's end date.
	CreatedBefore *time.Time

	// A status to filter on.
	Status ReadSetExportJobStatus

	noSmithyDocumentSerde
}

// Details about a read set export job.
type ExportReadSetJobDetail struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's destination in Amazon S3.
	//
	// This member is required.
	Destination *string

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The job's status.
	//
	// This member is required.
	Status ReadSetExportJobStatus

	// When the job completed.
	CompletionTime *time.Time

	noSmithyDocumentSerde
}

// Details about a file.
type FileInformation struct {

	// The file's content length.
	ContentLength *int64

	// The file's part size.
	PartSize *int64

	// The file's total parts.
	TotalParts *int32

	noSmithyDocumentSerde
}

// Formatting options for a file.
//
// The following types satisfy this interface:
//
//	FormatOptionsMemberTsvOptions
//	FormatOptionsMemberVcfOptions
type FormatOptions interface {
	isFormatOptions()
}

// Options for a TSV file.
type FormatOptionsMemberTsvOptions struct {
	Value TsvOptions

	noSmithyDocumentSerde
}

func (*FormatOptionsMemberTsvOptions) isFormatOptions() {}

// Options for a VCF file.
type FormatOptionsMemberVcfOptions struct {
	Value VcfOptions

	noSmithyDocumentSerde
}

func (*FormatOptionsMemberVcfOptions) isFormatOptions() {}

// A filter for import read set jobs.
type ImportReadSetFilter struct {

	// The filter's start date.
	CreatedAfter *time.Time

	// The filter's end date.
	CreatedBefore *time.Time

	// A status to filter on.
	Status ReadSetImportJobStatus

	noSmithyDocumentSerde
}

// An import read set job.
type ImportReadSetJobItem struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's service role ARN.
	//
	// This member is required.
	RoleArn *string

	// The job's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The job's status.
	//
	// This member is required.
	Status ReadSetImportJobStatus

	// When the job completed.
	CompletionTime *time.Time

	noSmithyDocumentSerde
}

// A source for an import read set job.
type ImportReadSetSourceItem struct {

	// The source's sample ID.
	//
	// This member is required.
	SampleId *string

	// The source's file type.
	//
	// This member is required.
	SourceFileType FileType

	// The source files' location in Amazon S3.
	//
	// This member is required.
	SourceFiles *SourceFiles

	// The source's status.
	//
	// This member is required.
	Status ReadSetImportJobItemStatus

	// The source's subject ID.
	//
	// This member is required.
	SubjectId *string

	// The source's description.
	Description *string

	// Where the source originated.
	GeneratedFrom *string

	// The source's name.
	Name *string

	// The source's genome reference ARN.
	ReferenceArn *string

	// The source's status message.
	StatusMessage *string

	// The source's tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A filter for import references.
type ImportReferenceFilter struct {

	// The filter's start date.
	CreatedAfter *time.Time

	// The filter's end date.
	CreatedBefore *time.Time

	// A status to filter on.
	Status ReferenceImportJobStatus

	noSmithyDocumentSerde
}

// An import reference job.
type ImportReferenceJobItem struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's reference store ID.
	//
	// This member is required.
	ReferenceStoreId *string

	// The job's service role ARN.
	//
	// This member is required.
	RoleArn *string

	// The job's status.
	//
	// This member is required.
	Status ReferenceImportJobStatus

	// When the job completed.
	CompletionTime *time.Time

	noSmithyDocumentSerde
}

// An genome reference source.
type ImportReferenceSourceItem struct {

	// The source's status.
	//
	// This member is required.
	Status ReferenceImportJobItemStatus

	// The source's description.
	Description *string

	// The source's name.
	Name *string

	// The source file's location in Amazon S3.
	SourceFile *string

	// The source's status message.
	StatusMessage *string

	// The source's tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A filter for annotation import jobs.
type ListAnnotationImportJobsFilter struct {

	// A status to filter on.
	Status JobStatus

	// A store name to filter on.
	StoreName *string

	noSmithyDocumentSerde
}

// A filter for annotation stores.
type ListAnnotationStoresFilter struct {

	// A status to filter on.
	Status StoreStatus

	noSmithyDocumentSerde
}

// A filter for variant import jobs.
type ListVariantImportJobsFilter struct {

	// A status to filter on.
	Status JobStatus

	// A store name to filter on.
	StoreName *string

	noSmithyDocumentSerde
}

// A filter for variant stores.
type ListVariantStoresFilter struct {

	// A status to filter on.
	Status StoreStatus

	noSmithyDocumentSerde
}

// Read options for an annotation import job.
type ReadOptions struct {

	// The file's comment character.
	Comment *string

	// The file's encoding.
	Encoding *string

	// A character for escaping quotes in the file.
	Escape *string

	// Whether quotes need to be escaped in the file.
	EscapeQuotes bool

	// Whether the file has a header row.
	Header bool

	// A line separator for the file.
	LineSep *string

	// The file's quote character.
	Quote *string

	// Whether all values need to be quoted, or just those that contain quotes.
	QuoteAll bool

	// The file's field separator.
	Sep *string

	noSmithyDocumentSerde
}

// An error from a batch read set operation.
type ReadSetBatchError struct {

	// The error's code.
	//
	// This member is required.
	Code *string

	// The error's ID.
	//
	// This member is required.
	Id *string

	// The error's message.
	//
	// This member is required.
	Message *string

	noSmithyDocumentSerde
}

// Files in a read set.
type ReadSetFiles struct {

	// The files' index.
	Index *FileInformation

	// The location of the first file in Amazon S3.
	Source1 *FileInformation

	// The location of the second file in Amazon S3.
	Source2 *FileInformation

	noSmithyDocumentSerde
}

// A filter for read sets.
type ReadSetFilter struct {

	// The filter's start date.
	CreatedAfter *time.Time

	// The filter's end date.
	CreatedBefore *time.Time

	// A name to filter on.
	Name *string

	// A genome reference ARN to filter on.
	ReferenceArn *string

	// A status to filter on.
	Status ReadSetStatus

	noSmithyDocumentSerde
}

// A read set.
type ReadSetListItem struct {

	// The read set's ARN.
	//
	// This member is required.
	Arn *string

	// When the read set was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The read set's file type.
	//
	// This member is required.
	FileType FileType

	// The read set's ID.
	//
	// This member is required.
	Id *string

	// The read set's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The read set's status.
	//
	// This member is required.
	Status ReadSetStatus

	// The read set's description.
	Description *string

	// The read set's name.
	Name *string

	// The read set's genome reference ARN.
	ReferenceArn *string

	// The read set's sample ID.
	SampleId *string

	// Details about a sequence.
	SequenceInformation *SequenceInformation

	// The read set's subject ID.
	SubjectId *string

	noSmithyDocumentSerde
}

// A set of genome reference files.
type ReferenceFiles struct {

	// The files' index.
	Index *FileInformation

	// The source file's location in Amazon S3.
	Source *FileInformation

	noSmithyDocumentSerde
}

// A filter for references.
type ReferenceFilter struct {

	// The filter's start date.
	CreatedAfter *time.Time

	// The filter's end date.
	CreatedBefore *time.Time

	// An MD5 checksum to filter on.
	Md5 *string

	// A name to filter on.
	Name *string

	noSmithyDocumentSerde
}

// A genome reference.
//
// The following types satisfy this interface:
//
//	ReferenceItemMemberReferenceArn
type ReferenceItem interface {
	isReferenceItem()
}

// The reference's ARN.
type ReferenceItemMemberReferenceArn struct {
	Value string

	noSmithyDocumentSerde
}

func (*ReferenceItemMemberReferenceArn) isReferenceItem() {}

// A genome reference.
type ReferenceListItem struct {

	// The reference's ARN.
	//
	// This member is required.
	Arn *string

	// When the reference was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The reference's ID.
	//
	// This member is required.
	Id *string

	// The reference's MD5 checksum.
	//
	// This member is required.
	Md5 *string

	// The reference's store ID.
	//
	// This member is required.
	ReferenceStoreId *string

	// When the reference was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The reference's description.
	Description *string

	// The reference's name.
	Name *string

	// The reference's status.
	Status ReferenceStatus

	noSmithyDocumentSerde
}

// Details about a reference store.
type ReferenceStoreDetail struct {

	// The store's ARN.
	//
	// This member is required.
	Arn *string

	// When the store was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The store's ID.
	//
	// This member is required.
	Id *string

	// The store's description.
	Description *string

	// The store's name.
	Name *string

	// The store's server-side encryption (SSE) settings.
	SseConfig *SseConfig

	noSmithyDocumentSerde
}

// A filter for reference stores.
type ReferenceStoreFilter struct {

	// The filter's start date.
	CreatedAfter *time.Time

	// The filter's end date.
	CreatedBefore *time.Time

	// The name to filter on.
	Name *string

	noSmithyDocumentSerde
}

// A run group.
type RunGroupListItem struct {

	// The group's ARN.
	Arn *string

	// When the group was created.
	CreationTime *time.Time

	// The group's ID.
	Id *string

	// The group's maximum CPU count setting.
	MaxCpus *int32

	// The group's maximum duration setting.
	MaxDuration *int32

	// The group's maximum concurrent run setting.
	MaxRuns *int32

	// The group's name.
	Name *string

	noSmithyDocumentSerde
}

// A workflow run.
type RunListItem struct {

	// The run's ARN.
	Arn *string

	// When the run was created.
	CreationTime *time.Time

	// The run's ID.
	Id *string

	// The run's name.
	Name *string

	// The run's priority.
	Priority *int32

	// When the run started.
	StartTime *time.Time

	// The run's status.
	Status RunStatus

	// When the run stopped.
	StopTime *time.Time

	// The run's storage capacity.
	StorageCapacity *int32

	// The run's workflow ID.
	WorkflowId *string

	noSmithyDocumentSerde
}

// Details about a sequence.
type SequenceInformation struct {

	// The sequence's alignment setting.
	Alignment *string

	// Where the sequence originated.
	GeneratedFrom *string

	// The sequence's total base count.
	TotalBaseCount *int64

	// The sequence's total read count.
	TotalReadCount *int64

	noSmithyDocumentSerde
}

// Details about a sequence store.
type SequenceStoreDetail struct {

	// The store's ARN.
	//
	// This member is required.
	Arn *string

	// When the store was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The store's ID.
	//
	// This member is required.
	Id *string

	// The store's description.
	Description *string

	// The store's name.
	Name *string

	// The store's server-side encryption (SSE) settings.
	SseConfig *SseConfig

	noSmithyDocumentSerde
}

// A filter for a sequence store.
type SequenceStoreFilter struct {

	// The filter's start date.
	CreatedAfter *time.Time

	// The filter's end date.
	CreatedBefore *time.Time

	// A name to filter on.
	Name *string

	noSmithyDocumentSerde
}

// Source files for a sequence.
type SourceFiles struct {

	// The location of the first file in Amazon S3.
	//
	// This member is required.
	Source1 *string

	// The location of the second file in Amazon S3.
	Source2 *string

	noSmithyDocumentSerde
}

// Server-side encryption (SSE) settings for a store.
type SseConfig struct {

	// The encryption type.
	//
	// This member is required.
	Type EncryptionType

	// An encryption key ARN.
	KeyArn *string

	noSmithyDocumentSerde
}

// A source for a read set activation job.
type StartReadSetActivationJobSourceItem struct {

	// The source's read set ID.
	//
	// This member is required.
	ReadSetId *string

	noSmithyDocumentSerde
}

// A source for a read set import job.
type StartReadSetImportJobSourceItem struct {

	// The source's reference ARN.
	//
	// This member is required.
	ReferenceArn *string

	// The source's sample ID.
	//
	// This member is required.
	SampleId *string

	// The source's file type.
	//
	// This member is required.
	SourceFileType FileType

	// The source files' location in Amazon S3.
	//
	// This member is required.
	SourceFiles *SourceFiles

	// The source's subject ID.
	//
	// This member is required.
	SubjectId *string

	// The source's description.
	Description *string

	// Where the source originated.
	GeneratedFrom *string

	// The source's name.
	Name *string

	// The source's tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A source for a reference import job.
type StartReferenceImportJobSourceItem struct {

	// The source's name.
	//
	// This member is required.
	Name *string

	// The source file's location in Amazon S3.
	//
	// This member is required.
	SourceFile *string

	// The source's description.
	Description *string

	// The source's tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Settings for a store.
//
// The following types satisfy this interface:
//
//	StoreOptionsMemberTsvStoreOptions
type StoreOptions interface {
	isStoreOptions()
}

// File settings for a TSV store.
type StoreOptionsMemberTsvStoreOptions struct {
	Value TsvStoreOptions

	noSmithyDocumentSerde
}

func (*StoreOptionsMemberTsvStoreOptions) isStoreOptions() {}

// A workflow run task.
type TaskListItem struct {

	// The task's CPU count.
	Cpus *int32

	// When the task was created.
	CreationTime *time.Time

	// The task's memory.
	Memory *int32

	// The task's name.
	Name *string

	// When the task started.
	StartTime *time.Time

	// The task's status.
	Status TaskStatus

	// When the task stopped.
	StopTime *time.Time

	// The task's ID.
	TaskId *string

	noSmithyDocumentSerde
}

// Formatting options for a TSV file.
type TsvOptions struct {

	// The file's read options.
	ReadOptions *ReadOptions

	noSmithyDocumentSerde
}

// File settings for a TSV store.
type TsvStoreOptions struct {

	// The store's annotation type.
	AnnotationType AnnotationType

	// The store's header key to column name mapping.
	FormatToHeader map[string]string

	// The store's schema.
	Schema []map[string]SchemaValueType

	noSmithyDocumentSerde
}

// Details about an imported variant item.
type VariantImportItemDetail struct {

	// The item's job status.
	//
	// This member is required.
	JobStatus JobStatus

	// The source file's location in Amazon S3.
	//
	// This member is required.
	Source *string

	noSmithyDocumentSerde
}

// A imported variant item's source.
type VariantImportItemSource struct {

	// The source file's location in Amazon S3.
	//
	// This member is required.
	Source *string

	noSmithyDocumentSerde
}

// A variant import job.
type VariantImportJobItem struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's destination variant store.
	//
	// This member is required.
	DestinationName *string

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's service role ARN.
	//
	// This member is required.
	RoleArn *string

	// The job's status.
	//
	// This member is required.
	Status JobStatus

	// When the job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// When the job completed.
	CompletionTime *time.Time

	// The job's left normalization setting.
	RunLeftNormalization bool

	noSmithyDocumentSerde
}

// A variant store.
type VariantStoreItem struct {

	// When the store was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The store's description.
	//
	// This member is required.
	Description *string

	// The store's ID.
	//
	// This member is required.
	Id *string

	// The store's name.
	//
	// This member is required.
	Name *string

	// The store's genome reference.
	//
	// This member is required.
	Reference ReferenceItem

	// The store's server-side encryption (SSE) settings.
	//
	// This member is required.
	SseConfig *SseConfig

	// The store's status.
	//
	// This member is required.
	Status StoreStatus

	// The store's status message.
	//
	// This member is required.
	StatusMessage *string

	// The store's ARN.
	//
	// This member is required.
	StoreArn *string

	// The store's size in bytes.
	//
	// This member is required.
	StoreSizeBytes *int64

	// When the store was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	noSmithyDocumentSerde
}

// Formatting options for a VCF file.
type VcfOptions struct {

	// The file's ignore filter field setting.
	IgnoreFilterField *bool

	// The file's ignore qual field setting.
	IgnoreQualField *bool

	noSmithyDocumentSerde
}

// A workflow.
type WorkflowListItem struct {

	// The workflow's ARN.
	Arn *string

	// When the workflow was created.
	CreationTime *time.Time

	// The workflow's digest.
	Digest *string

	// The workflow's ID.
	Id *string

	// The workflow's name.
	Name *string

	// The workflow's status.
	Status WorkflowStatus

	// The workflow's type.
	Type WorkflowType

	noSmithyDocumentSerde
}

// A workflow parameter.
type WorkflowParameter struct {

	// The parameter's description.
	Description *string

	// Whether the parameter is optional.
	Optional *bool

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isFormatOptions() {}
func (*UnknownUnionMember) isReferenceItem() {}
func (*UnknownUnionMember) isStoreOptions()  {}