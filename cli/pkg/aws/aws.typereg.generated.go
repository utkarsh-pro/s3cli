// Code Generated via typereg; DO NOT EDIT

package aws

import (
	"fmt"
	"reflect"

	s3 "github.com/aws/aws-sdk-go/service/s3"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(typedNil interface{}) {
    t := reflect.TypeOf(typedNil).Elem()
    typeRegistry[t.PkgPath()+"."+t.Name()] = t
}

func init() {
    registerType((*s3.GetObjectAclInput)(nil))
    registerType((*s3.GetBucketAnalyticsConfigurationInput)(nil))
    registerType((*s3.GetObjectTaggingOutput)(nil))
    registerType((*s3.PutObjectInput)(nil))
    registerType((*s3.Part)(nil))
    registerType((*s3.DeleteBucketLifecycleInput)(nil))
    registerType((*s3.DeleteBucketOwnershipControlsInput)(nil))
    registerType((*s3.DeleteBucketReplicationInput)(nil))
    registerType((*s3.DeleteBucketAnalyticsConfigurationOutput)(nil))
    registerType((*s3.ListMultipartUploadsInput)(nil))
    registerType((*s3.GetBucketPolicyStatusInput)(nil))
    registerType((*s3.DeleteBucketInventoryConfigurationInput)(nil))
    registerType((*s3.ObjectIdentifier)(nil))
    registerType((*s3.PutBucketVersioningOutput)(nil))
    registerType((*s3.PutObjectTaggingOutput)(nil))
    registerType((*s3.DeletePublicAccessBlockOutput)(nil))
    registerType((*s3.DeleteObjectsInput)(nil))
    registerType((*s3.AnalyticsConfiguration)(nil))
    registerType((*s3.WebsiteConfiguration)(nil))
    registerType((*s3.GetBucketVersioningInput)(nil))
    registerType((*s3.DeleteBucketCorsOutput)(nil))
    registerType((*s3.MetricsConfiguration)(nil))
    registerType((*s3.PutBucketMetricsConfigurationInput)(nil))
    registerType((*s3.GetObjectLockConfigurationOutput)(nil))
    registerType((*s3.Grantee)(nil))
    registerType((*s3.HeadBucketInput)(nil))
    registerType((*s3.ListBucketAnalyticsConfigurationsOutput)(nil))
    registerType((*s3.EndEvent)(nil))
    registerType((*s3.NotificationConfigurationFilter)(nil))
    registerType((*s3.S3)(nil))
    registerType((*s3.PutBucketReplicationInput)(nil))
    registerType((*s3.GetObjectLegalHoldOutput)(nil))
    registerType((*s3.GetObjectInput)(nil))
    registerType((*s3.InventoryEncryption)(nil))
    registerType((*s3.Progress)(nil))
    registerType((*s3.AccessControlTranslation)(nil))
    registerType((*s3.PutBucketCorsOutput)(nil))
    registerType((*s3.GetBucketMetricsConfigurationOutput)(nil))
    registerType((*s3.PutPublicAccessBlockInput)(nil))
    registerType((*s3.InventoryS3BucketDestination)(nil))
    registerType((*s3.AnalyticsS3BucketDestination)(nil))
    registerType((*s3.DeleteBucketIntelligentTieringConfigurationInput)(nil))
    registerType((*s3.MetadataEntry)(nil))
    registerType((*s3.DeleteBucketPolicyOutput)(nil))
    registerType((*s3.ListObjectsInput)(nil))
    registerType((*s3.Object)(nil))
    registerType((*s3.GetObjectLegalHoldInput)(nil))
    registerType((*s3.HeadBucketOutput)(nil))
    registerType((*s3.QueueConfiguration)(nil))
    registerType((*s3.DeleteBucketLifecycleOutput)(nil))
    registerType((*s3.PutBucketPolicyInput)(nil))
    registerType((*s3.ScanRange)(nil))
    registerType((*s3.DeleteObjectsOutput)(nil))
    registerType((*s3.PutBucketLifecycleConfigurationOutput)(nil))
    registerType((*s3.RedirectAllRequestsTo)(nil))
    registerType((*s3.GetBucketAclInput)(nil))
    registerType((*s3.RestoreObjectOutput)(nil))
    registerType((*s3.Delete)(nil))
    registerType((*s3.GetBucketMetricsConfigurationInput)(nil))
    registerType((*s3.CopyObjectInput)(nil))
    registerType((*s3.RecordsEvent)(nil))
    registerType((*s3.PutBucketLifecycleConfigurationInput)(nil))
    registerType((*s3.SSES3)(nil))
    registerType((*s3.ListObjectVersionsOutput)(nil))
    registerType((*s3.CopyObjectOutput)(nil))
    registerType((*s3.DeleteObjectInput)(nil))
    registerType((*s3.IntelligentTieringConfiguration)(nil))
    registerType((*s3.LambdaFunctionConfiguration)(nil))
    registerType((*s3.GetBucketWebsiteInput)(nil))
    registerType((*s3.LoggingEnabled)(nil))
    registerType((*s3.DeleteObjectTaggingInput)(nil))
    registerType((*s3.GetBucketInventoryConfigurationInput)(nil))
    registerType((*s3.CreateMultipartUploadInput)(nil))
    registerType((*s3.AbortMultipartUploadInput)(nil))
    registerType((*s3.GlacierJobParameters)(nil))
    registerType((*s3.Location)(nil))
    registerType((*s3.GetBucketPolicyInput)(nil))
    registerType((*s3.GetBucketPolicyStatusOutput)(nil))
    registerType((*s3.ListObjectVersionsInput)(nil))
    registerType((*s3.PutObjectRetentionInput)(nil))
    registerType((*s3.GetBucketLifecycleInput)(nil))
    registerType((*s3.GetBucketAnalyticsConfigurationOutput)(nil))
    registerType((*s3.PutBucketMetricsConfigurationOutput)(nil))
    registerType((*s3.SseKmsEncryptedObjects)(nil))
    registerType((*s3.CopyPartResult)(nil))
    registerType((*s3.BucketLifecycleConfiguration)(nil))
    registerType((*s3.DefaultRetention)(nil))
    registerType((*s3.GetBucketOwnershipControlsOutput)(nil))
    registerType((*s3.PutObjectRetentionOutput)(nil))
    registerType((*s3.FilterRule)(nil))
    registerType((*s3.IntelligentTieringAndOperator)(nil))
    registerType((*s3.UploadPartCopyInput)(nil))
    registerType((*s3.DeleteBucketWebsiteInput)(nil))
    registerType((*s3.CompletedPart)(nil))
    registerType((*s3.DeleteBucketOwnershipControlsOutput)(nil))
    registerType((*s3.ObjectVersion)(nil))
    registerType((*s3.PutBucketOwnershipControlsInput)(nil))
    registerType((*s3.StorageClassAnalysis)(nil))
    registerType((*s3.PutBucketEncryptionOutput)(nil))
    registerType((*s3.GetBucketAccelerateConfigurationInput)(nil))
    registerType((*s3.DeleteBucketCorsInput)(nil))
    registerType((*s3.ListPartsOutput)(nil))
    registerType((*s3.LifecycleConfiguration)(nil))
    registerType((*s3.ReplicationRuleAndOperator)(nil))
    registerType((*s3.PutBucketReplicationOutput)(nil))
    registerType((*s3.GetObjectTaggingInput)(nil))
    registerType((*s3.LifecycleExpiration)(nil))
    registerType((*s3.PutBucketLoggingInput)(nil))
    registerType((*s3.EventBridgeConfiguration)(nil))
    registerType((*s3.PutBucketNotificationConfigurationOutput)(nil))
    registerType((*s3.PolicyStatus)(nil))
    registerType((*s3.GetBucketVersioningOutput)(nil))
    registerType((*s3.CSVInput)(nil))
    registerType((*s3.CSVOutput)(nil))
    registerType((*s3.DeleteBucketWebsiteOutput)(nil))
    registerType((*s3.PutBucketAccelerateConfigurationInput)(nil))
    registerType((*s3.ListBucketInventoryConfigurationsOutput)(nil))
    registerType((*s3.GetBucketReplicationOutput)(nil))
    registerType((*s3.PutBucketLifecycleOutput)(nil))
    registerType((*s3.Owner)(nil))
    registerType((*s3.InventoryDestination)(nil))
    registerType((*s3.HeadObjectOutput)(nil))
    registerType((*s3.ListObjectsOutput)(nil))
    registerType((*s3.LifecycleRuleAndOperator)(nil))
    registerType((*s3.GetBucketTaggingOutput)(nil))
    registerType((*s3.PutBucketLoggingOutput)(nil))
    registerType((*s3.ServerSideEncryptionRule)(nil))
    registerType((*s3.GetBucketPolicyOutput)(nil))
    registerType((*s3.WriteGetObjectResponseInput)(nil))
    registerType((*s3.GetBucketEncryptionInput)(nil))
    registerType((*s3.DeleteBucketEncryptionOutput)(nil))
    registerType((*s3.ListBucketsInput)(nil))
    registerType((*s3.GetBucketInventoryConfigurationOutput)(nil))
    registerType((*s3.SelectObjectContentEventStream)(nil))
    registerType((*s3.GetBucketLocationInput)(nil))
    registerType((*s3.NoncurrentVersionTransition)(nil))
    registerType((*s3.Bucket)(nil))
    registerType((*s3.PutBucketIntelligentTieringConfigurationOutput)(nil))
    registerType((*s3.ExistingObjectReplication)(nil))
    registerType((*s3.InputSerialization)(nil))
    registerType((*s3.CloudFunctionConfiguration)(nil))
    registerType((*s3.ServerSideEncryptionConfiguration)(nil))
    registerType((*s3.PublicAccessBlockConfiguration)(nil))
    registerType((*s3.ReplicationTime)(nil))
    registerType((*s3.DeleteObjectOutput)(nil))
    registerType((*s3.ListBucketIntelligentTieringConfigurationsOutput)(nil))
    registerType((*s3.PutBucketAnalyticsConfigurationInput)(nil))
    registerType((*s3.PutBucketCorsInput)(nil))
    registerType((*s3.Grant)(nil))
    registerType((*s3.KeyFilter)(nil))
    registerType((*s3.PutObjectOutput)(nil))
    registerType((*s3.Metrics)(nil))
    registerType((*s3.Error)(nil))
    registerType((*s3.Tag)(nil))
    registerType((*s3.DeleteBucketEncryptionInput)(nil))
    registerType((*s3.GetBucketLoggingInput)(nil))
    registerType((*s3.PutBucketTaggingInput)(nil))
    registerType((*s3.SelectObjectContentEventStreamUnknownEvent)(nil))
    registerType((*s3.GetBucketRequestPaymentOutput)(nil))
    registerType((*s3.CreateMultipartUploadOutput)(nil))
    registerType((*s3.JSONOutput)(nil))
    registerType((*s3.AbortIncompleteMultipartUpload)(nil))
    registerType((*s3.PutBucketNotificationOutput)(nil))
    registerType((*s3.GetObjectTorrentOutput)(nil))
    registerType((*s3.PutBucketLifecycleInput)(nil))
    registerType((*s3.CommonPrefix)(nil))
    registerType((*s3.GetObjectAttributesOutput)(nil))
    registerType((*s3.Condition)(nil))
    registerType((*s3.GetBucketIntelligentTieringConfigurationInput)(nil))
    registerType((*s3.GetBucketLifecycleConfigurationOutput)(nil))
    registerType((*s3.ListBucketsOutput)(nil))
    registerType((*s3.ContinuationEvent)(nil))
    registerType((*s3.GetObjectRetentionOutput)(nil))
    registerType((*s3.DeleteBucketOutput)(nil))
    registerType((*s3.TopicConfigurationDeprecated)(nil))
    registerType((*s3.DeletePublicAccessBlockInput)(nil))
    registerType((*s3.AccelerateConfiguration)(nil))
    registerType((*s3.CopyObjectResult)(nil))
    registerType((*s3.DeleteBucketMetricsConfigurationOutput)(nil))
    registerType((*s3.PutBucketWebsiteOutput)(nil))
    registerType((*s3.TopicConfiguration)(nil))
    registerType((*s3.PutBucketNotificationInput)(nil))
    registerType((*s3.LifecycleRule)(nil))
    registerType((*s3.ReplicationTimeValue)(nil))
    registerType((*s3.ObjectLockLegalHold)(nil))
    registerType((*s3.ListBucketMetricsConfigurationsOutput)(nil))
    registerType((*s3.UploadPartInput)(nil))
    registerType((*s3.ReplicationRuleFilter)(nil))
    registerType((*s3.InventoryConfiguration)(nil))
    registerType((*s3.PutBucketIntelligentTieringConfigurationInput)(nil))
    registerType((*s3.GetObjectRetentionInput)(nil))
    registerType((*s3.PutBucketInventoryConfigurationOutput)(nil))
    registerType((*s3.ProgressEvent)(nil))
    registerType((*s3.MultipartUpload)(nil))
    registerType((*s3.PutBucketNotificationConfigurationInput)(nil))
    registerType((*s3.ListBucketAnalyticsConfigurationsInput)(nil))
    registerType((*s3.DeleteMarkerReplication)(nil))
    registerType((*s3.AnalyticsFilter)(nil))
    registerType((*s3.Encryption)(nil))
    registerType((*s3.VersioningConfiguration)(nil))
    registerType((*s3.ServerSideEncryptionByDefault)(nil))
    registerType((*s3.ReplicationRule)(nil))
    registerType((*s3.ObjectLockRule)(nil))
    registerType((*s3.CreateBucketInput)(nil))
    registerType((*s3.GetObjectLockConfigurationInput)(nil))
    registerType((*s3.BucketLoggingStatus)(nil))
    registerType((*s3.GetBucketLocationOutput)(nil))
    registerType((*s3.PutBucketAclInput)(nil))
    registerType((*s3.Initiator)(nil))
    registerType((*s3.IndexDocument)(nil))
    registerType((*s3.SelectParameters)(nil))
    registerType((*s3.StatsEvent)(nil))
    registerType((*s3.CompleteMultipartUploadInput)(nil))
    registerType((*s3.Stats)(nil))
    registerType((*s3.InventorySchedule)(nil))
    registerType((*s3.GetBucketLifecycleOutput)(nil))
    registerType((*s3.PutObjectLegalHoldInput)(nil))
    registerType((*s3.DeleteBucketMetricsConfigurationInput)(nil))
    registerType((*s3.PutBucketRequestPaymentInput)(nil))
    registerType((*s3.GetBucketTaggingInput)(nil))
    registerType((*s3.PutPublicAccessBlockOutput)(nil))
    registerType((*s3.PutObjectTaggingInput)(nil))
    registerType((*s3.Redirect)(nil))
    registerType((*s3.GetPublicAccessBlockOutput)(nil))
    registerType((*s3.PutBucketPolicyOutput)(nil))
    registerType((*s3.IntelligentTieringFilter)(nil))
    registerType((*s3.GetBucketIntelligentTieringConfigurationOutput)(nil))
    registerType((*s3.GetObjectTorrentInput)(nil))
    registerType((*s3.PutBucketAnalyticsConfigurationOutput)(nil))
    registerType((*s3.ListPartsInput)(nil))
    registerType((*s3.RequestPaymentConfiguration)(nil))
    registerType((*s3.CompletedMultipartUpload)(nil))
    registerType((*s3.PutObjectLegalHoldOutput)(nil))
    registerType((*s3.GetBucketAclOutput)(nil))
    registerType((*s3.DeleteBucketIntelligentTieringConfigurationOutput)(nil))
    registerType((*s3.GetObjectAttributesInput)(nil))
    registerType((*s3.PutBucketInventoryConfigurationInput)(nil))
    registerType((*s3.Tagging)(nil))
    registerType((*s3.PutBucketAclOutput)(nil))
    registerType((*s3.UploadPartCopyOutput)(nil))
    registerType((*s3.NotificationConfiguration)(nil))
    registerType((*s3.CompleteMultipartUploadOutput)(nil))
    registerType((*s3.GetPublicAccessBlockInput)(nil))
    registerType((*s3.AccessControlPolicy)(nil))
    registerType((*s3.PutBucketVersioningInput)(nil))
    registerType((*s3.Tiering)(nil))
    registerType((*s3.ListBucketInventoryConfigurationsInput)(nil))
    registerType((*s3.ObjectLockConfiguration)(nil))
    registerType((*s3.ObjectPart)(nil))
    registerType((*s3.InventoryFilter)(nil))
    registerType((*s3.GetBucketEncryptionOutput)(nil))
    registerType((*s3.AnalyticsAndOperator)(nil))
    registerType((*s3.OutputLocation)(nil))
    registerType((*s3.DeleteBucketTaggingOutput)(nil))
    registerType((*s3.CreateBucketOutput)(nil))
    registerType((*s3.OutputSerialization)(nil))
    registerType((*s3.PutBucketTaggingOutput)(nil))
    registerType((*s3.Checksum)(nil))
    registerType((*s3.OwnershipControls)(nil))
    registerType((*s3.LifecycleRuleFilter)(nil))
    registerType((*s3.DeleteBucketInput)(nil))
    registerType((*s3.DeleteBucketInventoryConfigurationOutput)(nil))
    registerType((*s3.GetObjectOutput)(nil))
    registerType((*s3.GetObjectAclOutput)(nil))
    registerType((*s3.DeleteBucketTaggingInput)(nil))
    registerType((*s3.GetBucketLifecycleConfigurationInput)(nil))
    registerType((*s3.ListBucketMetricsConfigurationsInput)(nil))
    registerType((*s3.CORSConfiguration)(nil))
    registerType((*s3.PutObjectAclInput)(nil))
    registerType((*s3.DeleteMarkerEntry)(nil))
    registerType((*s3.HeadObjectInput)(nil))
    registerType((*s3.SourceSelectionCriteria)(nil))
    registerType((*s3.WriteGetObjectResponseOutput)(nil))
    registerType((*s3.GetBucketOwnershipControlsInput)(nil))
    registerType((*s3.PutObjectAclOutput)(nil))
    registerType((*s3.EncryptionConfiguration)(nil))
    registerType((*s3.RestoreRequest)(nil))
    registerType((*s3.CORSRule)(nil))
    registerType((*s3.PutBucketRequestPaymentOutput)(nil))
    registerType((*s3.UploadPartOutput)(nil))
    registerType((*s3.JSONInput)(nil))
    registerType((*s3.StorageClassAnalysisDataExport)(nil))
    registerType((*s3.PutBucketAccelerateConfigurationOutput)(nil))
    registerType((*s3.ErrorDocument)(nil))
    registerType((*s3.ReplicaModifications)(nil))
    registerType((*s3.GetBucketWebsiteOutput)(nil))
    registerType((*s3.GetBucketRequestPaymentInput)(nil))
    registerType((*s3.AnalyticsExportDestination)(nil))
    registerType((*s3.GetBucketNotificationConfigurationRequest)(nil))
    registerType((*s3.ListBucketIntelligentTieringConfigurationsInput)(nil))
    registerType((*s3.GetBucketLoggingOutput)(nil))
    registerType((*s3.SelectObjectContentEventStreamEvent)(nil))
    registerType((*s3.SelectObjectContentEventStreamReader)(nil))
    registerType((*s3.MetricsAndOperator)(nil))
    registerType((*s3.SelectObjectContentInput)(nil))
    registerType((*s3.GetObjectAttributesParts)(nil))
    registerType((*s3.GetBucketCorsInput)(nil))
    registerType((*s3.TargetGrant)(nil))
    registerType((*s3.Rule)(nil))
    registerType((*s3.GetBucketAccelerateConfigurationOutput)(nil))
    registerType((*s3.PutBucketEncryptionInput)(nil))
    registerType((*s3.CreateBucketConfiguration)(nil))
    registerType((*s3.NoncurrentVersionExpiration)(nil))
    registerType((*s3.Destination)(nil))
    registerType((*s3.DeleteBucketPolicyInput)(nil))
    registerType((*s3.ObjectLockRetention)(nil))
    registerType((*s3.RestoreObjectInput)(nil))
    registerType((*s3.RoutingRule)(nil))
    registerType((*s3.PutBucketOwnershipControlsOutput)(nil))
    registerType((*s3.DeleteBucketReplicationOutput)(nil))
    registerType((*s3.GetBucketCorsOutput)(nil))
    registerType((*s3.Transition)(nil))
    registerType((*s3.AbortMultipartUploadOutput)(nil))
    registerType((*s3.ListObjectsV2Output)(nil))
    registerType((*s3.DeletedObject)(nil))
    registerType((*s3.ParquetInput)(nil))
    registerType((*s3.NotificationConfigurationDeprecated)(nil))
    registerType((*s3.RequestProgress)(nil))
    registerType((*s3.MetricsFilter)(nil))
    registerType((*s3.SSEKMS)(nil))
    registerType((*s3.ListObjectsV2Input)(nil))
    registerType((*s3.PutObjectLockConfigurationInput)(nil))
    registerType((*s3.PutBucketWebsiteInput)(nil))
    registerType((*s3.DeleteObjectTaggingOutput)(nil))
    registerType((*s3.OwnershipControlsRule)(nil))
    registerType((*s3.SelectObjectContentOutput)(nil))
    registerType((*s3.PutObjectLockConfigurationOutput)(nil))
    registerType((*s3.ListMultipartUploadsOutput)(nil))
    registerType((*s3.ReplicationConfiguration)(nil))
    registerType((*s3.QueueConfigurationDeprecated)(nil))
    registerType((*s3.DeleteBucketAnalyticsConfigurationInput)(nil))
    registerType((*s3.GetBucketReplicationInput)(nil))
}

func instance(name string) (interface{}, error) {
	typ, ok := typeRegistry[name]
	if !ok {
		return nil, fmt.Errorf("type not found in the type registry")
	}

    return reflect.New(typ).Elem().Addr().Interface(), nil
}