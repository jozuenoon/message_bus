// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import time "time"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateActiveDetectors provides a mock function with given fields: detectorID, activeDetectors, timestamp
func (_m *Repository) CreateActiveDetectors(detectorID string, activeDetectors string, timestamp time.Time) error {
	ret := _m.Called(detectorID, activeDetectors, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, time.Time) error); ok {
		r0 = rf(detectorID, activeDetectors, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBatteryVoltage provides a mock function with given fields: detectorID, voltage, timestamp
func (_m *Repository) CreateBatteryVoltage(detectorID string, voltage float64, timestamp time.Time) error {
	ret := _m.Called(detectorID, voltage, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64, time.Time) error); ok {
		r0 = rf(detectorID, voltage, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCoordinates provides a mock function with given fields: detectorID, latitude, longitude
func (_m *Repository) CreateCoordinates(detectorID string, latitude float64, longitude float64) error {
	ret := _m.Called(detectorID, latitude, longitude)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64, float64) error); ok {
		r0 = rf(detectorID, latitude, longitude)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDetectionCount provides a mock function with given fields: detectorID, detectorType, count, timestamp
func (_m *Repository) CreateDetectionCount(detectorID string, detectorType string, count int64, timestamp time.Time) {
	_m.Called(detectorID, detectorType, count, timestamp)
}

// CreateDetectionEvent provides a mock function with given fields: detectorID, deviceID, timestamp
func (_m *Repository) CreateDetectionEvent(detectorID string, deviceID string, timestamp time.Time) error {
	ret := _m.Called(detectorID, deviceID, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, time.Time) error); ok {
		r0 = rf(detectorID, deviceID, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
