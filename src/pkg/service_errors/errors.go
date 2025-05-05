package service_errors

const (
	// Token
	UnexpectedErrors string = "Unexpected errors"
	ClaimNotFound    string = "Claim not found"
	TokenNotFound    string = "Token not found"
	TokenExpired     string = "Token expired"
	TokenInvalid     string = "Token invalid"

	// User
	UserEmailAlreadyExists    string = "User email already exists"
	UserUsernameAlreadyExists string = "User username already exists"

	// OTP
	OtpExists   string = "Otp exists"
	OtpUsed     string = "Otp used"
	OtpNotValid string = "Otp not valid"

	// DB
	RecordNotFound string = "Record not found"
)
