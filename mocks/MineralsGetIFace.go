// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/manjdk/Carbon-Based-Life-Forms/domain"
	mock "github.com/stretchr/testify/mock"
)

// MineralsGetIFace is an autogenerated mock type for the MineralsGetIFace type
type MineralsGetIFace struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *MineralsGetIFace) GetAll() ([]domain.Mineral, error) {
	ret := _m.Called()

	var r0 []domain.Mineral
	if rf, ok := ret.Get(0).(func() []domain.Mineral); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Mineral)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
