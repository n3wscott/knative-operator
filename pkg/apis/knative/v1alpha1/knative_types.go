package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*

Planning:

spec.serving
 - version
spec.eventing
 - version
spec.build
 - version


- todo: customDomain



 */


// KnativeSpec defines the desired state of Knative
type KnativeSpec struct {
	Serving ServingSpec `json:"serving,omitempty"`
	Eventing EventingSpec `json:"eventing,omitempty"`
	Build BuildSpec `json:"build,omitempty"`

	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// ServingSpec the spec for Serving specific installation configuration.
type ServingSpec struct {
	Version string `json:"version,omitempty"`
}

// EventingSpec the spec for Eventing specific installation configuration.
type EventingSpec struct {
	Version string `json:"version,omitempty"`
}

// BuildSpec the spec for Build specific installation configuration.
type BuildSpec struct {
	Version string `json:"version,omitempty"`
}


// KnativeStatus defines the observed state of Knative
type KnativeStatus struct {
	// TODO: INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Knative is the Schema for the knatives API
// +k8s:openapi-gen=true
type Knative struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KnativeSpec   `json:"spec,omitempty"`
	Status KnativeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KnativeList contains a list of Knative
type KnativeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Knative `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Knative{}, &KnativeList{})
}
