// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "github.com/manjdk/Carbon-Based-Life-Forms/domain"
	mock "github.com/stretchr/testify/mock"
)

// MineralUpdateStateIFace is an autogenerated mock type for the MineralUpdateStateIFace type
type MineralUpdateStateIFace struct {
	mock.Mock
}

// Update provides a mock function with given fields: mineral
func (_m *MineralUpdateStateIFace) Update(mineral *domain.Mineral) error {
	ret := _m.Called(mineral)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Mineral) error); ok {
		r0 = rf(mineral)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}