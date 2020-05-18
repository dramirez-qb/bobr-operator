package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BobRUISpec defines the desired state of BobRUI
type BobRUISpec struct {
	// Image to deploy
	Image string `json:"image,omitempty"`
	// Replicas that we need
	Replicas int `json:"replicas,omitempty"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// BobRUIStatus defines the observed state of BobRUI
type BobRUIStatus struct {
	Replicas      int    `json:"replicas,omitempty"`
	LabelSelector string `json:"labelselector,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BobRUI is the Schema for the bobruis API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bobruis,scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=".spec.Replicas",statuspath=".status.Replicas",selectorpath=".status.LabelSelector"
// kubebuilder:printcolumn:name="Spec",type="integer",JSONPath=".spec.cronSpec",description="status of the kind"
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.Replicas",description="The number of jobs launched"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type BobRUI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BobRUISpec   `json:"spec,omitempty"`
	Status BobRUIStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BobRUIList contains a list of BobRUI
type BobRUIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BobRUI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BobRUI{}, &BobRUIList{})
}
