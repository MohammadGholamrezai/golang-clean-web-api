package middlewares

import (
	"net/http"
	"strings"

	"github.com/MohammadGholamrezai/golang-clean-web-api/api/helper"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/constants"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/service_errors"
	"github.com/MohammadGholamrezai/golang-clean-web-api/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)

	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}

		auth := c.GetHeader("Authorization")
		token := strings.Split(auth, " ")
		if auth == " " {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenNotFound}
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}

		if err != nil {
			// c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(nil, false, -2, err))
			return
		}

		c.Set("user_id", claimMap["user_id"])
		c.Set("mobile_number", claimMap["mobile_number"])
		c.Set("roles", claimMap["roles"])
		c.Set("exp", claimMap["exp"])

		c.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {

			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -300))
			return
		}

		// Get roles from token
		rolesVal := c.Keys[constants.RolesKey]
		if rolesVal == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -301))
			return
		}

		// Insert roles to dictionary
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}

		// Checking validRoles exists in dictionary or not
		for _, role := range validRoles {
			if _, exists := val[role]; exists {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -302))
	}
}
