//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetric) DeepCopyInto(out *KeptnMetric) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetric.
func (in *KeptnMetric) DeepCopy() *KeptnMetric {
	if in == nil {
		return nil
	}
	out := new(KeptnMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnMetric) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricList) DeepCopyInto(out *KeptnMetricList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricList.
func (in *KeptnMetricList) DeepCopy() *KeptnMetricList {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnMetricList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricSpec) DeepCopyInto(out *KeptnMetricSpec) {
	*out = *in
	out.Provider = in.Provider
	out.Range = in.Range
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricSpec.
func (in *KeptnMetricSpec) DeepCopy() *KeptnMetricSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricStatus) DeepCopyInto(out *KeptnMetricStatus) {
	*out = *in
	if in.RawValue != nil {
		in, out := &in.RawValue, &out.RawValue
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricStatus.
func (in *KeptnMetricStatus) DeepCopy() *KeptnMetricStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricsProvider) DeepCopyInto(out *KeptnMetricsProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricsProvider.
func (in *KeptnMetricsProvider) DeepCopy() *KeptnMetricsProvider {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricsProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnMetricsProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricsProviderList) DeepCopyInto(out *KeptnMetricsProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnMetricsProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricsProviderList.
func (in *KeptnMetricsProviderList) DeepCopy() *KeptnMetricsProviderList {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricsProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnMetricsProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricsProviderSpec) DeepCopyInto(out *KeptnMetricsProviderSpec) {
	*out = *in
	in.SecretKeyRef.DeepCopyInto(&out.SecretKeyRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricsProviderSpec.
func (in *KeptnMetricsProviderSpec) DeepCopy() *KeptnMetricsProviderSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricsProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnMetricsProviderStatus) DeepCopyInto(out *KeptnMetricsProviderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnMetricsProviderStatus.
func (in *KeptnMetricsProviderStatus) DeepCopy() *KeptnMetricsProviderStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnMetricsProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderRef) DeepCopyInto(out *ProviderRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderRef.
func (in *ProviderRef) DeepCopy() *ProviderRef {
	if in == nil {
		return nil
	}
	out := new(ProviderRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RangeSpec) DeepCopyInto(out *RangeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RangeSpec.
func (in *RangeSpec) DeepCopy() *RangeSpec {
	if in == nil {
		return nil
	}
	out := new(RangeSpec)
	in.DeepCopyInto(out)
	return out
}
