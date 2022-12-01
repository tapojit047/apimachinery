/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	core "k8s.io/api/core/v1"
)

// DeletionPolicy specifies what to do if a resource is deleted
// +kubebuilder:validation:Enum=Delete;WipeOut
type DeletionPolicy string

const (
	DeletionPolicyDelete  DeletionPolicy = "Delete"
	DeletionPolicyWipeOut DeletionPolicy = "WipeOut"
)

type StorageProvider string

const (
	ProviderLocal StorageProvider = "local"
	ProviderS3    StorageProvider = "s3"
	ProviderGCS   StorageProvider = "gcs"
	ProviderAzure StorageProvider = "azure"
	ProviderSwift StorageProvider = "swift"
	ProviderB2    StorageProvider = "b2"
	ProviderRest  StorageProvider = "rest"
)

type Backend struct {
	// Provider specifies the provider of the storage
	Provider StorageProvider `json:"provider,omitempty"`

	// Local specifies the storage information for local provider
	// +optional
	Local *LocalSpec `json:"local,omitempty"`

	// S3 specifies the storage information for AWS S3 and S3 compatible storage.
	// +optional
	S3 *S3Spec `json:"s3,omitempty"`

	// GCS specifies the storage information for GCS bucket
	// +optional
	GCS *GCSSpec `json:"gcs,omitempty"`

	// Azure specifies the storage information for Azure Blob container
	// +optional
	Azure *AzureSpec `json:"azure,omitempty"`

	// Swift specifies the storage information for Swift container
	// +optional
	Swift *SwiftSpec `json:"swift,omitempty"`

	// B2 specifies the storage information for B2 bucket
	// +optional
	B2 *B2Spec `json:"b2,omitempty"`

	// Rest specifies the storage information for rest storage server
	// +optional
	Rest *RestServerSpec `json:"rest,omitempty"`
}

type LocalSpec struct {
	// Represents the source of a volume to mount. Only one of its members may be specified.
	// Make sure the volume exist before using the volume as backend.
	core.VolumeSource `json:",inline"`

	// MountPath specifies the directory where this volume will be mounted
	MountPath string `json:"mountPath,omitempty"`

	// Path within the volume from which the container's volume should be mounted.
	// Defaults to "" (volume's root).
	// +optional
	SubPath string `json:"subPath,omitempty"`
}

type S3Spec struct {
	// Endpoint specifies the URL of the S3 or S3 compatible storage bucket.
	Endpoint string `json:"endpoint,omitempty"`

	// Bucket specifies the name of the bucket that will be used as storage backend.
	Bucket string `json:"bucket,omitempty"`

	// Prefix specifies a directory inside the bucket/container where the data for this backend will be stored.
	Prefix string `json:"prefix,omitempty"`

	// Region specifies the region where the bucket is located
	// +optional
	Region string `json:"region,omitempty"`

	// Secret specifies the name of the Secret that contains the access credential for this storage.
	// +optional
	Secret string `json:"secret,omitempty"`
}

type GCSSpec struct {
	// Bucket specifies the name of the bucket that will be used as storage backend.
	Bucket string `json:"bucket,omitempty"`

	// Prefix specifies a directory inside the bucket/container where the data for this backend will be stored.
	Prefix string `json:"prefix,omitempty"`

	// MaxConnections specifies the maximum number of concurrent connections to use to upload/download data to this backend.
	// +optional
	MaxConnections int64 `json:"maxConnections,omitempty"`

	// Secret specifies the name of the Secret that contains the access credential for this storage.
	// +optional
	Secret string `json:"secret,omitempty"`
}

type AzureSpec struct {
	// Container specifies the name of the Azure Blob container that will be used as storage backend.
	Container string `json:"container,omitempty"`

	// Prefix specifies a directory inside the bucket/container where the data for this backend will be stored.
	Prefix string `json:"prefix,omitempty"`

	// MaxConnections specifies the maximum number of concurrent connections to use to upload/download data to this backend.
	// +optional
	MaxConnections int64 `json:"maxConnections,omitempty"`

	// Secret specifies the name of the Secret that contains the access credential for this storage.
	// +optional
	Secret string `json:"secret,omitempty"`
}

type SwiftSpec struct {
	// Container specifies the name of the Swift container that will be used as storage backend.
	Container string `json:"container,omitempty"`

	// Prefix specifies a directory inside the bucket/container where the data for this backend will be stored.
	Prefix string `json:"prefix,omitempty"`

	// Secret specifies the name of the Secret that contains the access credential for this storage.
	// +optional
	Secret string `json:"secret,omitempty"`
}

type B2Spec struct {
	// Bucket specifies the name of the bucket that will be used as storage backend.
	Bucket string `json:"bucket,omitempty"`

	// Prefix specifies a directory inside the bucket/container where the data for this backend will be stored.
	Prefix string `json:"prefix,omitempty"`

	// MaxConnections specifies the maximum number of concurrent connections to use to upload/download data to this backend.
	// +optional
	MaxConnections int64 `json:"maxConnections,omitempty"`

	// Secret specifies the name of the Secret that contains the access credential for this storage.
	// +optional
	Secret string `json:"secret,omitempty"`
}

type RestServerSpec struct {
	// URL specifies the URL of the REST storage server
	URL string `json:"url,omitempty"`

	// Secret specifies the name of the Secret that contains the access credential for this storage.
	// +optional
	Secret string `json:"secret,omitempty"`
}
