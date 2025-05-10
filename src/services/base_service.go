package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	"github.com/MohammadGholamrezai/golang-clean-web-api/api/dto"
	"github.com/MohammadGholamrezai/golang-clean-web-api/common"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/db"
	"github.com/MohammadGholamrezai/golang-clean-web-api/data/models"
	"github.com/MohammadGholamrezai/golang-clean-web-api/pkg/service_errors"
	"gorm.io/gorm"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Preloads []preload
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

	// return common.TypeConvertor[Tr](model)
	baseModel, _ := common.TypeConvertor[models.BaseModel](model)
	return s.GetById(ctx, baseModel.Id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, req *Tu, id int) (*Tr, error) {
	updateMap, _ := common.TypeConvertor[map[string]interface{}](req)

	snakeMap := map[string]interface{}{}
	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}

	snakeMap["ModifiedBy"] = &sql.NullInt64{Int64: int64(ctx.Value("user_id").(float64)), Valid: true}
	snakeMap["ModifiedAt"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}

	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.Model(model).
		Where("id = ? AND deleted_by is null", id).
		Updates(snakeMap).
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

// GetByFilter
func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[Tr], error) {
	return Paginate[T, Tr](req, s.Preloads, s.Database)
}

// Paginate
func Paginate[T any, Tr any](pagination *dto.PaginationInputWithFilter, preloads []preload, db *gorm.DB) (*dto.PagedList[Tr], error) {
	model := new(T)
	var items *[]T
	var rItems *[]Tr

	db = Preload(db, preloads)
	query := getQuery[T](&pagination.DynamicFilter)
	sort := getSort[T](&pagination.DynamicFilter)

	var totalRows int64 = 0
	db.
		Model(model).
		Where(query).
		Count(&totalRows)

	err := db.
		Where(query).
		Offset(pagination.GetOffset()).
		Limit(pagination.GetPageSize()).
		Order(sort).
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}

	rItems, err = common.TypeConvertor[[]Tr](items)
	if err != nil {
		return nil, err
	}

	return NewPagedList(rItems, totalRows, pagination.PageNumber, int64(pagination.PageSize)), err
}

// Preload
func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, v := range preloads {
		db = db.Preload(v.string)
	}

	return db
}

// getQuery
func getQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)

	query := make([]string, 0)
	query = append(query, "deleted_by is null")
	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			fld, ok := typeT.FieldByName(name)
			if ok {
				fld.Name = common.ToSnakeCase(fld.Name)
				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "notContains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startsWith":
					query = append(query, fmt.Sprintf("%s ILike '%s%%'", fld.Name, filter.From))
				case "endsWith":
					query = append(query, fmt.Sprintf("%s ILike '%%%s'", fld.Name, filter.From))
				case "equals":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEqual":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lessThan":
					query = append(query, fmt.Sprintf("%s < '%s'", fld.Name, filter.From))
				case "lessThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > '%s'", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= %s", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= %s", fld.Name, filter.To))
					}
				}
			}
		}
	}

	return strings.Join(query, " AND ")
}

// getSort
func getSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)

	if filter.Sort != nil {
		for _, dto := range *filter.Sort {
			fld, ok := typeT.FieldByName(dto.ColId)
			if ok && (dto.Sort == "asc" || dto.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, dto.Sort))
			}
		}
	}

	return strings.Join(sort, ", ")
}

func NewPagedList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PagedList[T] {
	pl := &dto.PagedList[T]{
		PageNumber: pageNumber,
		TotalRows:  count,
		Items:      items,
	}

	pl.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasPreviousPage = pl.PageNumber > 1

	return pl
}
