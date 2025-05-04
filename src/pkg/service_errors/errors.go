package service_errors

const (
	UnexpectedErrors          string = "Unexpected errors"
	ClaimNotFound             string = "Claim not found"
	UserEmailAlreadyExists    string = "User email already exists"
	UserUsernameAlreadyExists string = "User username already exists"
	TokenNotFound             string = "Token not found"
	TokenExpired              string = "Token expired"
	TokenInvalid              string = "Token invalid"

	OtpExists   string = "Otp exists"
	OtpUsed     string = "Otp used"
	OtpNotValid string = "Otp not valid"
)
