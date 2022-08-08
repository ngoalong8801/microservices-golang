package server

import (
	"context"
	"proto/grpc/account"
	"proto/grpc/common"
)

type AccountServer struct {
	account.UnimplementedUserServiceServer
}

func (acc *AccountServer) GetUser(ctx context.Context, req *common.FindByIdRequest) (*account.GetUserResponse, error) {
	return &account.GetUserResponse{
		Success: true,
	}, nil
}
