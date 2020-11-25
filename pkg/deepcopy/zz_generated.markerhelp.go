// +build !ignore_autogenerated

/*
Copyright2020 The Kubernetes Authors.

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

// Code generated by helpgen. DO NOT EDIT.

package deepcopy

import (
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func (Generator) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "generates code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"HeaderFile": markers.DetailedHelp{
				Summary: "specifies the header text (e.g. license) to prepend to generated files.",
				Details: "",
			},
			"Year": markers.DetailedHelp{
				Summary: "specifies the year to substitute for \" YEAR\" in the header file.",
				Details: "",
			},
		},
	}
}
