// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
)

// IngressStatusApplyConfiguration represents a declarative configuration of the IngressStatus type for use
// with apply.
type IngressStatusApplyConfiguration struct {
	ComponentRoutes  []ComponentRouteStatusApplyConfiguration `json:"componentRoutes,omitempty"`
	DefaultPlacement *configv1.DefaultPlacement               `json:"defaultPlacement,omitempty"`
}

// IngressStatusApplyConfiguration constructs a declarative configuration of the IngressStatus type for use with
// apply.
func IngressStatus() *IngressStatusApplyConfiguration {
	return &IngressStatusApplyConfiguration{}
}

// WithComponentRoutes adds the given value to the ComponentRoutes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ComponentRoutes field.
func (b *IngressStatusApplyConfiguration) WithComponentRoutes(values ...*ComponentRouteStatusApplyConfiguration) *IngressStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithComponentRoutes")
		}
		b.ComponentRoutes = append(b.ComponentRoutes, *values[i])
	}
	return b
}

// WithDefaultPlacement sets the DefaultPlacement field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultPlacement field is set to the value of the last call.
func (b *IngressStatusApplyConfiguration) WithDefaultPlacement(value configv1.DefaultPlacement) *IngressStatusApplyConfiguration {
	b.DefaultPlacement = &value
	return b
}
