package cmd

import (
	"context"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jozuenoon/message_bus/collector"

	"github.com/spf13/cobra"
)

var trafficCmdDetectorIDs []string
var trafficCmdDeviceIDs []string
var trafficCmdRate int64

// trafficCmd represents the traffic command
var trafficCmd = &cobra.Command{
	Use:   "traffic",
	Short: "Traffic simulator",
	Run: func(cmd *cobra.Command, args []string) {
		ticker := time.NewTicker(time.Duration(trafficCmdRate) * time.Second)
		defer ticker.Stop()
		cli, err := NewCollectorClient(collectorHost)
		if err != nil {
			logger.Crit("failed to create collector client", "err", err)
			return
		}
	trafficCmdLoop:
		for {
			select {
			case <-signals:
				break trafficCmdLoop
			case <-ticker.C:
			}
			detector := randomElement(trafficCmdDetectorIDs)
			device := randomElement(trafficCmdDeviceIDs)
			now := time.Now()
			tnow := &timestamp.Timestamp{
				Seconds: now.Unix(),
				Nanos:   int32(now.Nanosecond()),
			}

			eventLog := &collector.EventLog{
				Loc: &collector.Coordinates{
					DetectorId: detector,
				},
				Events: []*collector.DetectionEvent{
					{
						Time: []*timestamp.Timestamp{tnow},
						DeviceId: &collector.DetectionEvent_Bluetooth{
							Bluetooth: device,
						},
					},
				},
			}

			_, err := cli.CreateEventLog(context.Background(), eventLog)
			if err != nil {
				logger.Crit("failed to create event log", "err", err)
				return
			}
			logger.Info("event", "detector", detector, "device", device, "time", now)
		}
	},
}

func init() {
	rootCmd.AddCommand(trafficCmd)

	trafficCmdDetectorIDs = *trafficCmd.PersistentFlags().StringArray("detector_ids",
		[]string{"xxx-1", "xxx-2", "xxx-3"}, "specify simulated detectors ids")

	trafficCmdDeviceIDs = *trafficCmd.PersistentFlags().StringArray("device_ids", []string{"ddd-1", "ddd-2", "ddd-3"},
		"specify simulated device ids")

	trafficCmdRate = *trafficCmd.PersistentFlags().Int64("rate", 5, "period of emitting random event in seconds")
}

func randomElement(s []string) string {
	return s[rand.Intn(len(s))]
}
