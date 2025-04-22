package models

type Country struct {
	BaseModel
	Name   string `gorm:"type:string;size:15;not null"`
	Cities *[]City
}
