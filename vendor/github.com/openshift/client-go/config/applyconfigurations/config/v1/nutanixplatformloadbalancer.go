// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
)

// NutanixPlatformLoadBalancerApplyConfiguration represents a declarative configuration of the NutanixPlatformLoadBalancer type for use
// with apply.
type NutanixPlatformLoadBalancerApplyConfiguration struct {
	Type *v1.PlatformLoadBalancerType `json:"type,omitempty"`
}

// NutanixPlatformLoadBalancerApplyConfiguration constructs a declarative configuration of the NutanixPlatformLoadBalancer type for use with
// apply.
func NutanixPlatformLoadBalancer() *NutanixPlatformLoadBalancerApplyConfiguration {
	return &NutanixPlatformLoadBalancerApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *NutanixPlatformLoadBalancerApplyConfiguration) WithType(value v1.PlatformLoadBalancerType) *NutanixPlatformLoadBalancerApplyConfiguration {
	b.Type = &value
	return b
}
