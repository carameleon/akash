// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
)

// MutatingWebhookConfigurationsGetter is an autogenerated mock type for the MutatingWebhookConfigurationsGetter type
type MutatingWebhookConfigurationsGetter struct {
	mock.Mock
}

// MutatingWebhookConfigurations provides a mock function with given fields:
func (_m *MutatingWebhookConfigurationsGetter) MutatingWebhookConfigurations() v1beta1.MutatingWebhookConfigurationInterface {
	ret := _m.Called()

	var r0 v1beta1.MutatingWebhookConfigurationInterface
	if rf, ok := ret.Get(0).(func() v1beta1.MutatingWebhookConfigurationInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.MutatingWebhookConfigurationInterface)
		}
	}

	return r0
}
