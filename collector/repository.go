package collector

import "time"
import "github.com/jozuenoon/message_bus/pkg/types"

type Repository interface {
	CreateDetectionEvent(detectorID string, deviceID string, timestamp time.Time) error
	CreateDetectorLink(destDetectorID, srcDetectorID string, maxSeconds int64) error

	// NOTICE: not implemented
	CreateActiveDetectors(detectorID string, activeDetectors string, timestamp time.Time) error
	CreateBatteryVoltage(detectorID string, voltage float64, timestamp time.Time) error
	CreateDetectionCount(detectorID string, detectorType string, count int64, timestamp time.Time) error
	CreateCoordinates(detectorID string, latitude, longitude types.DecimalDegrees) error
}
