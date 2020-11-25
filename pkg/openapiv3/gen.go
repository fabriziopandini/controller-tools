/*
Copyright 2020 The Kubernetes Authors.

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

package openapiv3

import (
	"fmt"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/controller-tools/pkg/crd"
	crdmarkers "sigs.k8s.io/controller-tools/pkg/crd/markers"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

// +controllertools:marker:generateHelp

// Generator outputs OpenAPI validation spec for the CRDs given as input.
type Generator struct {
	// AllowDangerousTypes allows types which are usually omitted from CRD generation
	// because they are not recommended.
	//
	// Currently the following additional types are allowed when this is true:
	// float32
	// float64
	//
	// Left unspecified, the default is false
	AllowDangerousTypes *bool `marker:",optional"`

	// MaxDescLen specifies the maximum description length for fields in CRD's OpenAPI schema.
	//
	// 0 indicates drop the description for all fields completely.
	// n indicates limit the description to at most n characters and truncate the description to
	// closest sentence boundary if it exceeds n characters.
	MaxDescLen *int `marker:",optional"`
}

var _ genall.Generator = &Generator{}

func (Generator) CheckFilter() loader.NodeFilter {
	return crd.FilterTypesForCRDs
}

func (Generator) RegisterMarkers(into *markers.Registry) error {
	return crdmarkers.Register(into)
}

func (g Generator) Generate(ctx *genall.GenerationContext) (result error) {
	// Use the crd generator to get CRDs.
	crds, err := crd.Generator{
		// Forward openapi generator marker to the crd generator.
		AllowDangerousTypes: g.AllowDangerousTypes,
		MaxDescLen:          g.MaxDescLen,

		// Use const for others crd generator markers.

		// Always use v1 for CRD generation
		CRDVersions: []string{"v1"},
		// Trivial version does not applies to v1, set to false (default).
		TrivialVersions: false,
		// PreserveUnknownFields does not impact on the openapi schema, set to nil (default).
		PreserveUnknownFields: nil,
	}.GetCRD(ctx)
	if err != nil {
		return err
	}

	// Extract OpenAPIV3 specification for each crd/Spec.Versions.
	for _, crd := range crds {
		crdObject := crd.Object.(*apiextensionsv1.CustomResourceDefinition)
		for _, version := range crdObject.Spec.Versions {
			out := map[string]interface{}{
				"openapi": "3.0.1",
				"info": map[string]interface{}{
					"title":   fmt.Sprintf("Validation schema for CRD %s/%s, %s", crd.GroupName, crd.Version, crd.PluralName),
					"version": "0.0.0", // TODO: this is a mandatory field, might be we want to use git tag to find a meaningful value to apply here.
				},
				"paths": map[string]interface{}{},
				"components": map[string]interface{}{
					"schemas": map[string]*apiextensionsv1.JSONSchemaProps{
						crdObject.Spec.Names.Kind: version.Schema.OpenAPIV3Schema,
					},
				},
			}

			fileName := fmt.Sprintf("%s_%s.%s.openapi.yaml", crd.GroupName, crd.PluralName, version.Name)
			if err := ctx.WriteYAML(fileName, out); err != nil {
				return err
			}
		}
	}
	return nil
}
