// Code generated by mockery v2.19.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// StatusClient is an autogenerated mock type for the StatusClient type
type StatusClient struct {
	mock.Mock
}

// Status provides a mock function with given fields:
func (_m *StatusClient) Status() client.SubResourceWriter {
	ret := _m.Called()

	var r0 client.SubResourceWriter
	if rf, ok := ret.Get(0).(func() client.SubResourceWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SubResourceWriter)
		}
	}

	return r0
}

type mockConstructorTestingTNewStatusClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatusClient creates a new instance of StatusClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatusClient(t mockConstructorTestingTNewStatusClient) *StatusClient {
	mock := &StatusClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
