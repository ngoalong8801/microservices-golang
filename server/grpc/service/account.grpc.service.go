package service

import (
	"context"
	"github.com/PhuMinh08082001/server-cobra/grpc/repository"
	"proto/grpc/account"
	"proto/grpc/common"
	"strconv"
)

type AccountService struct {
	AccountRepository *repository.AccountRepository
}

func NewAccountService(accountRepository *repository.AccountRepository) *AccountService {
	return &AccountService{
		AccountRepository: accountRepository,
	}
}

func (sv *AccountService) GetUser(ctx context.Context, req *common.FindByIdRequest) (*account.User, error) {

	user := sv.AccountRepository.GetUser(ctx, req)

	userRes := &account.User{
		UserName: user.Name,
		UserId:   strconv.FormatInt(user.ID, 10),
	}
	return userRes, nil
}
