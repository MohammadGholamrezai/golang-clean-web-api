package helper

import (
	"net/http"

	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	// Otp
	service_errors.OtpExists:   409,
	service_errors.OtpUsed:     410,
	service_errors.OtpNotValid: 400,

	// User
	service_errors.UserEmailAlreadyExists:    409,
	service_errors.UserUsernameAlreadyExists: 409,
	service_errors.ClaimNotFound:             404,
	service_errors.UnexpectedErrors:          500,
	service_errors.RecordNotFound:            404,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
