// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdminStatus string

// Enum values for AdminStatus
const (
	AdminStatusEnabled             AdminStatus = "ENABLED"
	AdminStatusDisablingInProgress AdminStatus = "DISABLING_IN_PROGRESS"
)

// Values returns all known values for AdminStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AdminStatus) Values() []AdminStatus {
	return []AdminStatus{
		"ENABLED",
		"DISABLING_IN_PROGRESS",
	}
}

type AllowListStatusCode string

// Enum values for AllowListStatusCode
const (
	AllowListStatusCodeOk                   AllowListStatusCode = "OK"
	AllowListStatusCodeS3ObjectNotFound     AllowListStatusCode = "S3_OBJECT_NOT_FOUND"
	AllowListStatusCodeS3UserAccessDenied   AllowListStatusCode = "S3_USER_ACCESS_DENIED"
	AllowListStatusCodeS3ObjectAccessDenied AllowListStatusCode = "S3_OBJECT_ACCESS_DENIED"
	AllowListStatusCodeS3Throttled          AllowListStatusCode = "S3_THROTTLED"
	AllowListStatusCodeS3ObjectOversize     AllowListStatusCode = "S3_OBJECT_OVERSIZE"
	AllowListStatusCodeS3ObjectEmpty        AllowListStatusCode = "S3_OBJECT_EMPTY"
	AllowListStatusCodeUnknownError         AllowListStatusCode = "UNKNOWN_ERROR"
)

// Values returns all known values for AllowListStatusCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AllowListStatusCode) Values() []AllowListStatusCode {
	return []AllowListStatusCode{
		"OK",
		"S3_OBJECT_NOT_FOUND",
		"S3_USER_ACCESS_DENIED",
		"S3_OBJECT_ACCESS_DENIED",
		"S3_THROTTLED",
		"S3_OBJECT_OVERSIZE",
		"S3_OBJECT_EMPTY",
		"UNKNOWN_ERROR",
	}
}

type AllowsUnencryptedObjectUploads string

// Enum values for AllowsUnencryptedObjectUploads
const (
	AllowsUnencryptedObjectUploadsTrue    AllowsUnencryptedObjectUploads = "TRUE"
	AllowsUnencryptedObjectUploadsFalse   AllowsUnencryptedObjectUploads = "FALSE"
	AllowsUnencryptedObjectUploadsUnknown AllowsUnencryptedObjectUploads = "UNKNOWN"
)

// Values returns all known values for AllowsUnencryptedObjectUploads. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AllowsUnencryptedObjectUploads) Values() []AllowsUnencryptedObjectUploads {
	return []AllowsUnencryptedObjectUploads{
		"TRUE",
		"FALSE",
		"UNKNOWN",
	}
}

type AutomatedDiscoveryStatus string

// Enum values for AutomatedDiscoveryStatus
const (
	AutomatedDiscoveryStatusEnabled  AutomatedDiscoveryStatus = "ENABLED"
	AutomatedDiscoveryStatusDisabled AutomatedDiscoveryStatus = "DISABLED"
)

// Values returns all known values for AutomatedDiscoveryStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutomatedDiscoveryStatus) Values() []AutomatedDiscoveryStatus {
	return []AutomatedDiscoveryStatus{
		"ENABLED",
		"DISABLED",
	}
}

type AvailabilityCode string

// Enum values for AvailabilityCode
const (
	AvailabilityCodeAvailable   AvailabilityCode = "AVAILABLE"
	AvailabilityCodeUnavailable AvailabilityCode = "UNAVAILABLE"
)

// Values returns all known values for AvailabilityCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AvailabilityCode) Values() []AvailabilityCode {
	return []AvailabilityCode{
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

type BucketMetadataErrorCode string

// Enum values for BucketMetadataErrorCode
const (
	BucketMetadataErrorCodeAccessDenied BucketMetadataErrorCode = "ACCESS_DENIED"
)

// Values returns all known values for BucketMetadataErrorCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BucketMetadataErrorCode) Values() []BucketMetadataErrorCode {
	return []BucketMetadataErrorCode{
		"ACCESS_DENIED",
	}
}

type ClassificationScopeUpdateOperation string

// Enum values for ClassificationScopeUpdateOperation
const (
	ClassificationScopeUpdateOperationAdd     ClassificationScopeUpdateOperation = "ADD"
	ClassificationScopeUpdateOperationReplace ClassificationScopeUpdateOperation = "REPLACE"
	ClassificationScopeUpdateOperationRemove  ClassificationScopeUpdateOperation = "REMOVE"
)

