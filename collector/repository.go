package collector

import (
	"context"
	"time"
)
import "github.com/jozuenoon/message_bus/pkg/types"

type Repository interface {
	CreateDetectionEvent(ctx context.Context, detectorID string, deviceID string, timestamp time.Time) error
	CreateDetectorLink(ctx context.Context, destDetectorID, srcDetectorID string, maxSeconds int64) error

	// NOTICE: not implemented
	CreateActiveDetectors(ctx context.Context, detectorID string, activeDetectors string, timestamp time.Time) error
	CreateBatteryVoltage(ctx context.Context, detectorID string, voltage float64, timestamp time.Time) error
	CreateDetectionCount(ctx context.Context, detectorID string, detectorType string, count int64, timestamp time.Time) error
	CreateCoordinates(ctx context.Context, detectorID string, latitude, longitude types.DecimalDegrees) error
}
