package gconfig

import (
	"gateway-service/config"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type AccountParams struct {
	fx.In
	Option        grpc.DialOption
	Configuration config.Configuration
}
type AccountConn struct {
	Cc *grpc.ClientConn
}

func NewAccountCC(params AccountParams) *AccountConn {
	conn, _ := grpc.Dial(params.Configuration.AccountClientUrl, params.Option)
	return &AccountConn{Cc: conn}
}
