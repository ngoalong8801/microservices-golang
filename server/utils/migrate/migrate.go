package migrate

import (
	"fmt"
	"github.com/PhuMinh08082001/server-cobra/config"
	"go.uber.org/fx"
	"log"
	"os/exec"
)

var Command *Cmd

type Cmd struct {
	Name string
	Step string
}

var Module = fx.Invoke(Start)

func Start(config config.Configuration) {

	cmd := exec.Command("migration", "-source", "file://db/migration", "-database", "postgres://postgres:123456@localhost:5432/grpc", "-verbose", "up")

	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Println(string(stdout))

	//output, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Output:\n%s\n", string(output))

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.Name)
	//
	//log.Println(psqlInfo)
	//
	//db, err := sql.Open("postgres", psqlInfo)
	//
	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//
	//	}
	//}(db)
	//
	//driver, err := postgres.WithInstance(db, &postgres.Config{})
	//
	//m, err := migration.NewWithDatabaseInstance(
	//	"file://db/migration",
	//	"postgres", driver)
	//
	//err = m.Up()
	//
	//if err != nil {
	//	log.Println("Database has " + err.Error())
	//}
}
