package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid"`
	Name string
}
