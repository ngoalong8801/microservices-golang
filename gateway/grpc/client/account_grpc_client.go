package client

import (
	"context"
	"gateway-service/common/constants"
	gconfig "gateway-service/grpc/config"
	"google.golang.org/grpc/metadata"
	"proto/grpc/account"
	"proto/grpc/common"
	"time"
)

type AccountClient struct {
	accountClient account.UserServiceClient
}

func NewAccountClient(accountConn *gconfig.AccountConn) *AccountClient {
	return &AccountClient{
		accountClient: account.NewUserServiceClient(accountConn.Cc),
	}
}

func (service *AccountClient) GetUser(req *common.FindByIdRequest, md metadata.MD) (res *account.GetUserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.accountClient.GetUser(ctx, req)
	return
}
