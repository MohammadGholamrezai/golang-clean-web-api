package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/MohammadGholamrezai/golang-clean-web-api/common"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/db"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/service_errors"
	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *gorm.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, _ := common.TypeConvertor[T](req)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.Create(model).
		Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return common.TypeConvertor[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, req *Tu, id int) (*Tr, error) {
	updateMap, _ := common.TypeConvertor[map[string]interface{}](req)
	(*updateMap)["ModifiedBy"] = &sql.NullInt64{Int64: int64(ctx.Value("user_id").(float64)), Valid: true}
	(*updateMap)["ModifiedAt"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}

	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.Model(model).
		Where("id = ? AND deleted_by is null", id).
		Updates(*updateMap).
		Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return s.GetById(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	deleteMap := map[string]interface{}{
		"DeletedBy": &sql.NullInt64{Int64: int64(ctx.Value("user_id").(float64)), Valid: true},
		"DeletedAt": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	deleteMap["ModifiedBy"] = &sql.NullInt64{Int64: int64(ctx.Value("user_id").(float64)), Valid: true}
	deleteMap["ModifiedAt"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}

	model := new(T)

	tx := s.Database.WithContext(ctx).Begin()

	if count := tx.
		Model(model).
		Where("id = ? AND deleted_by is null", id).
		Updates(deleteMap).
		RowsAffected; count == 0 {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}

	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	err := s.Database.
		Model(model).
		Where("id = ? AND deleted_by is null", id).
		First(model).
		Error

	if err != nil {
		return nil, err
	}

	return common.TypeConvertor[Tr](model)
}
