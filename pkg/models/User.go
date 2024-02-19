package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`

	FirstName  string
	MiddleName string
	LastName   string

	Email       string `gorm:"unique"`
	Username    string
	CountryCode string
	Phone       string
	Password    string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
