package main

import (
	"go/server/dal"
	"go/server/dal/model"
	"gorm.io/gen"
)

func init() {
	dal.DB = dal.ConnectDB().Debug()
}

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})

	g.UseDB(dal.DB)

	g.ApplyBasic(model.User{})

	g.Execute()
}
