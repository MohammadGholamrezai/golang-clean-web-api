package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors

	if errors.As(err, &ve) { // اروری که از ورودی اومده رو بریز توی &ve
		// در واقع یعنی اگر از نوع ولیدیشن اررور بود
		for _, err := range err.(validator.ValidationErrors) { // اروری که از ورودی اومد تبدیل بشه به validator.ValidationErrors
			var el ValidationError
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, el)
		}

		return &validationErrors
	}

	return nil
}
