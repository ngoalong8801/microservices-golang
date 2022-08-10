package migration

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go/server/config"
	"log"
)

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