// Values returns all known values for ClassificationScopeUpdateOperation. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ClassificationScopeUpdateOperation) Values() []ClassificationScopeUpdateOperation {
	return []ClassificationScopeUpdateOperation{
		"ADD",
		"REPLACE",
		"REMOVE",
	}
}

type Currency string

// Enum values for Currency
const (
	CurrencyUsd Currency = "USD"
)

// Values returns all known values for Currency. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Currency) Values() []Currency {
	return []Currency{
		"USD",
	}
}

type DataIdentifierSeverity string

// Enum values for DataIdentifierSeverity
const (
	DataIdentifierSeverityLow    DataIdentifierSeverity = "LOW"
	DataIdentifierSeverityMedium DataIdentifierSeverity = "MEDIUM"
	DataIdentifierSeverityHigh   DataIdentifierSeverity = "HIGH"
)

// Values returns all known values for DataIdentifierSeverity. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataIdentifierSeverity) Values() []DataIdentifierSeverity {
	return []DataIdentifierSeverity{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

type DataIdentifierType string

// Enum values for DataIdentifierType
const (
	DataIdentifierTypeCustom  DataIdentifierType = "CUSTOM"
	DataIdentifierTypeManaged DataIdentifierType = "MANAGED"
)

// Values returns all known values for DataIdentifierType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataIdentifierType) Values() []DataIdentifierType {
	return []DataIdentifierType{
		"CUSTOM",
		"MANAGED",
	}
}

type DayOfWeek string

// Enum values for DayOfWeek
const (
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
)

// Values returns all known values for DayOfWeek. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		"SUNDAY",
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
	}
}

type EffectivePermission string

// Enum values for EffectivePermission
const (
	EffectivePermissionPublic    EffectivePermission = "PUBLIC"
	EffectivePermissionNotPublic EffectivePermission = "NOT_PUBLIC"
	EffectivePermissionUnknown   EffectivePermission = "UNKNOWN"
)

// Values returns all known values for EffectivePermission. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EffectivePermission) Values() []EffectivePermission {
	return []EffectivePermission{
		"PUBLIC",
		"NOT_PUBLIC",
		"UNKNOWN",
	}
}

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeNone    EncryptionType = "NONE"
	EncryptionTypeAes256  EncryptionType = "AES256"
	EncryptionTypeAwsKms  EncryptionType = "aws:kms"
	EncryptionTypeUnknown EncryptionType = "UNKNOWN"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"NONE",
		"AES256",
		"aws:kms",
		"UNKNOWN",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeClientError   ErrorCode = "ClientError"
	ErrorCodeInternalError ErrorCode = "InternalError"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"ClientError",
		"InternalError",
	}
}

type FindingActionType string

// Enum values for FindingActionType
const (
	FindingActionTypeAwsApiCall FindingActionType = "AWS_API_CALL"
)

// Values returns all known values for FindingActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingActionType) Values() []FindingActionType {
	return []FindingActionType{
		"AWS_API_CALL",
	}
}

type FindingCategory string

// Enum values for FindingCategory
const (
	FindingCategoryClassification FindingCategory = "CLASSIFICATION"
	FindingCategoryPolicy         FindingCategory = "POLICY"
)

// Values returns all known values for FindingCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingCategory) Values() []FindingCategory {
	return []FindingCategory{
		"CLASSIFICATION",
		"POLICY",
	}
}

type FindingPublishingFrequency string

// Enum values for FindingPublishingFrequency
const (
	FindingPublishingFrequencyFifteenMinutes FindingPublishingFrequency = "FIFTEEN_MINUTES"
	FindingPublishingFrequencyOneHour        FindingPublishingFrequency = "ONE_HOUR"
	FindingPublishingFrequencySixHours       FindingPublishingFrequency = "SIX_HOURS"
)

// Values returns all known values for FindingPublishingFrequency. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (FindingPublishingFrequency) Values() []FindingPublishingFrequency {
	return []FindingPublishingFrequency{
		"FIFTEEN_MINUTES",
		"ONE_HOUR",
		"SIX_HOURS",
	}
}

type FindingsFilterAction string

// Enum values for FindingsFilterAction
const (
	FindingsFilterActionArchive FindingsFilterAction = "ARCHIVE"
	FindingsFilterActionNoop    FindingsFilterAction = "NOOP"
)

// Values returns all known values for FindingsFilterAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingsFilterAction) Values() []FindingsFilterAction {
	return []FindingsFilterAction{
		"ARCHIVE",
		"NOOP",
	}
}

type FindingStatisticsSortAttributeName string

