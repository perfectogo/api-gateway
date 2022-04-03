package service

import (
	"fmt"

	"github.com/perfectogo/api-gateway/config"
	"github.com/perfectogo/api-gateway/genproto/catalog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type InterfaceServiceManager interface {
	CatalogService() catalog.CatalogServiceClient
}
type serviceManager struct {
	catalogService catalog.CatalogServiceClient
}

func (s *serviceManager) CatalogService() catalog.CatalogServiceClient {
	return s.catalogService
}

func NewServiceManager(cfg *config.Config) (InterfaceServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		catalogService: catalog.NewCatalogServiceClient(connCatalog),
	}

	return serviceManager, nil
}
