package gen

import (
	"github.com/PhuMinh08082001/server-cobra/dal/model"
	"go.uber.org/fx"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var Module = fx.Invoke(generateQuery)

func generateQuery(db *gorm.DB) {

	g := gen.NewGenerator(gen.Config{
		OutPath: "dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
	})

	g.UseDB(db)

	g.ApplyBasic(
		model.User{},
		model.Ordertab{},
		model.Category{},
		model.OrderDetail{},
		model.Product{},
	)

	g.Execute()
}
