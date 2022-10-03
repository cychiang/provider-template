/*
Copyright 2022 The Crossplane Authors.

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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"reflect"
)

// AwsParameters are the configurable fields of a Certificate.
type AwsParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// AwsObservation are the observable fields of a Certificate.
type AwsObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A AwsSpec defines the desired state of a Certificate.
type AwsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AwsParameters `json:"forProvider"`
}

// A AwsStatus represents the observed state of a Certificate.
type AwsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AwsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Aws is an letsencrypt API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,letsencrypt,aws}
type Aws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AwsSpec   `json:"spec"`
	Status AwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsList contains a list of Certificate
type AwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Aws `json:"items"`
}

// Aws type metadata.
var (
	AwsKind             = reflect.TypeOf(Aws{}).Name()
	AwsGroupKind        = schema.GroupKind{Group: Group, Kind: AwsKind}.String()
	AwsKindAPIVersion   = AwsKind + "." + SchemeGroupVersion.String()
	AwsGroupVersionKind = SchemeGroupVersion.WithKind(AwsKind)
)

func init() {
	SchemeBuilder.Register(&Aws{}, &AwsList{})
}
