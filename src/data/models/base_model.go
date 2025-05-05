package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id         int          `gorm:"primaryKey"`
	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"` // soft delete

	CreatedBy  int            `gorm:"not null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"` // null if not modified
	DeletedBy  *sql.NullInt64 `gorm:"null"` // null if not deleted

	// All Null Types

	// sql.NullTime
	// sql.NullBool
	// sql.NullString
	// sql.NullInt64
	// sql.NullFloat64
	
	// *sql.NullInt64 means not only value could be null, but also all field could be null
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("user_id")
	// TODO: check userId type
	var userId = -1

	switch v := value.(type) {
	case int:
		userId = v
	case float64:
		userId = int(v) // برای مقادیر JSON یا JWT
	default:
		userId = -1 // یا یه لاگ بزن که "نوع ناشناخته"
	}
	

	// if value != nil {
	// 	userId, _ = value.(int)
	// }

	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("user_id")

	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}

	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.ModifiedBy = userId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("user_id")

	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}

	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.DeletedBy = userId
	return
}
