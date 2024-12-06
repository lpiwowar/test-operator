//go:build !ignore_autogenerated

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

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleTest) DeepCopyInto(out *AnsibleTest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleTest.
func (in *AnsibleTest) DeepCopy() *AnsibleTest {
	if in == nil {
		return nil
	}
	out := new(AnsibleTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsibleTest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleTestList) DeepCopyInto(out *AnsibleTestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnsibleTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleTestList.
func (in *AnsibleTestList) DeepCopy() *AnsibleTestList {
	if in == nil {
		return nil
	}
	out := new(AnsibleTestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnsibleTestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleTestSpec) DeepCopyInto(out *AnsibleTestSpec) {
	*out = *in
	in.CommonOptions.DeepCopyInto(&out.CommonOptions)
	out.CommonOpenstackConfig = in.CommonOpenstackConfig
	if in.Workflow != nil {
		in, out := &in.Workflow, &out.Workflow
		*out = make([]AnsibleTestWorkflowSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleTestSpec.
func (in *AnsibleTestSpec) DeepCopy() *AnsibleTestSpec {
	if in == nil {
		return nil
	}
	out := new(AnsibleTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnsibleTestWorkflowSpec) DeepCopyInto(out *AnsibleTestWorkflowSpec) {
	*out = *in
	in.WorkflowCommonParameters.DeepCopyInto(&out.WorkflowCommonParameters)
	out.CommonOpenstackConfig = in.CommonOpenstackConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnsibleTestWorkflowSpec.
func (in *AnsibleTestWorkflowSpec) DeepCopy() *AnsibleTestWorkflowSpec {
	if in == nil {
		return nil
	}
	out := new(AnsibleTestWorkflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonOpenstackConfig) DeepCopyInto(out *CommonOpenstackConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonOpenstackConfig.
func (in *CommonOpenstackConfig) DeepCopy() *CommonOpenstackConfig {
	if in == nil {
		return nil
	}
	out := new(CommonOpenstackConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonOptions) DeepCopyInto(out *CommonOptions) {
	*out = *in
	if in.BackoffLimit != nil {
		in, out := &in.BackoffLimit, &out.BackoffLimit
		*out = new(int32)
		**out = **in
	}
	if in.ExtraConfigmapsMounts != nil {
		in, out := &in.ExtraConfigmapsMounts, &out.ExtraConfigmapsMounts
		*out = make([]ExtraConfigmapsMounts, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonOptions.
func (in *CommonOptions) DeepCopy() *CommonOptions {
	if in == nil {
		return nil
	}
	out := new(CommonOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonTestStatus) DeepCopyInto(out *CommonTestStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonTestStatus.
func (in *CommonTestStatus) DeepCopy() *CommonTestStatus {
	if in == nil {
		return nil
	}
	out := new(CommonTestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalPluginType) DeepCopyInto(out *ExternalPluginType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalPluginType.
func (in *ExternalPluginType) DeepCopy() *ExternalPluginType {
	if in == nil {
		return nil
	}
	out := new(ExternalPluginType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraConfigmapsMounts) DeepCopyInto(out *ExtraConfigmapsMounts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraConfigmapsMounts.
func (in *ExtraConfigmapsMounts) DeepCopy() *ExtraConfigmapsMounts {
	if in == nil {
		return nil
	}
	out := new(ExtraConfigmapsMounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraImagesFlavorType) DeepCopyInto(out *ExtraImagesFlavorType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraImagesFlavorType.
func (in *ExtraImagesFlavorType) DeepCopy() *ExtraImagesFlavorType {
	if in == nil {
		return nil
	}
	out := new(ExtraImagesFlavorType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraImagesType) DeepCopyInto(out *ExtraImagesType) {
	*out = *in
	out.Flavor = in.Flavor
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraImagesType.
func (in *ExtraImagesType) DeepCopy() *ExtraImagesType {
	if in == nil {
		return nil
	}
	out := new(ExtraImagesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizonTest) DeepCopyInto(out *HorizonTest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizonTest.
func (in *HorizonTest) DeepCopy() *HorizonTest {
	if in == nil {
		return nil
	}
	out := new(HorizonTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HorizonTest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizonTestList) DeepCopyInto(out *HorizonTestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HorizonTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizonTestList.
func (in *HorizonTestList) DeepCopy() *HorizonTestList {
	if in == nil {
		return nil
	}
	out := new(HorizonTestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HorizonTestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizonTestSpec) DeepCopyInto(out *HorizonTestSpec) {
	*out = *in
	in.CommonOptions.DeepCopyInto(&out.CommonOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizonTestSpec.
func (in *HorizonTestSpec) DeepCopy() *HorizonTestSpec {
	if in == nil {
		return nil
	}
	out := new(HorizonTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tempest) DeepCopyInto(out *Tempest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tempest.
func (in *Tempest) DeepCopy() *Tempest {
	if in == nil {
		return nil
	}
	out := new(Tempest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tempest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempestDefaults) DeepCopyInto(out *TempestDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempestDefaults.
func (in *TempestDefaults) DeepCopy() *TempestDefaults {
	if in == nil {
		return nil
	}
	out := new(TempestDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempestList) DeepCopyInto(out *TempestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tempest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempestList.
func (in *TempestList) DeepCopy() *TempestList {
	if in == nil {
		return nil
	}
	out := new(TempestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TempestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempestRunSpec) DeepCopyInto(out *TempestRunSpec) {
	*out = *in
	if in.ExternalPlugin != nil {
		in, out := &in.ExternalPlugin, &out.ExternalPlugin
		*out = make([]ExternalPluginType, len(*in))
		copy(*out, *in)
	}
	if in.ExtraRPMs != nil {
		in, out := &in.ExtraRPMs, &out.ExtraRPMs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraImages != nil {
		in, out := &in.ExtraImages, &out.ExtraImages
		*out = make([]ExtraImagesType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempestRunSpec.
func (in *TempestRunSpec) DeepCopy() *TempestRunSpec {
	if in == nil {
		return nil
	}
	out := new(TempestRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempestSpec) DeepCopyInto(out *TempestSpec) {
	*out = *in
	in.CommonOptions.DeepCopyInto(&out.CommonOptions)
	out.CommonOpenstackConfig = in.CommonOpenstackConfig
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.TempestRun.DeepCopyInto(&out.TempestRun)
	out.TempestconfRun = in.TempestconfRun
	if in.ConfigOverwrite != nil {
		in, out := &in.ConfigOverwrite, &out.ConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Workflow != nil {
		in, out := &in.Workflow, &out.Workflow
		*out = make([]WorkflowTempestSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempestSpec.
func (in *TempestSpec) DeepCopy() *TempestSpec {
	if in == nil {
		return nil
	}
	out := new(TempestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TempestconfRunSpec) DeepCopyInto(out *TempestconfRunSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TempestconfRunSpec.
func (in *TempestconfRunSpec) DeepCopy() *TempestconfRunSpec {
	if in == nil {
		return nil
	}
	out := new(TempestconfRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tobiko) DeepCopyInto(out *Tobiko) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tobiko.
func (in *Tobiko) DeepCopy() *Tobiko {
	if in == nil {
		return nil
	}
	out := new(Tobiko)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tobiko) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TobikoList) DeepCopyInto(out *TobikoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tobiko, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TobikoList.
func (in *TobikoList) DeepCopy() *TobikoList {
	if in == nil {
		return nil
	}
	out := new(TobikoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TobikoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TobikoSpec) DeepCopyInto(out *TobikoSpec) {
	*out = *in
	in.CommonOptions.DeepCopyInto(&out.CommonOptions)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Workflow != nil {
		in, out := &in.Workflow, &out.Workflow
		*out = make([]TobikoWorkflowSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TobikoSpec.
func (in *TobikoSpec) DeepCopy() *TobikoSpec {
	if in == nil {
		return nil
	}
	out := new(TobikoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TobikoWorkflowSpec) DeepCopyInto(out *TobikoWorkflowSpec) {
	*out = *in
	in.WorkflowCommonParameters.DeepCopyInto(&out.WorkflowCommonParameters)
	if in.PreventCreate != nil {
		in, out := &in.PreventCreate, &out.PreventCreate
		*out = new(bool)
		**out = **in
	}
	if in.NumProcesses != nil {
		in, out := &in.NumProcesses, &out.NumProcesses
		*out = new(uint8)
		**out = **in
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TobikoWorkflowSpec.
func (in *TobikoWorkflowSpec) DeepCopy() *TobikoWorkflowSpec {
	if in == nil {
		return nil
	}
	out := new(TobikoWorkflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowCommonParameters) DeepCopyInto(out *WorkflowCommonParameters) {
	*out = *in
	if in.Privileged != nil {
		in, out := &in.Privileged, &out.Privileged
		*out = new(bool)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.SELinuxLevel != nil {
		in, out := &in.SELinuxLevel, &out.SELinuxLevel
		*out = new(string)
		**out = **in
	}
	if in.BackoffLimit != nil {
		in, out := &in.BackoffLimit, &out.BackoffLimit
		*out = new(int32)
		**out = **in
	}
	if in.ExtraConfigmapsMounts != nil {
		in, out := &in.ExtraConfigmapsMounts, &out.ExtraConfigmapsMounts
		*out = new([]ExtraConfigmapsMounts)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ExtraConfigmapsMounts, len(*in))
			copy(*out, *in)
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = new([]v1.Toleration)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.Toleration, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowCommonParameters.
func (in *WorkflowCommonParameters) DeepCopy() *WorkflowCommonParameters {
	if in == nil {
		return nil
	}
	out := new(WorkflowCommonParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTempestRunSpec) DeepCopyInto(out *WorkflowTempestRunSpec) {
	*out = *in
	if in.IncludeList != nil {
		in, out := &in.IncludeList, &out.IncludeList
		*out = new(string)
		**out = **in
	}
	if in.ExcludeList != nil {
		in, out := &in.ExcludeList, &out.ExcludeList
		*out = new(string)
		**out = **in
	}
	if in.ExpectedFailuresList != nil {
		in, out := &in.ExpectedFailuresList, &out.ExpectedFailuresList
		*out = new(string)
		**out = **in
	}
	if in.Concurrency != nil {
		in, out := &in.Concurrency, &out.Concurrency
		*out = new(int64)
		**out = **in
	}
	if in.Smoke != nil {
		in, out := &in.Smoke, &out.Smoke
		*out = new(bool)
		**out = **in
	}
	if in.Parallel != nil {
		in, out := &in.Parallel, &out.Parallel
		*out = new(bool)
		**out = **in
	}
	if in.Serial != nil {
		in, out := &in.Serial, &out.Serial
		*out = new(bool)
		**out = **in
	}
	if in.WorkerFile != nil {
		in, out := &in.WorkerFile, &out.WorkerFile
		*out = new(string)
		**out = **in
	}
	if in.ExternalPlugin != nil {
		in, out := &in.ExternalPlugin, &out.ExternalPlugin
		*out = new([]ExternalPluginType)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ExternalPluginType, len(*in))
			copy(*out, *in)
		}
	}
	if in.ExtraRPMs != nil {
		in, out := &in.ExtraRPMs, &out.ExtraRPMs
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.ExtraImages != nil {
		in, out := &in.ExtraImages, &out.ExtraImages
		*out = new([]ExtraImagesType)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ExtraImagesType, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTempestRunSpec.
func (in *WorkflowTempestRunSpec) DeepCopy() *WorkflowTempestRunSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowTempestRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTempestSpec) DeepCopyInto(out *WorkflowTempestSpec) {
	*out = *in
	in.WorkflowCommonParameters.DeepCopyInto(&out.WorkflowCommonParameters)
	out.CommonOpenstackConfig = in.CommonOpenstackConfig
	if in.Parallel != nil {
		in, out := &in.Parallel, &out.Parallel
		*out = new(bool)
		**out = **in
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	in.TempestRun.DeepCopyInto(&out.TempestRun)
	in.TempestconfRun.DeepCopyInto(&out.TempestconfRun)
	if in.SSHKeySecretName != nil {
		in, out := &in.SSHKeySecretName, &out.SSHKeySecretName
		*out = new(string)
		**out = **in
	}
	if in.ConfigOverwrite != nil {
		in, out := &in.ConfigOverwrite, &out.ConfigOverwrite
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTempestSpec.
func (in *WorkflowTempestSpec) DeepCopy() *WorkflowTempestSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowTempestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowTempestconfRunSpec) DeepCopyInto(out *WorkflowTempestconfRunSpec) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(bool)
		**out = **in
	}
	if in.CollectTiming != nil {
		in, out := &in.CollectTiming, &out.CollectTiming
		*out = new(bool)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.NoDefaultDeployer != nil {
		in, out := &in.NoDefaultDeployer, &out.NoDefaultDeployer
		*out = new(bool)
		**out = **in
	}
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(bool)
		**out = **in
	}
	if in.Verbose != nil {
		in, out := &in.Verbose, &out.Verbose
		*out = new(bool)
		**out = **in
	}
	if in.NonAdmin != nil {
		in, out := &in.NonAdmin, &out.NonAdmin
		*out = new(bool)
		**out = **in
	}
	if in.RetryImage != nil {
		in, out := &in.RetryImage, &out.RetryImage
		*out = new(bool)
		**out = **in
	}
	if in.ConvertToRaw != nil {
		in, out := &in.ConvertToRaw, &out.ConvertToRaw
		*out = new(bool)
		**out = **in
	}
	if in.Out != nil {
		in, out := &in.Out, &out.Out
		*out = new(string)
		**out = **in
	}
	if in.DeployerInput != nil {
		in, out := &in.DeployerInput, &out.DeployerInput
		*out = new(string)
		**out = **in
	}
	if in.TestAccounts != nil {
		in, out := &in.TestAccounts, &out.TestAccounts
		*out = new(string)
		**out = **in
	}
	if in.CreateAccountsFile != nil {
		in, out := &in.CreateAccountsFile, &out.CreateAccountsFile
		*out = new(string)
		**out = **in
	}
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = new(string)
		**out = **in
	}
	if in.GenerateProfile != nil {
		in, out := &in.GenerateProfile, &out.GenerateProfile
		*out = new(string)
		**out = **in
	}
	if in.ImageDiskFormat != nil {
		in, out := &in.ImageDiskFormat, &out.ImageDiskFormat
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.FlavorMinMem != nil {
		in, out := &in.FlavorMinMem, &out.FlavorMinMem
		*out = new(int64)
		**out = **in
	}
	if in.FlavorMinDisk != nil {
		in, out := &in.FlavorMinDisk, &out.FlavorMinDisk
		*out = new(int64)
		**out = **in
	}
	if in.NetworkID != nil {
		in, out := &in.NetworkID, &out.NetworkID
		*out = new(string)
		**out = **in
	}
	if in.Append != nil {
		in, out := &in.Append, &out.Append
		*out = new(string)
		**out = **in
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = new(string)
		**out = **in
	}
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowTempestconfRunSpec.
func (in *WorkflowTempestconfRunSpec) DeepCopy() *WorkflowTempestconfRunSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowTempestconfRunSpec)
	in.DeepCopyInto(out)
	return out
}
