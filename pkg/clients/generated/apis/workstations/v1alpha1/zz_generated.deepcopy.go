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
func (in *WorkstationclusterDetailsStatus) DeepCopyInto(out *WorkstationclusterDetailsStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationclusterDetailsStatus.
func (in *WorkstationclusterDetailsStatus) DeepCopy() *WorkstationclusterDetailsStatus {
	if in == nil {
		return nil
	}
	out := new(WorkstationclusterDetailsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationclusterPrivateClusterConfig) DeepCopyInto(out *WorkstationclusterPrivateClusterConfig) {
	*out = *in
	if in.ClusterHostname != nil {
		in, out := &in.ClusterHostname, &out.ClusterHostname
		*out = new(string)
		**out = **in
	}
	if in.ServiceAttachmentUri != nil {
		in, out := &in.ServiceAttachmentUri, &out.ServiceAttachmentUri
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationclusterPrivateClusterConfig.
func (in *WorkstationclusterPrivateClusterConfig) DeepCopy() *WorkstationclusterPrivateClusterConfig {
	if in == nil {
		return nil
	}
	out := new(WorkstationclusterPrivateClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationsWorkstationCluster) DeepCopyInto(out *WorkstationsWorkstationCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationsWorkstationCluster.
func (in *WorkstationsWorkstationCluster) DeepCopy() *WorkstationsWorkstationCluster {
	if in == nil {
		return nil
	}
	out := new(WorkstationsWorkstationCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkstationsWorkstationCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationsWorkstationClusterList) DeepCopyInto(out *WorkstationsWorkstationClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkstationsWorkstationCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationsWorkstationClusterList.
func (in *WorkstationsWorkstationClusterList) DeepCopy() *WorkstationsWorkstationClusterList {
	if in == nil {
		return nil
	}
	out := new(WorkstationsWorkstationClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkstationsWorkstationClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationsWorkstationClusterSpec) DeepCopyInto(out *WorkstationsWorkstationClusterSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.PrivateClusterConfig != nil {
		in, out := &in.PrivateClusterConfig, &out.PrivateClusterConfig
		*out = new(WorkstationclusterPrivateClusterConfig)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationsWorkstationClusterSpec.
func (in *WorkstationsWorkstationClusterSpec) DeepCopy() *WorkstationsWorkstationClusterSpec {
	if in == nil {
		return nil
	}
	out := new(WorkstationsWorkstationClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkstationsWorkstationClusterStatus) DeepCopyInto(out *WorkstationsWorkstationClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.Degraded != nil {
		in, out := &in.Degraded, &out.Degraded
		*out = new(bool)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
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
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkstationsWorkstationClusterStatus.
func (in *WorkstationsWorkstationClusterStatus) DeepCopy() *WorkstationsWorkstationClusterStatus {
	if in == nil {
		return nil
	}
	out := new(WorkstationsWorkstationClusterStatus)
	in.DeepCopyInto(out)
	return out
}
