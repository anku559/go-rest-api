package db

import (
	"github.com/impleopus/go-rest-api/pkg/models"
)

func SyncDB() {

	// Migrate Schemas
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Category{})
}
