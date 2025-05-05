package dto

//  Data Transfer Object (DTO)

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpireTime  int64  `json:"access_token_expire_time"`
	RefreshTokenExpireTime int64  `json:"refresh_token_expire_time"`
}

type RegisterUserByUsernameRequest struct {
	FirstName string `json:"FirstName" binding:"required,min=3"`
	LastName  string `json:"LastName" binding:"required,min=6"`
	Username  string `json:"Username" binding:"required,min=5"`
	Email     string `json:"Email" binding:"email,min=6"`
	Password  string `json:"Password" binding:"required,password,min=6"`
}

type RegisterLoginByMobileRequest struct {
	MobileNumber string `json:"MobileNumber" binding:"required,mobile,min=11,max=11"`
	Otp          string `json:"Otp" binding:"required,min=6,max=6"`
}

type LoginByUsernameRequest struct {
	Username string `json:"Username" binding:"required,min=5"`
	Password string `json:"Password" binding:"required,password,min=6"`
}
