package query

import (
	"context"
	"time"
)

type Event struct {
	DetectorID string
	DeviceID   string
	Time       time.Time
}

type Repository interface {
	GetEvents(ctx context.Context, detectors []string, after, before time.Time, limit int64) ([]*Event, error)
}
