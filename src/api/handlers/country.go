package handlers

import (
	"net/http"
	"strconv"

	"github.com/MohammadGholamrezai/golang-clean-web-api/api/dto"
	"github.com/MohammadGholamrezai/golang-clean-web-api/api/helper"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/services"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{
		service: services.NewCountryService(cfg),
	}
}

func (h *CountryHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	response, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, 131, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(response, true, 0))
}

func (h *CountryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := dto.CreateUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, 121, err))
		return
	}

	response, err := h.service.Update(c, &req, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, 131, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

func (h *CountryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	err := h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, 131, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, 0))
}

func (h *CountryHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponse(nil, false, 141))
		return
	}

	response, err := h.service.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, 131, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(response, true, 0))
}

// TODO
func (h *CountryHandler) getByFilter(c *gin.Context) {

}
