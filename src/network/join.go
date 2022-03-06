package network

import (
	"ludo/src/common"
	"net"
)

func Join() (*common.GobHandler, error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	return common.NewGobHandler(conn), err
}