// Enum values for FindingStatisticsSortAttributeName
const (
	FindingStatisticsSortAttributeNameGroupKey FindingStatisticsSortAttributeName = "groupKey"
	FindingStatisticsSortAttributeNameCount    FindingStatisticsSortAttributeName = "count"
)

// Values returns all known values for FindingStatisticsSortAttributeName. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (FindingStatisticsSortAttributeName) Values() []FindingStatisticsSortAttributeName {
	return []FindingStatisticsSortAttributeName{
		"groupKey",
		"count",
	}
}

type FindingType string

// Enum values for FindingType
const (
	FindingTypeSensitiveDataS3ObjectMultiple             FindingType = "SensitiveData:S3Object/Multiple"
	FindingTypeSensitiveDataS3ObjectFinancial            FindingType = "SensitiveData:S3Object/Financial"
	FindingTypeSensitiveDataS3ObjectPersonal             FindingType = "SensitiveData:S3Object/Personal"
	FindingTypeSensitiveDataS3ObjectCredentials          FindingType = "SensitiveData:S3Object/Credentials"
	FindingTypeSensitiveDataS3ObjectCustomIdentifier     FindingType = "SensitiveData:S3Object/CustomIdentifier"
	FindingTypePolicyIAMUserS3BucketPublic               FindingType = "Policy:IAMUser/S3BucketPublic"
	FindingTypePolicyIAMUserS3BucketSharedExternally     FindingType = "Policy:IAMUser/S3BucketSharedExternally"
	FindingTypePolicyIAMUserS3BucketReplicatedExternally FindingType = "Policy:IAMUser/S3BucketReplicatedExternally"
	FindingTypePolicyIAMUserS3BucketEncryptionDisabled   FindingType = "Policy:IAMUser/S3BucketEncryptionDisabled"
	FindingTypePolicyIAMUserS3BlockPublicAccessDisabled  FindingType = "Policy:IAMUser/S3BlockPublicAccessDisabled"
)

// Values returns all known values for FindingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FindingType) Values() []FindingType {
	return []FindingType{
		"SensitiveData:S3Object/Multiple",
		"SensitiveData:S3Object/Financial",
		"SensitiveData:S3Object/Personal",
		"SensitiveData:S3Object/Credentials",
		"SensitiveData:S3Object/CustomIdentifier",
		"Policy:IAMUser/S3BucketPublic",
		"Policy:IAMUser/S3BucketSharedExternally",
		"Policy:IAMUser/S3BucketReplicatedExternally",
		"Policy:IAMUser/S3BucketEncryptionDisabled",
		"Policy:IAMUser/S3BlockPublicAccessDisabled",
	}
}

type GroupBy string

// Enum values for GroupBy
const (
	GroupByResourcesAffectedS3BucketName GroupBy = "resourcesAffected.s3Bucket.name"
	GroupByType                          GroupBy = "type"
	GroupByClassificationDetailsJobId    GroupBy = "classificationDetails.jobId"
	GroupBySeverityDescription           GroupBy = "severity.description"
)

// Values returns all known values for GroupBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (GroupBy) Values() []GroupBy {
	return []GroupBy{
		"resourcesAffected.s3Bucket.name",
		"type",
		"classificationDetails.jobId",
		"severity.description",
	}
}

type IsDefinedInJob string

// Enum values for IsDefinedInJob
const (
	IsDefinedInJobTrue    IsDefinedInJob = "TRUE"
	IsDefinedInJobFalse   IsDefinedInJob = "FALSE"
	IsDefinedInJobUnknown IsDefinedInJob = "UNKNOWN"
)

// Values returns all known values for IsDefinedInJob. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IsDefinedInJob) Values() []IsDefinedInJob {
	return []IsDefinedInJob{
		"TRUE",
		"FALSE",
		"UNKNOWN",
	}
}

type IsMonitoredByJob string

// Enum values for IsMonitoredByJob
const (
	IsMonitoredByJobTrue    IsMonitoredByJob = "TRUE"
	IsMonitoredByJobFalse   IsMonitoredByJob = "FALSE"
	IsMonitoredByJobUnknown IsMonitoredByJob = "UNKNOWN"
)

// Values returns all known values for IsMonitoredByJob. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IsMonitoredByJob) Values() []IsMonitoredByJob {
	return []IsMonitoredByJob{
		"TRUE",
		"FALSE",
		"UNKNOWN",
	}
}

type JobComparator string

