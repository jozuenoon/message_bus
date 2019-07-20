// +build integration

package collector

import (
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/inconshreveable/log15"
	"github.com/jozuenoon/message_bus/pkg/types"
)

func Test_etcdRepository_CreateDetectionEvent(t *testing.T) {
	type fields struct {
		prefix string
		logger log15.Logger
		cli    *clientv3.Client
	}
	type args struct {
		detectorID string
		deviceID   string
		timestamp  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &etcdRepository{
				prefix: tt.fields.prefix,
				logger: tt.fields.logger,
				cli:    tt.fields.cli,
			}
			if err := r.CreateDetectionEvent(tt.args.detectorID, tt.args.deviceID, tt.args.timestamp); (err != nil) != tt.wantErr {
				t.Errorf("etcdRepository.CreateDetectionEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_etcdRepository_CreateActiveDetectors(t *testing.T) {
	type fields struct {
		prefix string
		logger log15.Logger
		cli    *clientv3.Client
	}
	type args struct {
		detectorID      string
		activeDetectors string
		timestamp       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &etcdRepository{
				prefix: tt.fields.prefix,
				logger: tt.fields.logger,
				cli:    tt.fields.cli,
			}
			if err := r.CreateActiveDetectors(tt.args.detectorID, tt.args.activeDetectors, tt.args.timestamp); (err != nil) != tt.wantErr {
				t.Errorf("etcdRepository.CreateActiveDetectors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_etcdRepository_CreateBatteryVoltage(t *testing.T) {
	type fields struct {
		prefix string
		logger log15.Logger
		cli    *clientv3.Client
	}
	type args struct {
		detectorID string
		voltage    float64
		timestamp  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &etcdRepository{
				prefix: tt.fields.prefix,
				logger: tt.fields.logger,
				cli:    tt.fields.cli,
			}
			if err := r.CreateBatteryVoltage(tt.args.detectorID, tt.args.voltage, tt.args.timestamp); (err != nil) != tt.wantErr {
				t.Errorf("etcdRepository.CreateBatteryVoltage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_etcdRepository_CreateDetectionCount(t *testing.T) {
	type fields struct {
		prefix string
		logger log15.Logger
		cli    *clientv3.Client
	}
	type args struct {
		detectorID   string
		detectorType string
		count        int64
		timestamp    time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &etcdRepository{
				prefix: tt.fields.prefix,
				logger: tt.fields.logger,
				cli:    tt.fields.cli,
			}
			if err := r.CreateDetectionCount(tt.args.detectorID, tt.args.detectorType, tt.args.count, tt.args.timestamp); (err != nil) != tt.wantErr {
				t.Errorf("etcdRepository.CreateDetectionCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_etcdRepository_CreateCoordinates(t *testing.T) {
	type fields struct {
		prefix string
		logger log15.Logger
		cli    *clientv3.Client
	}
	type args struct {
		detectorID string
		latitude   types.DecimalDegrees
		longitude  types.DecimalDegrees
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &etcdRepository{
				prefix: tt.fields.prefix,
				logger: tt.fields.logger,
				cli:    tt.fields.cli,
			}
			if err := r.CreateCoordinates(tt.args.detectorID, tt.args.latitude, tt.args.longitude); (err != nil) != tt.wantErr {
				t.Errorf("etcdRepository.CreateCoordinates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
