package collector

import "time"

var _ Repository = (*bboltRepository)(nil)

type bboltRepository struct {

}

func (r *bboltRepository) CreateDetectionEvent(detectorID string, deviceID string, timestamp time.Time) error {
	panic("implement me")
}

func (r *bboltRepository) CreateActiveDetectors(detectorID string, activeDetectors string, timestamp time.Time) error {
	panic("implement me")
}

func (r *bboltRepository) CreateBatteryVoltage(detectorID string, voltage float64, timestamp time.Time) error {
	panic("implement me")
}

func (r *bboltRepository) CreateDetectionCount(detectorID string, detectorType string, count int64, timestamp time.Time) {
	panic("implement me")
}

func (r *bboltRepository) CreateCoordinates(detectorID string, latitude, longitude float64) error {
	panic("implement me")
}






