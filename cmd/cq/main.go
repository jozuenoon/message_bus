package main

import (
	"context"

	log "github.com/inconshreveable/log15"

	"github.com/jozuenoon/message_bus/cmd/cq/cqserver"
	"github.com/stevenroose/gonfig"
)

var config = struct {
	CollectorPort   string `id:"collector_port" desc:"grpc collector port" default:":9000"`
	QueryPort       string `id:"query_port" desc:"grpc query port" default:":8000"`
	HealthCheckPort string `id:"healthcheck_port" desc:"grpc health check port" default:":5000"`
	ETCD            struct {
		Prefix    string   `id:"prefix" desc:"etcd app prefix" default:"tdc"`
		Endpoints []string `id:"endpoints" desc:"etcd endpoints" default:"http://127.0.0.1:2379"`
	} `id:"etcd" desc:"ETCD variables"`

	ConfigFile string `id:"config_file"`
}{}

func main() {
	logger := log.New("app", "cqserver")
	if err := gonfig.Load(&config, gonfig.Conf{
		ConfigFileVariable:  "config_file",
		FileDefaultFilename: "config/config.yaml",
		FileDecoder:         gonfig.DecoderYAML,
	}); err != nil {
		logger.Crit("failed to parse config", "err", err)
	}
	srv := cqserver.New(config.CollectorPort, config.QueryPort, config.HealthCheckPort, config.ETCD.Prefix, config.ETCD.Endpoints, logger)

	logger.Info("starting server",
		"collector_port", config.CollectorPort,
		"query_port", config.QueryPort,
		"healthcheck_port", config.HealthCheckPort)

	err := srv.Run(context.Background())
	if err != nil {
		logger.Crit("exit", "err", err, "msg", "server died")
	}
}
