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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"stash.appscode.dev/kubestash/apis"
)

// BackupStorage specifies the backend information where the backed up data of different applications will be stored.
// You can consider BackupStorage as a representation of a bucket in Kubernetes native way.
// This is a namespaced object. However, you can use the BackupStorage from any namespace
// as long as it is permitted by the `.spec.usagePolicy` field.
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
type BackupStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupStorageSpec   `json:"spec,omitempty"`
	Status BackupStorageStatus `json:"status,omitempty"`
}

// BackupStorageSpec defines information regarding remote backend, its access credentials, usage policy etc.
type BackupStorageSpec struct {
	// Storage specifies the remote storage information
	Storage Backend `json:"storage,omitempty"`
	// UsagePolicy specifies a policy of how this BackupStorage will be used. For example, you can use `allowedNamespaces`
	// policy to restrict the usage of this BackupStorage to particular namespaces.
	// This field is optional. If you don't provide the usagePolicy, then it can be used only from the current namespace.
	// +optional
	UsagePolicy *apis.UsagePolicy `json:"usagePolicy,omitempty"`

	// Default specifies whether to use this BackupStorage as default storage for the current namespace
	// as well as the allowed namespaces. One namespace can have at most one default RetentionPolicy configured.
	// +optional
	Default bool `json:"default,omitempty"`

	// DeletionPolicy specifies what to do when you delete a BackupStorage CR.
	// The valid values are:
	// "Delete": This will delete the respective Repository and Snapshot CRs from the cluster but keep the backed up data in the remote backend. This is the default behavior.
	// "WipeOut": This will delete the respective Repository and Snapshot CRs as well as the backed up data from the backend.
	// +kubebuilder:validation:default=Delete
	// +optional
	DeletionPolicy DeletionPolicy `json:"deletionPolicy,omitempty"`
}

// BackupStorageStatus defines the observed state of BackupStorage
type BackupStorageStatus struct {
	// Ready specifies whether the BackupStorage is ready to use or not.
	// +optional
	Ready bool `json:"ready,omitempty"`

	// TotalSize represent the total backed up data size in this storage.
	// This is simply the summation of sizes of all Repositories using this BackupStorage.
	// +optional
	TotalSize string `json:"totalSize,omitempty"`

	// Repositories holds the information of all Repositories using this BackupStorage
	// +optional
	Repositories []RepositoryInfo `json:"repositories,omitempty"`

	// Conditions represents list of conditions regarding this BackupStorage
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// RepositoryInfo specifies information regarding a Repository using the BackupStorage
type RepositoryInfo struct {
	// Name represent the name of the respective Repository CR
	Name string `json:"name,omitempty"`

	// Namespace represent the namespace where the Repository CR has been created
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Path represents the directory inside the BackupStorage where this Repository is storing its data
	// This path is relative to the path of BackupStorage.
	Path string `json:"path,omitempty"`

	// Size represents the size of the backed up data in this Repository
	// +optional
	Size string `json:"size,omitempty"`

	// Synced specifies whether this Repository state has been synced with the cloud state or not
	// +optional
	Synced bool `json:"synced,omitempty"`

	// Error specifies the reason in case of Repository sync failure.
	// +optional
	Error *string `json:"error,omitempty"`
}

//+kubebuilder:object:root=true

// BackupStorageList contains a list of BackupStorage
type BackupStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupStorage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupStorage{}, &BackupStorageList{})
}
