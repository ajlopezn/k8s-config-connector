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
func (in *PubSubLiteSubscription) DeepCopyInto(out *PubSubLiteSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubLiteSubscription.
func (in *PubSubLiteSubscription) DeepCopy() *PubSubLiteSubscription {
	if in == nil {
		return nil
	}
	out := new(PubSubLiteSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PubSubLiteSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubLiteSubscriptionList) DeepCopyInto(out *PubSubLiteSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PubSubLiteSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubLiteSubscriptionList.
func (in *PubSubLiteSubscriptionList) DeepCopy() *PubSubLiteSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(PubSubLiteSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PubSubLiteSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubLiteSubscriptionSpec) DeepCopyInto(out *PubSubLiteSubscriptionSpec) {
	*out = *in
	if in.DeliveryConfig != nil {
		in, out := &in.DeliveryConfig, &out.DeliveryConfig
		*out = new(SubscriptionDeliveryConfig)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.Region != nil {
		in, out := &in.Region, &out.Region
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubLiteSubscriptionSpec.
func (in *PubSubLiteSubscriptionSpec) DeepCopy() *PubSubLiteSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(PubSubLiteSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubLiteSubscriptionStatus) DeepCopyInto(out *PubSubLiteSubscriptionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubLiteSubscriptionStatus.
func (in *PubSubLiteSubscriptionStatus) DeepCopy() *PubSubLiteSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(PubSubLiteSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubLiteTopic) DeepCopyInto(out *PubSubLiteTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubLiteTopic.
func (in *PubSubLiteTopic) DeepCopy() *PubSubLiteTopic {
	if in == nil {
		return nil
	}
	out := new(PubSubLiteTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PubSubLiteTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubLiteTopicList) DeepCopyInto(out *PubSubLiteTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PubSubLiteTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubLiteTopicList.
func (in *PubSubLiteTopicList) DeepCopy() *PubSubLiteTopicList {
	if in == nil {
		return nil
	}
	out := new(PubSubLiteTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PubSubLiteTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubLiteTopicSpec) DeepCopyInto(out *PubSubLiteTopicSpec) {
	*out = *in
	if in.PartitionConfig != nil {
		in, out := &in.PartitionConfig, &out.PartitionConfig
		*out = new(TopicPartitionConfig)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ReservationConfig != nil {
		in, out := &in.ReservationConfig, &out.ReservationConfig
		*out = new(TopicReservationConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.RetentionConfig != nil {
		in, out := &in.RetentionConfig, &out.RetentionConfig
		*out = new(TopicRetentionConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubLiteTopicSpec.
func (in *PubSubLiteTopicSpec) DeepCopy() *PubSubLiteTopicSpec {
	if in == nil {
		return nil
	}
	out := new(PubSubLiteTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PubSubLiteTopicStatus) DeepCopyInto(out *PubSubLiteTopicStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PubSubLiteTopicStatus.
func (in *PubSubLiteTopicStatus) DeepCopy() *PubSubLiteTopicStatus {
	if in == nil {
		return nil
	}
	out := new(PubSubLiteTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionDeliveryConfig) DeepCopyInto(out *SubscriptionDeliveryConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionDeliveryConfig.
func (in *SubscriptionDeliveryConfig) DeepCopy() *SubscriptionDeliveryConfig {
	if in == nil {
		return nil
	}
	out := new(SubscriptionDeliveryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicCapacity) DeepCopyInto(out *TopicCapacity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicCapacity.
func (in *TopicCapacity) DeepCopy() *TopicCapacity {
	if in == nil {
		return nil
	}
	out := new(TopicCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicPartitionConfig) DeepCopyInto(out *TopicPartitionConfig) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(TopicCapacity)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicPartitionConfig.
func (in *TopicPartitionConfig) DeepCopy() *TopicPartitionConfig {
	if in == nil {
		return nil
	}
	out := new(TopicPartitionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicReservationConfig) DeepCopyInto(out *TopicReservationConfig) {
	*out = *in
	if in.ThroughputReservation != nil {
		in, out := &in.ThroughputReservation, &out.ThroughputReservation
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicReservationConfig.
func (in *TopicReservationConfig) DeepCopy() *TopicReservationConfig {
	if in == nil {
		return nil
	}
	out := new(TopicReservationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopicRetentionConfig) DeepCopyInto(out *TopicRetentionConfig) {
	*out = *in
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopicRetentionConfig.
func (in *TopicRetentionConfig) DeepCopy() *TopicRetentionConfig {
	if in == nil {
		return nil
	}
	out := new(TopicRetentionConfig)
	in.DeepCopyInto(out)
	return out
}
