//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageCheck) DeepCopyInto(out *PackageCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make([]PackageCheckSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageCheck.
func (in *PackageCheck) DeepCopy() *PackageCheck {
	if in == nil {
		return nil
	}
	out := new(PackageCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageCheckList) DeepCopyInto(out *PackageCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PackageCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageCheckList.
func (in *PackageCheckList) DeepCopy() *PackageCheckList {
	if in == nil {
		return nil
	}
	out := new(PackageCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageCheckSpec) DeepCopyInto(out *PackageCheckSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageCheckSpec.
func (in *PackageCheckSpec) DeepCopy() *PackageCheckSpec {
	if in == nil {
		return nil
	}
	out := new(PackageCheckSpec)
	in.DeepCopyInto(out)
	return out
}