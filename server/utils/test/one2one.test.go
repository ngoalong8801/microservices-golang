package utils

import (
	"fmt"
	"github.com/PhuMinh08082001/server-cobra/dal"
	"github.com/PhuMinh08082001/server-cobra/dal/query"
)

func GetOrderDetail() {
	od := query.Use(dal.DB).OrderDetail
	orderdetail, _ := od.Preload(od.Product, od.Ordertab).First()
	fmt.Println(orderdetail)
}
