/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VirtualMachineParameters are the configurable fields of a VirtualMachine.
type VirtualMachineParameters struct {
	ConfigurableField string `json:"configurableField"`
	ID string `json:"id, omitempty"`
}

// VirtualMachineObservation are the observable fields of a VirtualMachine.
type VirtualMachineObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A VirtualMachineSpec defines the desired state of a VirtualMachine.
type VirtualMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualMachineParameters `json:"forProvider, omitempty"`
}

// A VirtualMachineStatus represents the observed state of a VirtualMachine.
type VirtualMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A VirtualMachine is an example API type.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineList contains a list of VirtualMachine
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

// VirtualMachine type metadata.
var (
	VirtualMachineKind             = reflect.TypeOf(VirtualMachine{}).Name()
	VirtualMachineGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualMachineKind}.String()
	VirtualMachineKindAPIVersion   = VirtualMachineKind + "." + SchemeGroupVersion.String()
	VirtualMachineGroupVersionKind = SchemeGroupVersion.WithKind(VirtualMachineKind)
)

func init() {
	SchemeBuilder.Register(&VirtualMachine{}, &VirtualMachineList{})
}