// Enum values for JobComparator
const (
	JobComparatorEq         JobComparator = "EQ"
	JobComparatorGt         JobComparator = "GT"
	JobComparatorGte        JobComparator = "GTE"
	JobComparatorLt         JobComparator = "LT"
	JobComparatorLte        JobComparator = "LTE"
	JobComparatorNe         JobComparator = "NE"
	JobComparatorContains   JobComparator = "CONTAINS"
	JobComparatorStartsWith JobComparator = "STARTS_WITH"
)

// Values returns all known values for JobComparator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JobComparator) Values() []JobComparator {
	return []JobComparator{
		"EQ",
		"GT",
		"GTE",
		"LT",
		"LTE",
		"NE",
		"CONTAINS",
		"STARTS_WITH",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusRunning    JobStatus = "RUNNING"
	JobStatusPaused     JobStatus = "PAUSED"
	JobStatusCancelled  JobStatus = "CANCELLED"
	JobStatusComplete   JobStatus = "COMPLETE"
	JobStatusIdle       JobStatus = "IDLE"
	JobStatusUserPaused JobStatus = "USER_PAUSED"
)

// Values returns all known values for JobStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"RUNNING",
		"PAUSED",
		"CANCELLED",
		"COMPLETE",
		"IDLE",
		"USER_PAUSED",
	}
}

type JobType string

// Enum values for JobType
const (
	JobTypeOneTime   JobType = "ONE_TIME"
	JobTypeScheduled JobType = "SCHEDULED"
)

// Values returns all known values for JobType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobType) Values() []JobType {
	return []JobType{
		"ONE_TIME",
		"SCHEDULED",
	}
}

type LastRunErrorStatusCode string

// Enum values for LastRunErrorStatusCode
const (
	LastRunErrorStatusCodeNone  LastRunErrorStatusCode = "NONE"
	LastRunErrorStatusCodeError LastRunErrorStatusCode = "ERROR"
)

// Values returns all known values for LastRunErrorStatusCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LastRunErrorStatusCode) Values() []LastRunErrorStatusCode {
	return []LastRunErrorStatusCode{
		"NONE",
		"ERROR",
	}
}

type ListJobsFilterKey string

// Enum values for ListJobsFilterKey
const (
	ListJobsFilterKeyJobType   ListJobsFilterKey = "jobType"
	ListJobsFilterKeyJobStatus ListJobsFilterKey = "jobStatus"
	ListJobsFilterKeyCreatedAt ListJobsFilterKey = "createdAt"
	ListJobsFilterKeyName      ListJobsFilterKey = "name"
)

// Values returns all known values for ListJobsFilterKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ListJobsFilterKey) Values() []ListJobsFilterKey {
	return []ListJobsFilterKey{
		"jobType",
		"jobStatus",
		"createdAt",
		"name",
	}
}

type ListJobsSortAttributeName string

// Enum values for ListJobsSortAttributeName
const (
	ListJobsSortAttributeNameCreatedAt ListJobsSortAttributeName = "createdAt"
	ListJobsSortAttributeNameJobStatus ListJobsSortAttributeName = "jobStatus"
	ListJobsSortAttributeNameName      ListJobsSortAttributeName = "name"
	ListJobsSortAttributeNameJobType   ListJobsSortAttributeName = "jobType"
)

// Values returns all known values for ListJobsSortAttributeName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListJobsSortAttributeName) Values() []ListJobsSortAttributeName {
	return []ListJobsSortAttributeName{
		"createdAt",
		"jobStatus",
		"name",
		"jobType",
	}
}

type MacieStatus string

// Enum values for MacieStatus
const (
	MacieStatusPaused  MacieStatus = "PAUSED"
	MacieStatusEnabled MacieStatus = "ENABLED"
)

// Values returns all known values for MacieStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MacieStatus) Values() []MacieStatus {
	return []MacieStatus{
		"PAUSED",
		"ENABLED",
	}
}

type ManagedDataIdentifierSelector string

// Enum values for ManagedDataIdentifierSelector
const (
	ManagedDataIdentifierSelectorAll     ManagedDataIdentifierSelector = "ALL"
	ManagedDataIdentifierSelectorExclude ManagedDataIdentifierSelector = "EXCLUDE"
	ManagedDataIdentifierSelectorInclude ManagedDataIdentifierSelector = "INCLUDE"
	ManagedDataIdentifierSelectorNone    ManagedDataIdentifierSelector = "NONE"
)

// Values returns all known values for ManagedDataIdentifierSelector. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ManagedDataIdentifierSelector) Values() []ManagedDataIdentifierSelector {
	return []ManagedDataIdentifierSelector{
		"ALL",
		"EXCLUDE",
		"INCLUDE",
		"NONE",
	}
}

