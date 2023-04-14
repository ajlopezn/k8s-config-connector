//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonsconfigAddonsConfig) DeepCopyInto(out *AddonsconfigAddonsConfig) {
	*out = *in
	if in.AdvancedApiOpsConfig != nil {
		in, out := &in.AdvancedApiOpsConfig, &out.AdvancedApiOpsConfig
		*out = new(AddonsconfigAdvancedApiOpsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ApiSecurityConfig != nil {
		in, out := &in.ApiSecurityConfig, &out.ApiSecurityConfig
		*out = new(AddonsconfigApiSecurityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ConnectorsPlatformConfig != nil {
		in, out := &in.ConnectorsPlatformConfig, &out.ConnectorsPlatformConfig
		*out = new(AddonsconfigConnectorsPlatformConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IntegrationConfig != nil {
		in, out := &in.IntegrationConfig, &out.IntegrationConfig
		*out = new(AddonsconfigIntegrationConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MonetizationConfig != nil {
		in, out := &in.MonetizationConfig, &out.MonetizationConfig
		*out = new(AddonsconfigMonetizationConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonsconfigAddonsConfig.
func (in *AddonsconfigAddonsConfig) DeepCopy() *AddonsconfigAddonsConfig {
	if in == nil {
		return nil
	}
	out := new(AddonsconfigAddonsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonsconfigAdvancedApiOpsConfig) DeepCopyInto(out *AddonsconfigAdvancedApiOpsConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonsconfigAdvancedApiOpsConfig.
func (in *AddonsconfigAdvancedApiOpsConfig) DeepCopy() *AddonsconfigAdvancedApiOpsConfig {
	if in == nil {
		return nil
	}
	out := new(AddonsconfigAdvancedApiOpsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonsconfigApiSecurityConfig) DeepCopyInto(out *AddonsconfigApiSecurityConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonsconfigApiSecurityConfig.
func (in *AddonsconfigApiSecurityConfig) DeepCopy() *AddonsconfigApiSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(AddonsconfigApiSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonsconfigConnectorsPlatformConfig) DeepCopyInto(out *AddonsconfigConnectorsPlatformConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExpiresAt != nil {
		in, out := &in.ExpiresAt, &out.ExpiresAt
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonsconfigConnectorsPlatformConfig.
func (in *AddonsconfigConnectorsPlatformConfig) DeepCopy() *AddonsconfigConnectorsPlatformConfig {
	if in == nil {
		return nil
	}
	out := new(AddonsconfigConnectorsPlatformConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonsconfigIntegrationConfig) DeepCopyInto(out *AddonsconfigIntegrationConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonsconfigIntegrationConfig.
func (in *AddonsconfigIntegrationConfig) DeepCopy() *AddonsconfigIntegrationConfig {
	if in == nil {
		return nil
	}
	out := new(AddonsconfigIntegrationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonsconfigMonetizationConfig) DeepCopyInto(out *AddonsconfigMonetizationConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonsconfigMonetizationConfig.
func (in *AddonsconfigMonetizationConfig) DeepCopy() *AddonsconfigMonetizationConfig {
	if in == nil {
		return nil
	}
	out := new(AddonsconfigMonetizationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeAddonsConfig) DeepCopyInto(out *ApigeeAddonsConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeAddonsConfig.
func (in *ApigeeAddonsConfig) DeepCopy() *ApigeeAddonsConfig {
	if in == nil {
		return nil
	}
	out := new(ApigeeAddonsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeAddonsConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeAddonsConfigList) DeepCopyInto(out *ApigeeAddonsConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApigeeAddonsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeAddonsConfigList.
func (in *ApigeeAddonsConfigList) DeepCopy() *ApigeeAddonsConfigList {
	if in == nil {
		return nil
	}
	out := new(ApigeeAddonsConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeAddonsConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeAddonsConfigSpec) DeepCopyInto(out *ApigeeAddonsConfigSpec) {
	*out = *in
	if in.AddonsConfig != nil {
		in, out := &in.AddonsConfig, &out.AddonsConfig
		*out = new(AddonsconfigAddonsConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeAddonsConfigSpec.
func (in *ApigeeAddonsConfigSpec) DeepCopy() *ApigeeAddonsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ApigeeAddonsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeAddonsConfigStatus) DeepCopyInto(out *ApigeeAddonsConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeAddonsConfigStatus.
func (in *ApigeeAddonsConfigStatus) DeepCopy() *ApigeeAddonsConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ApigeeAddonsConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEndpointAttachment) DeepCopyInto(out *ApigeeEndpointAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEndpointAttachment.
func (in *ApigeeEndpointAttachment) DeepCopy() *ApigeeEndpointAttachment {
	if in == nil {
		return nil
	}
	out := new(ApigeeEndpointAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeEndpointAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEndpointAttachmentList) DeepCopyInto(out *ApigeeEndpointAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApigeeEndpointAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEndpointAttachmentList.
func (in *ApigeeEndpointAttachmentList) DeepCopy() *ApigeeEndpointAttachmentList {
	if in == nil {
		return nil
	}
	out := new(ApigeeEndpointAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeEndpointAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEndpointAttachmentSpec) DeepCopyInto(out *ApigeeEndpointAttachmentSpec) {
	*out = *in
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEndpointAttachmentSpec.
func (in *ApigeeEndpointAttachmentSpec) DeepCopy() *ApigeeEndpointAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(ApigeeEndpointAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEndpointAttachmentStatus) DeepCopyInto(out *ApigeeEndpointAttachmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ConnectionState != nil {
		in, out := &in.ConnectionState, &out.ConnectionState
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEndpointAttachmentStatus.
func (in *ApigeeEndpointAttachmentStatus) DeepCopy() *ApigeeEndpointAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(ApigeeEndpointAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEnvgroup) DeepCopyInto(out *ApigeeEnvgroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEnvgroup.
func (in *ApigeeEnvgroup) DeepCopy() *ApigeeEnvgroup {
	if in == nil {
		return nil
	}
	out := new(ApigeeEnvgroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeEnvgroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEnvgroupAttachment) DeepCopyInto(out *ApigeeEnvgroupAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEnvgroupAttachment.
func (in *ApigeeEnvgroupAttachment) DeepCopy() *ApigeeEnvgroupAttachment {
	if in == nil {
		return nil
	}
	out := new(ApigeeEnvgroupAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeEnvgroupAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEnvgroupAttachmentList) DeepCopyInto(out *ApigeeEnvgroupAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApigeeEnvgroupAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEnvgroupAttachmentList.
func (in *ApigeeEnvgroupAttachmentList) DeepCopy() *ApigeeEnvgroupAttachmentList {
	if in == nil {
		return nil
	}
	out := new(ApigeeEnvgroupAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeEnvgroupAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEnvgroupAttachmentSpec) DeepCopyInto(out *ApigeeEnvgroupAttachmentSpec) {
	*out = *in
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEnvgroupAttachmentSpec.
func (in *ApigeeEnvgroupAttachmentSpec) DeepCopy() *ApigeeEnvgroupAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(ApigeeEnvgroupAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEnvgroupAttachmentStatus) DeepCopyInto(out *ApigeeEnvgroupAttachmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEnvgroupAttachmentStatus.
func (in *ApigeeEnvgroupAttachmentStatus) DeepCopy() *ApigeeEnvgroupAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(ApigeeEnvgroupAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEnvgroupList) DeepCopyInto(out *ApigeeEnvgroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApigeeEnvgroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEnvgroupList.
func (in *ApigeeEnvgroupList) DeepCopy() *ApigeeEnvgroupList {
	if in == nil {
		return nil
	}
	out := new(ApigeeEnvgroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeEnvgroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEnvgroupSpec) DeepCopyInto(out *ApigeeEnvgroupSpec) {
	*out = *in
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEnvgroupSpec.
func (in *ApigeeEnvgroupSpec) DeepCopy() *ApigeeEnvgroupSpec {
	if in == nil {
		return nil
	}
	out := new(ApigeeEnvgroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeEnvgroupStatus) DeepCopyInto(out *ApigeeEnvgroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeEnvgroupStatus.
func (in *ApigeeEnvgroupStatus) DeepCopy() *ApigeeEnvgroupStatus {
	if in == nil {
		return nil
	}
	out := new(ApigeeEnvgroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeInstance) DeepCopyInto(out *ApigeeInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeInstance.
func (in *ApigeeInstance) DeepCopy() *ApigeeInstance {
	if in == nil {
		return nil
	}
	out := new(ApigeeInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeInstanceAttachment) DeepCopyInto(out *ApigeeInstanceAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeInstanceAttachment.
func (in *ApigeeInstanceAttachment) DeepCopy() *ApigeeInstanceAttachment {
	if in == nil {
		return nil
	}
	out := new(ApigeeInstanceAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeInstanceAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeInstanceAttachmentList) DeepCopyInto(out *ApigeeInstanceAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApigeeInstanceAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeInstanceAttachmentList.
func (in *ApigeeInstanceAttachmentList) DeepCopy() *ApigeeInstanceAttachmentList {
	if in == nil {
		return nil
	}
	out := new(ApigeeInstanceAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeInstanceAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeInstanceAttachmentSpec) DeepCopyInto(out *ApigeeInstanceAttachmentSpec) {
	*out = *in
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeInstanceAttachmentSpec.
func (in *ApigeeInstanceAttachmentSpec) DeepCopy() *ApigeeInstanceAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(ApigeeInstanceAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeInstanceAttachmentStatus) DeepCopyInto(out *ApigeeInstanceAttachmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeInstanceAttachmentStatus.
func (in *ApigeeInstanceAttachmentStatus) DeepCopy() *ApigeeInstanceAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(ApigeeInstanceAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeInstanceList) DeepCopyInto(out *ApigeeInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApigeeInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeInstanceList.
func (in *ApigeeInstanceList) DeepCopy() *ApigeeInstanceList {
	if in == nil {
		return nil
	}
	out := new(ApigeeInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeInstanceSpec) DeepCopyInto(out *ApigeeInstanceSpec) {
	*out = *in
	if in.ConsumerAcceptList != nil {
		in, out := &in.ConsumerAcceptList, &out.ConsumerAcceptList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionKeyName != nil {
		in, out := &in.DiskEncryptionKeyName, &out.DiskEncryptionKeyName
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.IpRange != nil {
		in, out := &in.IpRange, &out.IpRange
		*out = new(string)
		**out = **in
	}
	if in.PeeringCidrRange != nil {
		in, out := &in.PeeringCidrRange, &out.PeeringCidrRange
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeInstanceSpec.
func (in *ApigeeInstanceSpec) DeepCopy() *ApigeeInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ApigeeInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeInstanceStatus) DeepCopyInto(out *ApigeeInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.ServiceAttachment != nil {
		in, out := &in.ServiceAttachment, &out.ServiceAttachment
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeInstanceStatus.
func (in *ApigeeInstanceStatus) DeepCopy() *ApigeeInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ApigeeInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeNATAddress) DeepCopyInto(out *ApigeeNATAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeNATAddress.
func (in *ApigeeNATAddress) DeepCopy() *ApigeeNATAddress {
	if in == nil {
		return nil
	}
	out := new(ApigeeNATAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeNATAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeNATAddressList) DeepCopyInto(out *ApigeeNATAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApigeeNATAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeNATAddressList.
func (in *ApigeeNATAddressList) DeepCopy() *ApigeeNATAddressList {
	if in == nil {
		return nil
	}
	out := new(ApigeeNATAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeNATAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeNATAddressSpec) DeepCopyInto(out *ApigeeNATAddressSpec) {
	*out = *in
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeNATAddressSpec.
func (in *ApigeeNATAddressSpec) DeepCopy() *ApigeeNATAddressSpec {
	if in == nil {
		return nil
	}
	out := new(ApigeeNATAddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeNATAddressStatus) DeepCopyInto(out *ApigeeNATAddressStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeNATAddressStatus.
func (in *ApigeeNATAddressStatus) DeepCopy() *ApigeeNATAddressStatus {
	if in == nil {
		return nil
	}
	out := new(ApigeeNATAddressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeSyncAuthorization) DeepCopyInto(out *ApigeeSyncAuthorization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeSyncAuthorization.
func (in *ApigeeSyncAuthorization) DeepCopy() *ApigeeSyncAuthorization {
	if in == nil {
		return nil
	}
	out := new(ApigeeSyncAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeSyncAuthorization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeSyncAuthorizationList) DeepCopyInto(out *ApigeeSyncAuthorizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApigeeSyncAuthorization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeSyncAuthorizationList.
func (in *ApigeeSyncAuthorizationList) DeepCopy() *ApigeeSyncAuthorizationList {
	if in == nil {
		return nil
	}
	out := new(ApigeeSyncAuthorizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApigeeSyncAuthorizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeSyncAuthorizationSpec) DeepCopyInto(out *ApigeeSyncAuthorizationSpec) {
	*out = *in
	if in.Identities != nil {
		in, out := &in.Identities, &out.Identities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeSyncAuthorizationSpec.
func (in *ApigeeSyncAuthorizationSpec) DeepCopy() *ApigeeSyncAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(ApigeeSyncAuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApigeeSyncAuthorizationStatus) DeepCopyInto(out *ApigeeSyncAuthorizationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApigeeSyncAuthorizationStatus.
func (in *ApigeeSyncAuthorizationStatus) DeepCopy() *ApigeeSyncAuthorizationStatus {
	if in == nil {
		return nil
	}
	out := new(ApigeeSyncAuthorizationStatus)
	in.DeepCopyInto(out)
	return out
}
