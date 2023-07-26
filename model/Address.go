package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

)

type Address struct {
	gorm.Model
	ID      uuid.UUID `gorm:"column:id" gorm:"type:uuid" json:"id"`
	UserID  uuid.UUID `gorm:"column:user_id" gorm:"type:uuid" json:"user_id"`
	Address string    `gorm:"column:address" json:"address"`
	User    User      `gorm:"foreignKey:UserID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
