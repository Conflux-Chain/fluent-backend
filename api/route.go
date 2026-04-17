package api

import (
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/api/middleware"
	"github.com/gin-gonic/gin"
)

// TODO swagger doc

type Config struct {
	api.Config `mapstructure:",squash"`
}

func MustServe(config Config, services service.Services) {
	api.MustServe(config.Config, func(router *gin.Engine) {
		api := router.Group("api")

		aaController := NewAccountAbstractController(services)

		// account abstract
		api.POST("/aa/auth", middleware.Metrics("api.aa.auth.send"), middleware.Wrap(aaController.SendAuth))
		api.GET("/aa/auth/:txhash", middleware.Metrics("api.aa.auth.status"), middleware.Wrap(aaController.GetAuthStatus))
	})
}
