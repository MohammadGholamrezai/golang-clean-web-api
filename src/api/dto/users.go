package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpireTime  int64    `json:"access_token_expire_time"`
	RefreshTokenExpireTime int64    `json:"refresh_token_expire_time"`
}
