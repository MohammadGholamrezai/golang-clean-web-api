package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/MohammadGholamrezai/golang-clean-web-api/api/dto"
	"github.com/MohammadGholamrezai/golang-clean-web-api/config"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/db"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/models"
	"gorm.io/gorm"
)

type CountryService struct {
	database *gorm.DB
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		database: db.GetDb(),
	}
}

func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CreateUpdateCountryResponse, error) {
	// In Go when using context.Context or json and getting data by interface{} type, usually parse as float64
	// UTC (Coordinated Universal Time)

	country := models.Country{Name: req.Name}
	country.CreatedBy = int(ctx.Value("user_id").(float64))
	country.CreatedAt = time.Now().UTC()

	tx := s.database.WithContext(ctx).Begin()
	if err := tx.Create(&country).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	dto := &dto.CreateUpdateCountryResponse{Name: country.Name, Id: country.Id}
	return dto, nil
}

func (s *CountryService) Update(ctx context.Context, req *dto.CreateUpdateCountryRequest, id int) (*dto.CreateUpdateCountryResponse, error) {
	updateMap := map[string]interface{}{
		"Name":       req.Name,
		"ModifiedBy": &sql.NullInt64{Int64: int64(ctx.Value("user_id").(float64)), Valid: true},
		"ModifiedAt": &sql.NullTime{Valid: true, Time: time.Now().UTC()},
	}

	tx := s.database.WithContext(ctx).Begin()
	err := tx.Model(&models.Country{}).
		Where("id = ?", id).
		Updates(updateMap).
		Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	country := &models.Country{}
	err = tx.Model(&models.Country{}).
		Where("id = ? AND deleted_by is null", id).
		Find(&country).
		Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	dto := &dto.CreateUpdateCountryResponse{Name: country.Name, Id: country.Id}
	return dto, nil
}

func (s *CountryService) Delete(ctx context.Context, id int) error {
	updateMap := map[string]interface{}{
		"DeletedBy": &sql.NullInt64{Int64: int64(ctx.Value("user_id").(float64)), Valid: true},
		"DeletedAt": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	tx := s.database.WithContext(ctx).Begin()
	err := tx.Model(&models.Country{}).
		Where("id = ?", id).
		Updates(updateMap). // soft delete
		Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *CountryService) GetById(ctx context.Context, id int) (*dto.GetByIdResponse, error) {
	country := &models.Country{}

	err := s.database.
		Model(&models.Country{}).
		Where("id = ?", id).
		Find(&country).
		Error

	if err != nil {
		return nil, err
	}

	dto := &dto.GetByIdResponse{Id: country.Id, Name: country.Name}
	return dto, nil
}