type OrderBy string

// Enum values for OrderBy
const (
	OrderByAsc  OrderBy = "ASC"
	OrderByDesc OrderBy = "DESC"
)

// Values returns all known values for OrderBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OrderBy) Values() []OrderBy {
	return []OrderBy{
		"ASC",
		"DESC",
	}
}

type OriginType string

// Enum values for OriginType
const (
	OriginTypeSensitiveDataDiscoveryJob       OriginType = "SENSITIVE_DATA_DISCOVERY_JOB"
	OriginTypeAutomatedSensitiveDataDiscovery OriginType = "AUTOMATED_SENSITIVE_DATA_DISCOVERY"
)

// Values returns all known values for OriginType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OriginType) Values() []OriginType {
	return []OriginType{
		"SENSITIVE_DATA_DISCOVERY_JOB",
		"AUTOMATED_SENSITIVE_DATA_DISCOVERY",
	}
}

type RelationshipStatus string

// Enum values for RelationshipStatus
const (
	RelationshipStatusEnabled                     RelationshipStatus = "Enabled"
	RelationshipStatusPaused                      RelationshipStatus = "Paused"
	RelationshipStatusInvited                     RelationshipStatus = "Invited"
	RelationshipStatusCreated                     RelationshipStatus = "Created"
	RelationshipStatusRemoved                     RelationshipStatus = "Removed"
	RelationshipStatusResigned                    RelationshipStatus = "Resigned"
	RelationshipStatusEmailVerificationInProgress RelationshipStatus = "EmailVerificationInProgress"
	RelationshipStatusEmailVerificationFailed     RelationshipStatus = "EmailVerificationFailed"
	RelationshipStatusRegionDisabled              RelationshipStatus = "RegionDisabled"
	RelationshipStatusAccountSuspended            RelationshipStatus = "AccountSuspended"
)

// Values returns all known values for RelationshipStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RelationshipStatus) Values() []RelationshipStatus {
	return []RelationshipStatus{
		"Enabled",
		"Paused",
		"Invited",
		"Created",
		"Removed",
		"Resigned",
		"EmailVerificationInProgress",
		"EmailVerificationFailed",
		"RegionDisabled",
		"AccountSuspended",
	}
}

type RevealRequestStatus string

// Enum values for RevealRequestStatus
const (
	RevealRequestStatusSuccess    RevealRequestStatus = "SUCCESS"
	RevealRequestStatusProcessing RevealRequestStatus = "PROCESSING"
	RevealRequestStatusError      RevealRequestStatus = "ERROR"
)

// Values returns all known values for RevealRequestStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RevealRequestStatus) Values() []RevealRequestStatus {
	return []RevealRequestStatus{
		"SUCCESS",
		"PROCESSING",
		"ERROR",
	}
}

type RevealStatus string

// Enum values for RevealStatus
const (
	RevealStatusEnabled  RevealStatus = "ENABLED"
	RevealStatusDisabled RevealStatus = "DISABLED"
)

// Values returns all known values for RevealStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RevealStatus) Values() []RevealStatus {
	return []RevealStatus{
		"ENABLED",
		"DISABLED",
	}
}

type ScopeFilterKey string

// Enum values for ScopeFilterKey
const (
	ScopeFilterKeyObjectExtension        ScopeFilterKey = "OBJECT_EXTENSION"
	ScopeFilterKeyObjectLastModifiedDate ScopeFilterKey = "OBJECT_LAST_MODIFIED_DATE"
	ScopeFilterKeyObjectSize             ScopeFilterKey = "OBJECT_SIZE"
	ScopeFilterKeyObjectKey              ScopeFilterKey = "OBJECT_KEY"
)

// Values returns all known values for ScopeFilterKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScopeFilterKey) Values() []ScopeFilterKey {
	return []ScopeFilterKey{
		"OBJECT_EXTENSION",
		"OBJECT_LAST_MODIFIED_DATE",
		"OBJECT_SIZE",
		"OBJECT_KEY",
	}
}

type SearchResourcesComparator string

// Enum values for SearchResourcesComparator
const (
	SearchResourcesComparatorEq SearchResourcesComparator = "EQ"
	SearchResourcesComparatorNe SearchResourcesComparator = "NE"
)

// Values returns all known values for SearchResourcesComparator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SearchResourcesComparator) Values() []SearchResourcesComparator {
	return []SearchResourcesComparator{
		"EQ",
		"NE",
	}
}

type SearchResourcesSimpleCriterionKey string

