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

func MustServe(config Config, services service.Services) {
	api.MustServe(config.Config, func(router *gin.Engine) {
		// set default rate limit config if not configured
		config.RateLimit.Add(middleware.RateLimitTierFree, "overall", 5, 20)
		// EIP-7702 tx
		config.RateLimit.Add(middleware.RateLimitTierFree, "setAuth", 1, 5)
		// gas tank
		config.RateLimit.Add(middleware.RateLimitTierFree, "signUserOp", 1, 5)
		// token pay
		config.RateLimit.Add(middleware.RateLimitTierFree, "getPrice", 1, 5)
		config.RateLimit.Add(middleware.RateLimitTierFree, "sponsor", 1, 5)

		// rate limit by client real IP address
		rateLimiters := middleware.NewRateLimitManager(config.RateLimit)
		go rateLimiters.ScheduleExpire()
		router.Use(rateLimiters.Middleware("overall"))

		// swagger docs
		docs.SwaggerInfo.BasePath = "/api"
		if config.SwaggerEnabled {
			router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		api := router.Group("api")

		// test client IP address
		api.GET("/ip", middleware.Wrap(testClientIP))

		// account abstract - auth
		if services.AccountAbstract != nil {
			controller := NewAccountAbstractController(services)
			api.POST("/aa/auth", rateLimiters.Middleware("setAuth"), middleware.Metrics("api.aa.auth.send"), middleware.Wrap(controller.SendAuth))
			api.GET("/aa/auth/:txHash", middleware.Metrics("api.aa.auth.status"), middleware.Wrap(controller.GetAuthStatus))
		}

		// verifying paymaster
		if services.VerifyingPaymaster != nil {
			controller := NewVerifyingPaymasterController(services)
			api.GET("/aa/paymaster/stub", middleware.Metrics("api.aa.paymaster.stub"), middleware.Wrap(controller.Stub))
			api.POST("/aa/paymaster/sign", rateLimiters.Middleware("signUserOp"), middleware.Metrics("api.aa.paymaster.sign"), middleware.Wrap(controller.Sign))
		}

		// Gas tank
		if services.GasTank != nil {
			controller := NewGasTankController(services)
			api.POST("/aa/gastank/stub/credit", middleware.Metrics("api.aa.gastank.stub.credit"), middleware.Wrap(controller.PrepareCredit))
			api.POST("/aa/gastank/stub/refund", middleware.Metrics("api.aa.gastank.stub.refund"), middleware.Wrap(controller.PrepareRefund))
			api.POST("/aa/gastank/sign", rateLimiters.Middleware("signUserOp"), middleware.Metrics("api.aa.gastank.sign"), middleware.Wrap(controller.Sign))
		}

		// token pay
		tokenPayController := NewTokenPayController(services)
		api.GET("/tokenpay/config", middleware.Metrics("api.tokenpay.config"), middleware.Wrap(tokenPayController.Config))
		api.GET("/tokenpay/price", rateLimiters.Middleware("getPrice"), middleware.Metrics("api.tokenpay.price"), middleware.Wrap(tokenPayController.GetETHPrice))
		api.POST("/tokenpay/submit", rateLimiters.Middleware("sponsor"), middleware.Metrics("api.tokenpay.submit"), middleware.Wrap(tokenPayController.Submit))
	})
}

func testClientIP(c *gin.Context) (any, error) {
	return middleware.GetRealIP(c), nil
}
