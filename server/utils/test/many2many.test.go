package utils

import (
	"fmt"
	"github.com/PhuMinh08082001/server-cobra/dal"
	"github.com/PhuMinh08082001/server-cobra/dal/query"
)

func GetProductTest() {
	p := query.Use(dal.DB).Product

	product, _ := p.Preload(p.Categories).First()

	fmt.Println(product)

}

func GetCategoryTest() {
	c := query.Use(dal.DB).Category

	category, _ := c.Preload(c.Products).First()

	fmt.Println(category)

}
