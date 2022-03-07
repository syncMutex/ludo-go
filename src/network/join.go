package network

import (
	"net"
)

func Join() (*GobHandler, error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	return NewGobHandler(conn), err
}
