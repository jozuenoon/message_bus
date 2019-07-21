package cmd

import (
	"context"
	"os"
	"sort"
	"time"

	"github.com/jozuenoon/message_bus/query"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var eventsDetectorsIDs []string
var eventsLimit int64

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "List latest events.",
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := NewQueryClient(queryHost)
		if err != nil {
			logger.Crit("failed to create collector client", "err", err)
			return
		}
		req := &query.GetEventsRequest{
			DetectorIds: eventsDetectorsIDs,
			Limit:       eventsLimit,
		}
		resp, err := cli.GetEvents(context.Background(), req)
		if err != nil {
			logger.Crit("failed to get events", "err", err)
			return
		}

		// Print out events.
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Detector ID", "Device ID", "Event time"})
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)

		for _, det := range resp.Events {
			bulk := make([][]string, 0, len(det.Time))
			for _, ev := range det.Time {
				ts := time.Unix(ev.Seconds, int64(ev.Nanos))
				bulk = append(bulk, []string{det.DetectorId, det.DeviceId, ts.Format(time.RFC3339)})
			}
			sort.Slice(bulk, func(i, j int) bool {
				return bulk[i][2] > bulk[j][2]
			})
			table.AppendBulk(bulk)

		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)

	eventsCmd.PersistentFlags().StringArrayVar(&eventsDetectorsIDs, "detector_ids",
		[]string{"xxx-1"}, "filtered detectors ids")

	eventsCmd.PersistentFlags().Int64Var(&eventsLimit, "event_limit", 1,
		"events limit")
}
