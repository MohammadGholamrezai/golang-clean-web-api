package api

import (
	"fmt"

	"github.com/MohammadGholamrezai/golang-clean-web-api/api/middlewares"
	"github.com/MohammadGholamrezai/golang-clean-web-api/api/routers"
	"github.com/MohammadGholamrezai/golang-clean-web-api/api/validations"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	registerValidators()

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	registerRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func registerValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validations.PasswordValidator, true)
	}
}

func registerRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		// middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"})
	
		health := v1.Group("/health")
		users := v1.Group("/users")
		countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		routers.Health(health)
		routers.User(users, cfg)
		routers.Country(countries, cfg)
	}
}
