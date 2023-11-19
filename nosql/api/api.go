package api

import (
	"nosql/config"
	"nosql/pkg/logger"
	"nosql/storage"

	v1 "nosql/api/handler"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Config  config.Config
	Log     logger.Logger
	Storage storage.StorageI
}

func New(ro RouterOptions) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Storage: ro.Storage,
		Logger:  ro.Log,
		Cfg:     ro.Config,
	})
	// Category endpoints
	router.GET("/v1/product", handlerV1.GetAllProducts)
	router.GET("/v1/product/:product_id", handlerV1.GetProduct)
	router.POST("/v1/product", handlerV1.CreateProduct)
	router.PUT("/v1/product/:product_id", handlerV1.UpdateProduct)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