// Enum values for SearchResourcesSimpleCriterionKey
const (
	SearchResourcesSimpleCriterionKeyAccountId                   SearchResourcesSimpleCriterionKey = "ACCOUNT_ID"
	SearchResourcesSimpleCriterionKeyS3BucketName                SearchResourcesSimpleCriterionKey = "S3_BUCKET_NAME"
	SearchResourcesSimpleCriterionKeyS3BucketEffectivePermission SearchResourcesSimpleCriterionKey = "S3_BUCKET_EFFECTIVE_PERMISSION"
	SearchResourcesSimpleCriterionKeyS3BucketSharedAccess        SearchResourcesSimpleCriterionKey = "S3_BUCKET_SHARED_ACCESS"
)

// Values returns all known values for SearchResourcesSimpleCriterionKey. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (SearchResourcesSimpleCriterionKey) Values() []SearchResourcesSimpleCriterionKey {
	return []SearchResourcesSimpleCriterionKey{
		"ACCOUNT_ID",
		"S3_BUCKET_NAME",
		"S3_BUCKET_EFFECTIVE_PERMISSION",
		"S3_BUCKET_SHARED_ACCESS",
	}
}

type SearchResourcesSortAttributeName string

// Enum values for SearchResourcesSortAttributeName
const (
	SearchResourcesSortAttributeNameAccountId                 SearchResourcesSortAttributeName = "ACCOUNT_ID"
	SearchResourcesSortAttributeNameResourceName              SearchResourcesSortAttributeName = "RESOURCE_NAME"
	SearchResourcesSortAttributeNameS3ClassifiableObjectCount SearchResourcesSortAttributeName = "S3_CLASSIFIABLE_OBJECT_COUNT"
	SearchResourcesSortAttributeNameS3ClassifiableSizeInBytes SearchResourcesSortAttributeName = "S3_CLASSIFIABLE_SIZE_IN_BYTES"
)

// Values returns all known values for SearchResourcesSortAttributeName. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (SearchResourcesSortAttributeName) Values() []SearchResourcesSortAttributeName {
	return []SearchResourcesSortAttributeName{
		"ACCOUNT_ID",
		"RESOURCE_NAME",
		"S3_CLASSIFIABLE_OBJECT_COUNT",
		"S3_CLASSIFIABLE_SIZE_IN_BYTES",
	}
}

type SensitiveDataItemCategory string

// Enum values for SensitiveDataItemCategory
const (
	SensitiveDataItemCategoryFinancialInformation SensitiveDataItemCategory = "FINANCIAL_INFORMATION"
	SensitiveDataItemCategoryPersonalInformation  SensitiveDataItemCategory = "PERSONAL_INFORMATION"
	SensitiveDataItemCategoryCredentials          SensitiveDataItemCategory = "CREDENTIALS"
	SensitiveDataItemCategoryCustomIdentifier     SensitiveDataItemCategory = "CUSTOM_IDENTIFIER"
)

// Values returns all known values for SensitiveDataItemCategory. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SensitiveDataItemCategory) Values() []SensitiveDataItemCategory {
	return []SensitiveDataItemCategory{
		"FINANCIAL_INFORMATION",
		"PERSONAL_INFORMATION",
		"CREDENTIALS",
		"CUSTOM_IDENTIFIER",
	}
}

type SeverityDescription string

// Enum values for SeverityDescription
const (
	SeverityDescriptionLow    SeverityDescription = "Low"
	SeverityDescriptionMedium SeverityDescription = "Medium"
	SeverityDescriptionHigh   SeverityDescription = "High"
)

// Values returns all known values for SeverityDescription. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SeverityDescription) Values() []SeverityDescription {
	return []SeverityDescription{
		"Low",
		"Medium",
		"High",
	}
}

type SharedAccess string

// Enum values for SharedAccess
const (
	SharedAccessExternal  SharedAccess = "EXTERNAL"
	SharedAccessInternal  SharedAccess = "INTERNAL"
	SharedAccessNotShared SharedAccess = "NOT_SHARED"
	SharedAccessUnknown   SharedAccess = "UNKNOWN"
)

// Values returns all known values for SharedAccess. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SharedAccess) Values() []SharedAccess {
	return []SharedAccess{
		"EXTERNAL",
		"INTERNAL",
		"NOT_SHARED",
		"UNKNOWN",
	}
}

type SimpleCriterionKeyForJob string

