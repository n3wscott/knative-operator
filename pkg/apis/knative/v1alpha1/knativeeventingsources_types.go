package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KnativeEventingSourcesSpec defines the desired state of KnativeEventingSources
type KnativeEventingSourcesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// KnativeEventingSourcesStatus defines the observed state of KnativeEventingSources
type KnativeEventingSourcesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KnativeEventingSources is the Schema for the knativeeventingsources API
// +k8s:openapi-gen=true
type KnativeEventingSources struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KnativeEventingSourcesSpec   `json:"spec,omitempty"`
	Status KnativeEventingSourcesStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KnativeEventingSourcesList contains a list of KnativeEventingSources
type KnativeEventingSourcesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KnativeEventingSources `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KnativeEventingSources{}, &KnativeEventingSourcesList{})
}
