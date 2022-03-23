package network

import (
	"net"
)

func Join(url string) (*GobHandler, error) {
	conn, err := net.Dial("tcp", url)
	return NewGobHandler(conn), err
}
