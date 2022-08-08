package config

import "net"

func NewNetListener(config *Configuration) (net.Listener, error) {
	return net.Listen(config.Grpc.Network, config.Grpc.Host+string(rune(config.Grpc.Port)))
}
