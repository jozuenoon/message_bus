package collector

import "time"

type Repository interface {
	CreateDetectionEvent(detectorID string, deviceID string, timestamp time.Time) error
	CreateActiveDetectors(detectorID string, activeDetectors string, timestamp time.Time) error
	CreateBatteryVoltage(detectorID string, voltage float64, timestamp time.Time) error
	CreateDetectionCount(detectorID string, detectorType string, count int64, timestamp time.Time)
	CreateCoordinates(detectorID string, latitude, longitude float64) error
}