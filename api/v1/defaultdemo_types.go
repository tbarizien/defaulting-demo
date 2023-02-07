/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DefaultDemoSpec defines the desired state of DefaultDemo
type DefaultDemoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DefaultDemo. Edit defaultdemo_types.go to remove/update
	Foo string `json:"foo,omitempty"`

	//+optional
	//+kubebuilder:default:={horsePower:40}
	Car *Car `json:"car,omitempty"`
}

type Car struct {
	//+optional
	//+kubebuilder:default:=50
	Horsepower *int `json:"horsePower,omitempty"`

	//+optional
	//+kubebuilder:default:=4
	Wheels *int `json:"wheels,omitempty"`
}

// DefaultDemoStatus defines the observed state of DefaultDemo
type DefaultDemoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DefaultDemo is the Schema for the defaultdemoes API
type DefaultDemo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DefaultDemoSpec   `json:"spec,omitempty"`
	Status DefaultDemoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DefaultDemoList contains a list of DefaultDemo
type DefaultDemoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultDemo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DefaultDemo{}, &DefaultDemoList{})
}