// Enum values for SimpleCriterionKeyForJob
const (
	SimpleCriterionKeyForJobAccountId                   SimpleCriterionKeyForJob = "ACCOUNT_ID"
	SimpleCriterionKeyForJobS3BucketName                SimpleCriterionKeyForJob = "S3_BUCKET_NAME"
	SimpleCriterionKeyForJobS3BucketEffectivePermission SimpleCriterionKeyForJob = "S3_BUCKET_EFFECTIVE_PERMISSION"
	SimpleCriterionKeyForJobS3BucketSharedAccess        SimpleCriterionKeyForJob = "S3_BUCKET_SHARED_ACCESS"
)

// Values returns all known values for SimpleCriterionKeyForJob. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SimpleCriterionKeyForJob) Values() []SimpleCriterionKeyForJob {
	return []SimpleCriterionKeyForJob{
		"ACCOUNT_ID",
		"S3_BUCKET_NAME",
		"S3_BUCKET_EFFECTIVE_PERMISSION",
		"S3_BUCKET_SHARED_ACCESS",
	}
}

type StorageClass string

// Enum values for StorageClass
const (
	StorageClassStandard           StorageClass = "STANDARD"
	StorageClassReducedRedundancy  StorageClass = "REDUCED_REDUNDANCY"
	StorageClassStandardIa         StorageClass = "STANDARD_IA"
	StorageClassIntelligentTiering StorageClass = "INTELLIGENT_TIERING"
	StorageClassDeepArchive        StorageClass = "DEEP_ARCHIVE"
	StorageClassOnezoneIa          StorageClass = "ONEZONE_IA"
	StorageClassGlacier            StorageClass = "GLACIER"
	StorageClassGlacierIr          StorageClass = "GLACIER_IR"
	StorageClassOutposts           StorageClass = "OUTPOSTS"
)

// Values returns all known values for StorageClass. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StorageClass) Values() []StorageClass {
	return []StorageClass{
		"STANDARD",
		"REDUCED_REDUNDANCY",
		"STANDARD_IA",
		"INTELLIGENT_TIERING",
		"DEEP_ARCHIVE",
		"ONEZONE_IA",
		"GLACIER",
		"GLACIER_IR",
		"OUTPOSTS",
	}
}

type TagTarget string

// Enum values for TagTarget
const (
	TagTargetS3Object TagTarget = "S3_OBJECT"
)

// Values returns all known values for TagTarget. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TagTarget) Values() []TagTarget {
	return []TagTarget{
		"S3_OBJECT",
	}
}

type TimeRange string

// Enum values for TimeRange
const (
	TimeRangeMonthToDate TimeRange = "MONTH_TO_DATE"
	TimeRangePast30Days  TimeRange = "PAST_30_DAYS"
)

// Values returns all known values for TimeRange. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TimeRange) Values() []TimeRange {
	return []TimeRange{
		"MONTH_TO_DATE",
		"PAST_30_DAYS",
	}
}

type Type string

// Enum values for Type
const (
	TypeNone   Type = "NONE"
	TypeAes256 Type = "AES256"
	TypeAwsKms Type = "aws:kms"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"NONE",
		"AES256",
		"aws:kms",
	}
}

type UnavailabilityReasonCode string

// Enum values for UnavailabilityReasonCode
const (
	UnavailabilityReasonCodeObjectExceedsSizeQuota      UnavailabilityReasonCode = "OBJECT_EXCEEDS_SIZE_QUOTA"
	UnavailabilityReasonCodeUnsupportedObjectType       UnavailabilityReasonCode = "UNSUPPORTED_OBJECT_TYPE"
	UnavailabilityReasonCodeUnsupportedFindingType      UnavailabilityReasonCode = "UNSUPPORTED_FINDING_TYPE"
	UnavailabilityReasonCodeInvalidClassificationResult UnavailabilityReasonCode = "INVALID_CLASSIFICATION_RESULT"
	UnavailabilityReasonCodeObjectUnavailable           UnavailabilityReasonCode = "OBJECT_UNAVAILABLE"
)

// Values returns all known values for UnavailabilityReasonCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UnavailabilityReasonCode) Values() []UnavailabilityReasonCode {
	return []UnavailabilityReasonCode{
		"OBJECT_EXCEEDS_SIZE_QUOTA",
		"UNSUPPORTED_OBJECT_TYPE",
		"UNSUPPORTED_FINDING_TYPE",
		"INVALID_CLASSIFICATION_RESULT",
		"OBJECT_UNAVAILABLE",
	}
}

type Unit string

// Enum values for Unit
const (
	UnitTerabytes Unit = "TERABYTES"
)

// Values returns all known values for Unit. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Unit) Values() []Unit {
	return []Unit{
		"TERABYTES",
	}
}

