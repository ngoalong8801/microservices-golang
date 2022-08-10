package migration

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/fx"
	"go/server/config"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "test1"
)

var Module = fx.Invoke(Start)

func Start(config config.Configuration) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Name)

	log.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		"postgres", driver)

	err = m.Up()

	if err != nil {
		log.Println("Database has " + err.Error())
	}
}
