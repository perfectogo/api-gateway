package main

import (
	"github.com/perfectogo/api-gateway/api"

	"github.com/perfectogo/api-gateway/config"
	"github.com/perfectogo/api-gateway/pkg/logger"
	"github.com/perfectogo/api-gateway/service"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := service.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
