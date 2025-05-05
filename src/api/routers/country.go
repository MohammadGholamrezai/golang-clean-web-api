package routers

import (
	"github.com/MohammadGholamrezai/golang-clean-web-api/api/handlers"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/gin-gonic/gin"
)

func Country(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewCountryHandler(cfg)

	router.POST("/", h.Create)
	router.PUT("/:id", h.Update)
	router.DELETE("/:id", h.Delete)
	router.GET("/:id", h.GetById)
}
