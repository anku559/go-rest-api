package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	Id uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`

	Name string `gorm:"unique;not null"`
	Slug string `gorm:"unique;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
