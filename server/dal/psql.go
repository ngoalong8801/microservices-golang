package dal

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (db *gorm.DB) {
	var err error

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=123456 dbname=grpc port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	return db
}
