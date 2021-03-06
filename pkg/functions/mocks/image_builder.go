// Code generated by mockery v1.0.0
package mocks

import functions "github.com/vmware/dispatch/pkg/functions"
import mock "github.com/stretchr/testify/mock"

// ImageBuilder is an autogenerated mock type for the ImageBuilder type
type ImageBuilder struct {
	mock.Mock
}

// BuildImage provides a mock function with given fields: faas, fnID, e
func (_m *ImageBuilder) BuildImage(faas string, fnID string, e *functions.Exec) (string, error) {
	ret := _m.Called(faas, fnID, e)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, *functions.Exec) string); ok {
		r0 = rf(faas, fnID, e)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *functions.Exec) error); ok {
		r1 = rf(faas, fnID, e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
