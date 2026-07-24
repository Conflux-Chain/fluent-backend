package api

import (
	"github.com/Conflux-Chain/fluent-backend/docs"
	"github.com/Conflux-Chain/fluent-backend/service"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/Conflux-Chain/go-conflux-util/api/middleware"
	"github.com/Conflux-Chain/go-conflux-util/api/middleware/rate"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	api.Config `mapstructure:",squash"`

	SwaggerEnabled bool

	RateLimit rate.Config
}

// @title           Fluent Backend API
// @version         1.0
// @description     REST API for upgrading an EOA (Externally Owned Account) to an EIP-7702 Account Abstraction smart-contract wallet.
// @BasePath /api
func MustServe(config Config, services service.Services) {
	api.MustServe(config.Config, func(router *gin.Engine) {
		// set default rate limit config if not configured
		config.RateLimit.Add(middleware.RateLimitTierFree, "overall", 10, 20)
		// EIP-7702 tx
		config.RateLimit.Add(middleware.RateLimitTierFree, "setAuth", 1, 5)
		// gas tank
		config.RateLimit.Add(middleware.RateLimitTierFree, "signUserOp", 1, 5)
		// token pay
		config.RateLimit.Add(middleware.RateLimitTierFree, "getPrice", 5, 10)
		config.RateLimit.Add(middleware.RateLimitTierFree, "sponsor", 1, 5)

		// rate limit by client real IP address
		rateLimiters := middleware.NewRateLimitManager(config.RateLimit, middleware.DefaultRateLimitKeyExtractor)
		go rateLimiters.ScheduleExpire()
		router.Use(rateLimiters.Middleware("overall"))

		// swagger docs
		docs.SwaggerInfo.BasePath = "/api"
		if config.SwaggerEnabled {
			router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		api := router.Group("api")

		aaController := NewAccountAbstractController(services)
		gasTankController := NewGasTankController(services)
		tokenPayController := NewTokenPayController(services)

		// account abstract - auth
		api.POST("/aa/auth", rateLimiters.Middleware("setAuth"), middleware.Metrics("api.aa.auth.send"), middleware.Wrap(aaController.SendAuth))
		api.GET("/aa/auth/:txHash", middleware.Metrics("api.aa.auth.status"), middleware.Wrap(aaController.GetAuthStatus))

		// Gas tank
		api.POST("/aa/gastank/prepare/credit", middleware.Metrics("api.aa.gastank.prepare.credit"), middleware.Wrap(gasTankController.PrepareCredit))
		api.POST("/aa/gastank/prepare/refund", middleware.Metrics("api.aa.gastank.prepare.refund"), middleware.Wrap(gasTankController.PrepareRefund))
		api.POST("/aa/gastank/sign", rateLimiters.Middleware("signUserOp"), middleware.Metrics("api.aa.gastank.signature"), middleware.Wrap(gasTankController.Sign))

		// token pay
		api.GET("/tokenpay/config", middleware.Metrics("api.tokenpay.config"), middleware.Wrap(tokenPayController.Config))
		api.GET("/tokenpay/price", rateLimiters.Middleware("getPrice"), middleware.Metrics("api.tokenpay.price"), middleware.Wrap(tokenPayController.GetETHPrice))
		api.POST("/tokenpay/submit", rateLimiters.Middleware("sponsor"), middleware.Metrics("api.tokenpay.submit"), middleware.Wrap(tokenPayController.Submit))
	})
}
