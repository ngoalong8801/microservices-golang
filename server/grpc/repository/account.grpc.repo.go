package repository

import (
	"context"
	"fmt"
	"github.com/PhuMinh08082001/server-cobra/dal/model"
	"github.com/PhuMinh08082001/server-cobra/dal/query"
	"gorm.io/gorm"
	"proto/grpc/common"
	"strconv"
)

type AccountRepository struct {
	Db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		Db: db,
	}
}

func (repo *AccountRepository) GetUser(ctx context.Context, req *common.FindByIdRequest) *model.User {
	userId, _ := strconv.ParseInt(req.GetId(), 10, 64)
	u := query.Use(repo.Db).User

	user, err := u.WithContext(ctx).Where(u.ID.Eq(userId)).First()

	if err != nil {
		fmt.Println("Fail when fetch user " + req.GetId())
	}

	return user
}
