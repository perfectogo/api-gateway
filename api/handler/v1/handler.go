package v1

import (
	"github.com/perfectogo/api-gateway/config"
	"github.com/perfectogo/api-gateway/pkg/logger"
	"github.com/perfectogo/api-gateway/service"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager service.InterfaceServiceManager
	cfg            config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager service.InterfaceServiceManager
	Cfg            config.Config
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}
