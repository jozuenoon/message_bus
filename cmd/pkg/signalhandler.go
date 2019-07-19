package pkg

import (
	"context"
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/oklog/run"
	"os"
	"os/signal"
	"syscall"
)

var shutdownSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

func SignalHandler(_ context.Context, g *run.Group, logger log15.Logger) {
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, shutdownSignals...)
	g.Add(func() error {
		for sig := range signals {
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				logger.Info("interrupt", "sig", sig)
				return fmt.Errorf("terminated: %v", sig)
			}
		}
		return nil
	}, func(error) {})
}
