// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import types "github.com/jozuenoon/message_bus/pkg/types"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetDetectors provides a mock function with given fields: latitude, longitude, radius
func (_m *Repository) GetDetectors(latitude types.DecimalDegrees, longitude types.DecimalDegrees, radius int64) {
	_m.Called(latitude, longitude, radius)
}
