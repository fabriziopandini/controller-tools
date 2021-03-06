/*

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

//go:generate ../../../../.run-controller-gen.sh paths=. output:dir=.
// +groupName=bar.example.com
package foo

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FooSpec struct {
	// This tests that defaulted fields are stripped for v1beta1,
	// but not for v1
	// +kubebuilder:default=fooDefaultString
	DefaultedString string `json:"defaultedString"`
}
type FooStatus struct{}

type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooSpec   `json:"spec,omitempty"`
	Status FooStatus `json:"status,omitempty"`
}