type UsageStatisticsFilterComparator string

// Enum values for UsageStatisticsFilterComparator
const (
	UsageStatisticsFilterComparatorGt       UsageStatisticsFilterComparator = "GT"
	UsageStatisticsFilterComparatorGte      UsageStatisticsFilterComparator = "GTE"
	UsageStatisticsFilterComparatorLt       UsageStatisticsFilterComparator = "LT"
	UsageStatisticsFilterComparatorLte      UsageStatisticsFilterComparator = "LTE"
	UsageStatisticsFilterComparatorEq       UsageStatisticsFilterComparator = "EQ"
	UsageStatisticsFilterComparatorNe       UsageStatisticsFilterComparator = "NE"
	UsageStatisticsFilterComparatorContains UsageStatisticsFilterComparator = "CONTAINS"
)

// Values returns all known values for UsageStatisticsFilterComparator. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (UsageStatisticsFilterComparator) Values() []UsageStatisticsFilterComparator {
	return []UsageStatisticsFilterComparator{
		"GT",
		"GTE",
		"LT",
		"LTE",
		"EQ",
		"NE",
		"CONTAINS",
	}
}

type UsageStatisticsFilterKey string

// Enum values for UsageStatisticsFilterKey
const (
	UsageStatisticsFilterKeyAccountId          UsageStatisticsFilterKey = "accountId"
	UsageStatisticsFilterKeyServiceLimit       UsageStatisticsFilterKey = "serviceLimit"
	UsageStatisticsFilterKeyFreeTrialStartDate UsageStatisticsFilterKey = "freeTrialStartDate"
	UsageStatisticsFilterKeyTotal              UsageStatisticsFilterKey = "total"
)

// Values returns all known values for UsageStatisticsFilterKey. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageStatisticsFilterKey) Values() []UsageStatisticsFilterKey {
	return []UsageStatisticsFilterKey{
		"accountId",
		"serviceLimit",
		"freeTrialStartDate",
		"total",
	}
}

type UsageStatisticsSortKey string

// Enum values for UsageStatisticsSortKey
const (
	UsageStatisticsSortKeyAccountId          UsageStatisticsSortKey = "accountId"
	UsageStatisticsSortKeyTotal              UsageStatisticsSortKey = "total"
	UsageStatisticsSortKeyServiceLimitValue  UsageStatisticsSortKey = "serviceLimitValue"
	UsageStatisticsSortKeyFreeTrialStartDate UsageStatisticsSortKey = "freeTrialStartDate"
)

// Values returns all known values for UsageStatisticsSortKey. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageStatisticsSortKey) Values() []UsageStatisticsSortKey {
	return []UsageStatisticsSortKey{
		"accountId",
		"total",
		"serviceLimitValue",
		"freeTrialStartDate",
	}
}

type UsageType string

// Enum values for UsageType
const (
	UsageTypeDataInventoryEvaluation         UsageType = "DATA_INVENTORY_EVALUATION"
	UsageTypeSensitiveDataDiscovery          UsageType = "SENSITIVE_DATA_DISCOVERY"
	UsageTypeAutomatedSensitiveDataDiscovery UsageType = "AUTOMATED_SENSITIVE_DATA_DISCOVERY"
	UsageTypeAutomatedObjectMonitoring       UsageType = "AUTOMATED_OBJECT_MONITORING"
)

// Values returns all known values for UsageType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UsageType) Values() []UsageType {
	return []UsageType{
		"DATA_INVENTORY_EVALUATION",
		"SENSITIVE_DATA_DISCOVERY",
		"AUTOMATED_SENSITIVE_DATA_DISCOVERY",
		"AUTOMATED_OBJECT_MONITORING",
	}
}

type UserIdentityType string

// Enum values for UserIdentityType
const (
	UserIdentityTypeAssumedRole   UserIdentityType = "AssumedRole"
	UserIdentityTypeIAMUser       UserIdentityType = "IAMUser"
	UserIdentityTypeFederatedUser UserIdentityType = "FederatedUser"
	UserIdentityTypeRoot          UserIdentityType = "Root"
	UserIdentityTypeAWSAccount    UserIdentityType = "AWSAccount"
	UserIdentityTypeAWSService    UserIdentityType = "AWSService"
)

// Values returns all known values for UserIdentityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UserIdentityType) Values() []UserIdentityType {
	return []UserIdentityType{
		"AssumedRole",
		"IAMUser",
		"FederatedUser",
		"Root",
		"AWSAccount",
		"AWSService",
	}
}
