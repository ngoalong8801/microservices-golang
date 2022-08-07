package gconfig

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTransportOption() grpc.DialOption {
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}
