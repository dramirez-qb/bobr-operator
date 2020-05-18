package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BobRCryptoSpec defines the desired state of BobRCrypto
type BobRCryptoSpec struct {
	// Image to deploy
	Image string `json:"image,omitempty"`
	// Replicas that we need
	Replicas int `json:"replicas,omitempty"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// BobRCryptoStatus defines the observed state of BobRCrypto
type BobRCryptoStatus struct {
	Replicas      int    `json:"replicas,omitempty"`
	LabelSelector string `json:"labelselector,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BobRCrypto is the Schema for the bobrcryptos API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bobrcryptos,scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=".spec.Replicas",statuspath=".status.Replicas",selectorpath=".status.LabelSelector"
// kubebuilder:printcolumn:name="Spec",type="integer",JSONPath=".spec.cronSpec",description="status of the kind"
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.Replicas",description="The number of jobs launched"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type BobRCrypto struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BobRCryptoSpec   `json:"spec,omitempty"`
	Status BobRCryptoStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BobRCryptoList contains a list of BobRCrypto
type BobRCryptoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BobRCrypto `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BobRCrypto{}, &BobRCryptoList{})
}
