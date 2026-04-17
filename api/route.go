package api

import (
	"github.com/Conflux-Chain/fluent-backend/docs"
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	api.Config `mapstructure:",squash"`

	SwaggerEnabled bool
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server for demonstrating Swagger in Golang.
// @BasePath /api
func MustServe(config Config, services service.Services) {
	api.MustServe(config.Config, func(router *gin.Engine) {
		// swagger docs
		docs.SwaggerInfo.BasePath = "/api"
		if config.SwaggerEnabled {
			router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		api := router.Group("api")

		aaController := NewAccountAbstractController(services)

		// account abstract
		api.POST("/aa/auth", middleware.Metrics("api.aa.auth.send"), middleware.Wrap(aaController.SendAuth))
		api.GET("/aa/auth/:txHash", middleware.Metrics("api.aa.auth.status"), middleware.Wrap(aaController.GetAuthStatus))
	})
}
