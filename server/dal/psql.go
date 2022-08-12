package dal

import (
	"fmt"
	config2 "github.com/PhuMinh08082001/server-cobra/config"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Module = fx.Provide(NewDB)
var DB *gorm.DB

func NewDB(config config2.Configuration) (db *gorm.DB) {
	var err error

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Database.Username, config.Database.Password, config.Database.Name, config.Database.Port, config.Database.Sslmode)
	fmt.Println(dsn)
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}
