package db

import (
	"fmt"

	"github.com/impleopus/go-rest-api/internal/config/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var err error

	urlString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		env.PG_HOST,
		env.PG_USER,
		env.PG_PASSWORD,
		env.PG_DB_NAME,
		env.PG_PORT,
		env.PG_SSL_MODE,
	)
	println(urlString)

	DB, err = gorm.Open(postgres.Open(urlString), &gorm.Config{})

	if err != nil {
		panic("Error Connecting Database")
	}
}
