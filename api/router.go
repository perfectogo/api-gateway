package api

import (
	v1 "github.com/perfectogo/api-gateway/api/handler/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/perfectogo/api-gateway/api/docs" // swag
	"github.com/perfectogo/api-gateway/config"
	"github.com/perfectogo/api-gateway/pkg/logger"
	"github.com/perfectogo/api-gateway/service"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager service.InterfaceServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")

	api.POST("/authors", handlerV1.PostAuthor)
	api.GET("/authors", handlerV1.GetAuthors)
	api.GET("/authors/:id", handlerV1.GetAuthor)
	api.PUT("/authors/:id", handlerV1.UpdateAuthor)
	api.DELETE("/authors/:id", handlerV1.DeleteAuthor)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
