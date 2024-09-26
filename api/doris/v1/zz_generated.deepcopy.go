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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdminUser) DeepCopyInto(out *AdminUser) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdminUser.
func (in *AdminUser) DeepCopy() *AdminUser {
	if in == nil {
		return nil
	}
	out := new(AdminUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingPolicy) DeepCopyInto(out *AutoScalingPolicy) {
	*out = *in
	if in.HPAPolicy != nil {
		in, out := &in.HPAPolicy, &out.HPAPolicy
		*out = new(HPAPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingPolicy.
func (in *AutoScalingPolicy) DeepCopy() *AutoScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(AutoScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseSpec) DeepCopyInto(out *BaseSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ExportService)
		(*in).DeepCopyInto(*out)
	}
	if in.FeAddress != nil {
		in, out := &in.FeAddress, &out.FeAddress
		*out = new(FeAddress)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	in.ConfigMapInfo.DeepCopyInto(&out.ConfigMapInfo)
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]corev1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PersistentVolumes != nil {
		in, out := &in.PersistentVolumes, &out.PersistentVolumes
		*out = make([]PersistentVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SystemInitialization != nil {
		in, out := &in.SystemInitialization, &out.SystemInitialization
		*out = new(SystemInitialization)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]Secret, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseSpec.
func (in *BaseSpec) DeepCopy() *BaseSpec {
	if in == nil {
		return nil
	}
	out := new(BaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BeSpec) DeepCopyInto(out *BeSpec) {
	*out = *in
	in.BaseSpec.DeepCopyInto(&out.BaseSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BeSpec.
func (in *BeSpec) DeepCopy() *BeSpec {
	if in == nil {
		return nil
	}
	out := new(BeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSpec) DeepCopyInto(out *BrokerSpec) {
	*out = *in
	in.BaseSpec.DeepCopyInto(&out.BaseSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSpec.
func (in *BrokerSpec) DeepCopy() *BrokerSpec {
	if in == nil {
		return nil
	}
	out := new(BrokerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnSpec) DeepCopyInto(out *CnSpec) {
	*out = *in
	in.BaseSpec.DeepCopyInto(&out.BaseSpec)
	if in.AutoScalingPolicy != nil {
		in, out := &in.AutoScalingPolicy, &out.AutoScalingPolicy
		*out = new(AutoScalingPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnSpec.
func (in *CnSpec) DeepCopy() *CnSpec {
	if in == nil {
		return nil
	}
	out := new(CnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnStatus) DeepCopyInto(out *CnStatus) {
	*out = *in
	in.ComponentStatus.DeepCopyInto(&out.ComponentStatus)
	if in.HorizontalScaler != nil {
		in, out := &in.HorizontalScaler, &out.HorizontalScaler
		*out = new(HorizontalScaler)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnStatus.
func (in *CnStatus) DeepCopy() *CnStatus {
	if in == nil {
		return nil
	}
	out := new(CnStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentCondition) DeepCopyInto(out *ComponentCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentCondition.
func (in *ComponentCondition) DeepCopy() *ComponentCondition {
	if in == nil {
		return nil
	}
	out := new(ComponentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatus) DeepCopyInto(out *ComponentStatus) {
	*out = *in
	if in.FailedMembers != nil {
		in, out := &in.FailedMembers, &out.FailedMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreatingMembers != nil {
		in, out := &in.CreatingMembers, &out.CreatingMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RunningMembers != nil {
		in, out := &in.RunningMembers, &out.RunningMembers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ComponentCondition.DeepCopyInto(&out.ComponentCondition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatus.
func (in *ComponentStatus) DeepCopy() *ComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapInfo) DeepCopyInto(out *ConfigMapInfo) {
	*out = *in
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]MountConfigMapInfo, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapInfo.
func (in *ConfigMapInfo) DeepCopy() *ConfigMapInfo {
	if in == nil {
		return nil
	}
	out := new(ConfigMapInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerResourceMetricSource) DeepCopyInto(out *ContainerResourceMetricSource) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerResourceMetricSource.
func (in *ContainerResourceMetricSource) DeepCopy() *ContainerResourceMetricSource {
	if in == nil {
		return nil
	}
	out := new(ContainerResourceMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossVersionObjectReference) DeepCopyInto(out *CrossVersionObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossVersionObjectReference.
func (in *CrossVersionObjectReference) DeepCopy() *CrossVersionObjectReference {
	if in == nil {
		return nil
	}
	out := new(CrossVersionObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisCluster) DeepCopyInto(out *DorisCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisCluster.
func (in *DorisCluster) DeepCopy() *DorisCluster {
	if in == nil {
		return nil
	}
	out := new(DorisCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DorisCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisClusterList) DeepCopyInto(out *DorisClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DorisCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisClusterList.
func (in *DorisClusterList) DeepCopy() *DorisClusterList {
	if in == nil {
		return nil
	}
	out := new(DorisClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DorisClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisClusterSpec) DeepCopyInto(out *DorisClusterSpec) {
	*out = *in
	if in.FeSpec != nil {
		in, out := &in.FeSpec, &out.FeSpec
		*out = new(FeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BeSpec != nil {
		in, out := &in.BeSpec, &out.BeSpec
		*out = new(BeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CnSpec != nil {
		in, out := &in.CnSpec, &out.CnSpec
		*out = new(CnSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BrokerSpec != nil {
		in, out := &in.BrokerSpec, &out.BrokerSpec
		*out = new(BrokerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AdminUser != nil {
		in, out := &in.AdminUser, &out.AdminUser
		*out = new(AdminUser)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisClusterSpec.
func (in *DorisClusterSpec) DeepCopy() *DorisClusterSpec {
	if in == nil {
		return nil
	}
	out := new(DorisClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisClusterStatus) DeepCopyInto(out *DorisClusterStatus) {
	*out = *in
	if in.FEStatus != nil {
		in, out := &in.FEStatus, &out.FEStatus
		*out = new(ComponentStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.BEStatus != nil {
		in, out := &in.BEStatus, &out.BEStatus
		*out = new(ComponentStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.CnStatus != nil {
		in, out := &in.CnStatus, &out.CnStatus
		*out = new(CnStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.BrokerStatus != nil {
		in, out := &in.BrokerStatus, &out.BrokerStatus
		*out = new(ComponentStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisClusterStatus.
func (in *DorisClusterStatus) DeepCopy() *DorisClusterStatus {
	if in == nil {
		return nil
	}
	out := new(DorisClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DorisServicePort) DeepCopyInto(out *DorisServicePort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DorisServicePort.
func (in *DorisServicePort) DeepCopy() *DorisServicePort {
	if in == nil {
		return nil
	}
	out := new(DorisServicePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoints) DeepCopyInto(out *Endpoints) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoints.
func (in *Endpoints) DeepCopy() *Endpoints {
	if in == nil {
		return nil
	}
	out := new(Endpoints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportService) DeepCopyInto(out *ExportService) {
	*out = *in
	if in.ServicePorts != nil {
		in, out := &in.ServicePorts, &out.ServicePorts
		*out = make([]DorisServicePort, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportService.
func (in *ExportService) DeepCopy() *ExportService {
	if in == nil {
		return nil
	}
	out := new(ExportService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetricSource) DeepCopyInto(out *ExternalMetricSource) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricSource.
func (in *ExternalMetricSource) DeepCopy() *ExternalMetricSource {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeAddress) DeepCopyInto(out *FeAddress) {
	*out = *in
	in.Endpoints.DeepCopyInto(&out.Endpoints)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeAddress.
func (in *FeAddress) DeepCopy() *FeAddress {
	if in == nil {
		return nil
	}
	out := new(FeAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeSpec) DeepCopyInto(out *FeSpec) {
	*out = *in
	if in.ElectionNumber != nil {
		in, out := &in.ElectionNumber, &out.ElectionNumber
		*out = new(int32)
		**out = **in
	}
	in.BaseSpec.DeepCopyInto(&out.BaseSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeSpec.
func (in *FeSpec) DeepCopy() *FeSpec {
	if in == nil {
		return nil
	}
	out := new(FeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAPolicy) DeepCopyInto(out *HPAPolicy) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAPolicy.
func (in *HPAPolicy) DeepCopy() *HPAPolicy {
	if in == nil {
		return nil
	}
	out := new(HPAPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAScalingPolicy) DeepCopyInto(out *HPAScalingPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAScalingPolicy.
func (in *HPAScalingPolicy) DeepCopy() *HPAScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(HPAScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPAScalingRules) DeepCopyInto(out *HPAScalingRules) {
	*out = *in
	if in.StabilizationWindowSeconds != nil {
		in, out := &in.StabilizationWindowSeconds, &out.StabilizationWindowSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SelectPolicy != nil {
		in, out := &in.SelectPolicy, &out.SelectPolicy
		*out = new(ScalingPolicySelect)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]HPAScalingPolicy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPAScalingRules.
func (in *HPAScalingRules) DeepCopy() *HPAScalingRules {
	if in == nil {
		return nil
	}
	out := new(HPAScalingRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalPodAutoscalerBehavior) DeepCopyInto(out *HorizontalPodAutoscalerBehavior) {
	*out = *in
	if in.ScaleUp != nil {
		in, out := &in.ScaleUp, &out.ScaleUp
		*out = new(HPAScalingRules)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleDown != nil {
		in, out := &in.ScaleDown, &out.ScaleDown
		*out = new(HPAScalingRules)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalPodAutoscalerBehavior.
func (in *HorizontalPodAutoscalerBehavior) DeepCopy() *HorizontalPodAutoscalerBehavior {
	if in == nil {
		return nil
	}
	out := new(HorizontalPodAutoscalerBehavior)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalScaler) DeepCopyInto(out *HorizontalScaler) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalScaler.
func (in *HorizontalScaler) DeepCopy() *HorizontalScaler {
	if in == nil {
		return nil
	}
	out := new(HorizontalScaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricIdentifier) DeepCopyInto(out *MetricIdentifier) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricIdentifier.
func (in *MetricIdentifier) DeepCopy() *MetricIdentifier {
	if in == nil {
		return nil
	}
	out := new(MetricIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(ObjectMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(PodsMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerResource != nil {
		in, out := &in.ContainerResource, &out.ContainerResource
		*out = new(ContainerResourceMetricSource)
		(*in).DeepCopyInto(*out)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalMetricSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTarget) DeepCopyInto(out *MetricTarget) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AverageValue != nil {
		in, out := &in.AverageValue, &out.AverageValue
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.AverageUtilization != nil {
		in, out := &in.AverageUtilization, &out.AverageUtilization
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTarget.
func (in *MetricTarget) DeepCopy() *MetricTarget {
	if in == nil {
		return nil
	}
	out := new(MetricTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountConfigMapInfo) DeepCopyInto(out *MountConfigMapInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountConfigMapInfo.
func (in *MountConfigMapInfo) DeepCopy() *MountConfigMapInfo {
	if in == nil {
		return nil
	}
	out := new(MountConfigMapInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMetricSource) DeepCopyInto(out *ObjectMetricSource) {
	*out = *in
	out.DescribedObject = in.DescribedObject
	in.Target.DeepCopyInto(&out.Target)
	in.Metric.DeepCopyInto(&out.Metric)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMetricSource.
func (in *ObjectMetricSource) DeepCopy() *ObjectMetricSource {
	if in == nil {
		return nil
	}
	out := new(ObjectMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolume) DeepCopyInto(out *PersistentVolume) {
	*out = *in
	in.PersistentVolumeClaimSpec.DeepCopyInto(&out.PersistentVolumeClaimSpec)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolume.
func (in *PersistentVolume) DeepCopy() *PersistentVolume {
	if in == nil {
		return nil
	}
	out := new(PersistentVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodsMetricSource) DeepCopyInto(out *PodsMetricSource) {
	*out = *in
	in.Metric.DeepCopyInto(&out.Metric)
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodsMetricSource.
func (in *PodsMetricSource) DeepCopy() *PodsMetricSource {
	if in == nil {
		return nil
	}
	out := new(PodsMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetricSource) DeepCopyInto(out *ResourceMetricSource) {
	*out = *in
	in.Target.DeepCopyInto(&out.Target)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetricSource.
func (in *ResourceMetricSource) DeepCopy() *ResourceMetricSource {
	if in == nil {
		return nil
	}
	out := new(ResourceMetricSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemInitialization) DeepCopyInto(out *SystemInitialization) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemInitialization.
func (in *SystemInitialization) DeepCopy() *SystemInitialization {
	if in == nil {
		return nil
	}
	out := new(SystemInitialization)
	in.DeepCopyInto(out)
	return out
}
