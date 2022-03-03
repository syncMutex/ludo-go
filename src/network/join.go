package network

import (
	"net"
)

func Join() (net.Conn, error) {
	return net.Dial("tcp", "127.0.0.1:8080")
}
