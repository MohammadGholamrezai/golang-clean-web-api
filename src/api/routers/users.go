package routers

import (
	"github.com/MohammadGholamrezai/golang-clean-web-api/api/handlers"
	"github.com/MohammadGholamrezai/golang-clean-web-api/api/middlewares"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
}
