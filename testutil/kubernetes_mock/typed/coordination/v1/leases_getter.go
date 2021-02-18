// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
)

// LeasesGetter is an autogenerated mock type for the LeasesGetter type
type LeasesGetter struct {
	mock.Mock
}

// Leases provides a mock function with given fields: namespace
func (_m *LeasesGetter) Leases(namespace string) v1.LeaseInterface {
	ret := _m.Called(namespace)

	var r0 v1.LeaseInterface
	if rf, ok := ret.Get(0).(func(string) v1.LeaseInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1.LeaseInterface)
		}
	}

	return r0
}
