package services

import (
	"github.com/MohammadGholamrezai/golang-clean-web-api/api/dto"
	"github.com/MohammadGholamrezai/golang-clean-web-api/common"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/db"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	database   *gorm.DB
	otpService *OtpService
}

func NewUserService(cfg *config.Config) *UserService {
	return &UserService{
		logger:     logging.NewLogger(cfg),
		cfg:        cfg,
		database:   db.GetDb(),
		otpService: NewOtpService(cfg),
	}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}

	return nil
}
