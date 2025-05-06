package services

import (
	"context"

	"github.com/MohammadGholamrezai/golang-clean-web-api/api/dto"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/db"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/models"
	"gorm.io/gorm"
)

type CountryService struct {
	database *gorm.DB
	// base is new for generic crud
	base *BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	// return &CountryService{
	// 	database: db.GetDb(),
	// }
	return &CountryService{
		base: &BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryResponse]{
			Database: db.GetDb(),
		},
	}
}

func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CreateUpdateCountryResponse, error) {
	// In Go when using context.Context or json and getting data by interface{} type, usually parse as float64
	// UTC (Coordinated Universal Time)

	return s.base.Create(ctx, req)
}

func (s *CountryService) Update(ctx context.Context, req *dto.CreateUpdateCountryRequest, id int) (*dto.CreateUpdateCountryResponse, error) {
	return s.base.Update(ctx, req, id)
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CreateUpdateCountryResponse, error) {
	return s.base.GetById(ctx, id)
}
