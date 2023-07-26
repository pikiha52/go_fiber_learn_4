package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid"`
	UserID uuid.UUID `gorm:"type:uuid"`
	RoleID uuid.UUID `gorm:"type:uuid"`
	User   User      `gorm:"foreignKey:UserID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role   Role      `gorm:"foreignKey:RoleID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
