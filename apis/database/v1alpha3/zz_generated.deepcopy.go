// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServer) DeepCopyInto(out *MySQLServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServer.
func (in *MySQLServer) DeepCopy() *MySQLServer {
	if in == nil {
		return nil
	}
	out := new(MySQLServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerList) DeepCopyInto(out *MySQLServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerList.
func (in *MySQLServerList) DeepCopy() *MySQLServerList {
	if in == nil {
		return nil
	}
	out := new(MySQLServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerNameReferencer) DeepCopyInto(out *MySQLServerNameReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerNameReferencer.
func (in *MySQLServerNameReferencer) DeepCopy() *MySQLServerNameReferencer {
	if in == nil {
		return nil
	}
	out := new(MySQLServerNameReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerVirtualNetworkRule) DeepCopyInto(out *MySQLServerVirtualNetworkRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerVirtualNetworkRule.
func (in *MySQLServerVirtualNetworkRule) DeepCopy() *MySQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(MySQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerVirtualNetworkRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopyInto(out *MySQLServerVirtualNetworkRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLServerVirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerVirtualNetworkRuleList.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopy() *MySQLServerVirtualNetworkRuleList {
	if in == nil {
		return nil
	}
	out := new(MySQLServerVirtualNetworkRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLVirtualNetworkRuleSpec) DeepCopyInto(out *MySQLVirtualNetworkRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(MySQLServerNameReferencer)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(ResourceGroupNameReferencerForVirtualNetworkRule)
		**out = **in
	}
	in.VirtualNetworkRuleProperties.DeepCopyInto(&out.VirtualNetworkRuleProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLVirtualNetworkRuleSpec.
func (in *MySQLVirtualNetworkRuleSpec) DeepCopy() *MySQLVirtualNetworkRuleSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLVirtualNetworkRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServer) DeepCopyInto(out *PostgreSQLServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServer.
func (in *PostgreSQLServer) DeepCopy() *PostgreSQLServer {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerList) DeepCopyInto(out *PostgreSQLServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerList.
func (in *PostgreSQLServerList) DeepCopy() *PostgreSQLServerList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerNameReferencer) DeepCopyInto(out *PostgreSQLServerNameReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerNameReferencer.
func (in *PostgreSQLServerNameReferencer) DeepCopy() *PostgreSQLServerNameReferencer {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerNameReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopyInto(out *PostgreSQLServerVirtualNetworkRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerVirtualNetworkRule.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopy() *PostgreSQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopyInto(out *PostgreSQLServerVirtualNetworkRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLServerVirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerVirtualNetworkRuleList.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopy() *PostgreSQLServerVirtualNetworkRuleList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerVirtualNetworkRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLVirtualNetworkRuleSpec) DeepCopyInto(out *PostgreSQLVirtualNetworkRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(PostgreSQLServerNameReferencer)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(ResourceGroupNameReferencerForVirtualNetworkRule)
		**out = **in
	}
	in.VirtualNetworkRuleProperties.DeepCopyInto(&out.VirtualNetworkRuleProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLVirtualNetworkRuleSpec.
func (in *PostgreSQLVirtualNetworkRuleSpec) DeepCopy() *PostgreSQLVirtualNetworkRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLVirtualNetworkRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PricingTierSpec) DeepCopyInto(out *PricingTierSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PricingTierSpec.
func (in *PricingTierSpec) DeepCopy() *PricingTierSpec {
	if in == nil {
		return nil
	}
	out := new(PricingTierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupNameReferencerForSQLServer) DeepCopyInto(out *ResourceGroupNameReferencerForSQLServer) {
	*out = *in
	out.ResourceGroupNameReferencer = in.ResourceGroupNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupNameReferencerForSQLServer.
func (in *ResourceGroupNameReferencerForSQLServer) DeepCopy() *ResourceGroupNameReferencerForSQLServer {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupNameReferencerForSQLServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroupNameReferencerForVirtualNetworkRule) DeepCopyInto(out *ResourceGroupNameReferencerForVirtualNetworkRule) {
	*out = *in
	out.ResourceGroupNameReferencer = in.ResourceGroupNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroupNameReferencerForVirtualNetworkRule.
func (in *ResourceGroupNameReferencerForVirtualNetworkRule) DeepCopy() *ResourceGroupNameReferencerForVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(ResourceGroupNameReferencerForVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerClass) DeepCopyInto(out *SQLServerClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerClass.
func (in *SQLServerClass) DeepCopy() *SQLServerClass {
	if in == nil {
		return nil
	}
	out := new(SQLServerClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLServerClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerClassList) DeepCopyInto(out *SQLServerClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SQLServerClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerClassList.
func (in *SQLServerClassList) DeepCopy() *SQLServerClassList {
	if in == nil {
		return nil
	}
	out := new(SQLServerClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SQLServerClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerClassSpecTemplate) DeepCopyInto(out *SQLServerClassSpecTemplate) {
	*out = *in
	in.ClassSpecTemplate.DeepCopyInto(&out.ClassSpecTemplate)
	in.SQLServerParameters.DeepCopyInto(&out.SQLServerParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerClassSpecTemplate.
func (in *SQLServerClassSpecTemplate) DeepCopy() *SQLServerClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(SQLServerClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerParameters) DeepCopyInto(out *SQLServerParameters) {
	*out = *in
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(ResourceGroupNameReferencerForSQLServer)
		**out = **in
	}
	out.PricingTier = in.PricingTier
	out.StorageProfile = in.StorageProfile
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerParameters.
func (in *SQLServerParameters) DeepCopy() *SQLServerParameters {
	if in == nil {
		return nil
	}
	out := new(SQLServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerSpec) DeepCopyInto(out *SQLServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.SQLServerParameters.DeepCopyInto(&out.SQLServerParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerSpec.
func (in *SQLServerSpec) DeepCopy() *SQLServerSpec {
	if in == nil {
		return nil
	}
	out := new(SQLServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLServerStatus) DeepCopyInto(out *SQLServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLServerStatus.
func (in *SQLServerStatus) DeepCopy() *SQLServerStatus {
	if in == nil {
		return nil
	}
	out := new(SQLServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerNameReferencerForMySQLServerVirtualNetworkRule) DeepCopyInto(out *ServerNameReferencerForMySQLServerVirtualNetworkRule) {
	*out = *in
	out.MySQLServerNameReferencer = in.MySQLServerNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerNameReferencerForMySQLServerVirtualNetworkRule.
func (in *ServerNameReferencerForMySQLServerVirtualNetworkRule) DeepCopy() *ServerNameReferencerForMySQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(ServerNameReferencerForMySQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerNameReferencerForPostgreSQLServerVirtualNetworkRule) DeepCopyInto(out *ServerNameReferencerForPostgreSQLServerVirtualNetworkRule) {
	*out = *in
	out.PostgreSQLServerNameReferencer = in.PostgreSQLServerNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerNameReferencerForPostgreSQLServerVirtualNetworkRule.
func (in *ServerNameReferencerForPostgreSQLServerVirtualNetworkRule) DeepCopy() *ServerNameReferencerForPostgreSQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(ServerNameReferencerForPostgreSQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfileSpec) DeepCopyInto(out *StorageProfileSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfileSpec.
func (in *StorageProfileSpec) DeepCopy() *StorageProfileSpec {
	if in == nil {
		return nil
	}
	out := new(StorageProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetIDReferencerForVirtualNetworkRule) DeepCopyInto(out *SubnetIDReferencerForVirtualNetworkRule) {
	*out = *in
	out.SubnetIDReferencer = in.SubnetIDReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetIDReferencerForVirtualNetworkRule.
func (in *SubnetIDReferencerForVirtualNetworkRule) DeepCopy() *SubnetIDReferencerForVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(SubnetIDReferencerForVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleProperties) DeepCopyInto(out *VirtualNetworkRuleProperties) {
	*out = *in
	if in.VirtualNetworkSubnetIDRef != nil {
		in, out := &in.VirtualNetworkSubnetIDRef, &out.VirtualNetworkSubnetIDRef
		*out = new(SubnetIDReferencerForVirtualNetworkRule)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleProperties.
func (in *VirtualNetworkRuleProperties) DeepCopy() *VirtualNetworkRuleProperties {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleStatus) DeepCopyInto(out *VirtualNetworkRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleStatus.
func (in *VirtualNetworkRuleStatus) DeepCopy() *VirtualNetworkRuleStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleStatus)
	in.DeepCopyInto(out)
	return out
}
