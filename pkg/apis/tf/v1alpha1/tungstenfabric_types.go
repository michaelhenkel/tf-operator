package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TungstenfabricSpec defines the desired state of Tungstenfabric
// +k8s:openapi-gen=true
type TungstenfabricSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	// Size is the size of the memcached deployment
	Size int32 `json:"size"`
}

// TungstenfabricStatus defines the observed state of Tungstenfabric
// +k8s:openapi-gen=true
type TungstenfabricStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	// Nodes are the names of the memcached pods
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Tungstenfabric is the Schema for the tungstenfabrics API
// +k8s:openapi-gen=true
type Tungstenfabric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TungstenfabricSpec   `json:"spec,omitempty"`
	Status TungstenfabricStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TungstenfabricList contains a list of Tungstenfabric
type TungstenfabricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tungstenfabric `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tungstenfabric{}, &TungstenfabricList{})
}
