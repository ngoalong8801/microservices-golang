package config

import "net"

const addr = "localhost:9999"

func NewNetListener(config Configuration) (net.Listener, error) {
	return net.Listen("tcp", addr)
}
