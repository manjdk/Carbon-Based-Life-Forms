// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CallClientIFace is an autogenerated mock type for the CallClientIFace type
type CallClientIFace struct {
	mock.Mock
}

// Call provides a mock function with given fields: traceID, method, url, queryParams, body
func (_m *CallClientIFace) Call(traceID string, method string, url string, queryParams map[string]string, body []byte) ([]byte, int, error) {
	ret := _m.Called(traceID, method, url, queryParams, body)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]string, []byte) []byte); ok {
		r0 = rf(traceID, method, url, queryParams, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string, string, string, map[string]string, []byte) int); ok {
		r1 = rf(traceID, method, url, queryParams, body)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, map[string]string, []byte) error); ok {
		r2 = rf(traceID, method, url, queryParams, body)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
