package dal

import (
	"fmt"
	"go.uber.org/fx"
	config2 "go/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Provide(NewDB)
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

func NewDB(config config2.Configuration) (db *gorm.DB) {
	var err error

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Database.Username, config.Database.Password, config.Database.Name, config.Database.Port, config.Database.Sslmode)

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}
