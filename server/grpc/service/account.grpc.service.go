package service

import (
	"context"
	"go/server/grpc/repository"
	"proto/grpc/account"
	"proto/grpc/common"
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
		UserId:   string(user.ID),
	}
	return userRes, nil
}
