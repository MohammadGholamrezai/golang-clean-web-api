package services

import (
	"fmt"

	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/constants"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/cache"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/logging"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/service_errors"
	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redis  *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, cfg: cfg, redis: redis}
}

func (s *OtpService) SetOtp(MobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, MobileNumber)
	val := &OtpDto{Value: otp, Used: false}

	res, err := cache.Get[OtpDto](s.redis, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{
			EndUserMessage: service_errors.OtpExists,
		}
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{
			EndUserMessage: service_errors.OtpUsed,
		}
	}

	err = cache.Set(s.redis, key, val, s.cfg.Otp.ExpireTime)
	if err != nil {
		return err
	}

	return nil
}
func (s *OtpService) ValidateOtp(MobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, MobileNumber)
	res, err := cache.Get[OtpDto](s.redis, key)

	if err != nil {
		return err
	} else if res.Used {
		return &service_errors.ServiceError{
			EndUserMessage: service_errors.OtpUsed,
		}

	} else if !res.Used && res.Value != otp {
		return &service_errors.ServiceError{
			EndUserMessage: service_errors.OtpNotValid,
		}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redis, key, res, s.cfg.Otp.ExpireTime)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}
