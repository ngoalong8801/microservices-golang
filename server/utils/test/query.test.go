package utils

import (
	"fmt"
	"github.com/PhuMinh08082001/server-cobra/dal"
	"github.com/PhuMinh08082001/server-cobra/dal/model"
	"github.com/PhuMinh08082001/server-cobra/dal/query"
)

func GetUserHaveGTOneOrder() []*model.User {
	u := query.Use(dal.DB).User
	o := query.Use(dal.DB).Ordertab

	ua := u
	oa := o

	users, err := ua.Preload(ua.Ordertabs).Exists(
		oa.Select(oa.OrderID.Count()).
			Group(oa.UserID).Where(ua.ID.EqCol(oa.UserID)).
			Having(oa.OrderID.Count().Gt(0)),
	).Find()

	if err != nil {
		panic(fmt.Errorf("query fail #{err}"))
	}

	return users
}
