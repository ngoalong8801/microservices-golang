package server

import (
	"context"
	"go/server/grpc/service"
	"proto/grpc/account"
	"proto/grpc/common"
)

type AccountServer struct {
	account.UnimplementedUserServiceServer
	AccountService *service.AccountService
}

func NewAccountServer(accountService *service.AccountService) *AccountServer {
	return &AccountServer{
		AccountService: accountService,
	}
}

func (acc *AccountServer) GetUser(ctx context.Context, req *common.FindByIdRequest) (*account.GetUserResponse, error) {
	user, _ := acc.AccountService.GetUser(ctx, req)

	resp := &account.GetUserResponse_Data{User: user}

	a := &account.GetUserResponse{
		Success:  true,
		Response: &account.GetUserResponse_Data_{Data: resp},
	}

	return a, nil
}
