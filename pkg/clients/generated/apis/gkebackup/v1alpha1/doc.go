// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Package v1alpha1 contains API Schema definitions for the gkebackup v1alpha1 API group.
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/pkg/apis/gkebackup
// +k8s:defaulter-gen=TypeMeta
// +groupName=gkebackup.cnrm.cloud.google.com

// Generate deepcopy object for gkebackup/v1alpha1 API group
//
//go:generate go run ../../../../../../scripts/deepcopy-gen/main.go -O zz_generated.deepcopy -i . -h ../../../../../../hack/boilerplate_client_alpha.go.txt
package v1alpha1
