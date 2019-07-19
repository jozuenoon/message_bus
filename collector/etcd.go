package collector

import "time"
import "github.com/jozuenoon/message_bus/pkg/types"

var _ Repository = (*etcdRepository)(nil)

type etcdRepository struct {

}

func (r *etcdRepository) CreateDetectionEvent(detectorID string, deviceID string, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateActiveDetectors(detectorID string, activeDetectors string, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateBatteryVoltage(detectorID string, voltage float64, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateDetectionCount(detectorID string, detectorType string, count int64, timestamp time.Time) error {
	panic("implement me")
}

func (r *etcdRepository) CreateCoordinates(detectorID string, latitude, longitude types.DecimalDegrees) error {
	panic("implement me")
}






