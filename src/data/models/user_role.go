package models

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnDelete:NO ACTION;OnUpdate:NO ACTION;"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnDelete:NO ACTION;OnUpdate:NO ACTION;"`
	UserId int
	RoleId int
}
