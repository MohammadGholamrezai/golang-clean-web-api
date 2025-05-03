package services

import (
	"time"

	"github.com/MohammadGholamrezai/golang-clean-web-api/api/dto"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/logging"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/service_errors"
	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	MobileNumber string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	return &TokenService{
		logger: logging.NewLogger(cfg),
		cfg:    cfg,
	}
}

func (t *TokenService) GenerateToken(payload *tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(t.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(t.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	atc := jwt.MapClaims{}
	atc["user_id"] = payload.UserId
	atc["mobile_number"] = payload.MobileNumber
	atc["roles"] = payload.Roles
	atc["exp"] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	var err error
	td.AccessToken, err = at.SignedString([]byte(t.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	rft := jwt.MapClaims{}
	rft["user_id"] = payload.UserId
	rft["exp"] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rft)
	td.RefreshToken, err = rt.SignedString([]byte(t.cfg.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (t *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	ac, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnexpectedErrors}
		}
		return []byte(t.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (t *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{} // empty interface {}
	verifyToken, err := t.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}

		return claimMap, nil
	}

	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimNotFound}
}
