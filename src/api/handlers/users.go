package handlers

import (
	"net/http"

	"github.com/MohammadGholamrezai/golang-clean-web-api/api/dto"
	"github.com/MohammadGholamrezai/golang-clean-web-api/api/helper"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/services"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	return &UsersHandler{
		service: services.NewUserService(cfg),
	}
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UsersHandler) SendOtp(c *gin.Context) {
	otp := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&otp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	if err = h.service.SendOtp(otp); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.GenerateBaseResponseWithError(err, false, -1, err))
		return
	}

	// call sms service to send otp
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}
