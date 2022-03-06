package network

import (
	"ludo/src/common"
	"net"
	"time"
)

type Player struct {
	common.PlayerData
	conn net.Conn
}

type ServerArena struct {
	players           []Player
	dice              common.Dice
	curTurn           int
	nextWinningPos    int
	participantsCount int
	bots              map[int][4]int
}

func handleClient(conn net.Conn, server ServerArena) {
	defer conn.Close()
	gh := common.NewGobHandler(conn)

	for {
		time.Sleep(time.Second)
		toSend := common.ConnRes{Ok: true, Msg: "Connected successfully."}
		gh.SendResponse(common.CONN_RES, toSend)
		gh.SendResponse(common.TEST_RES, common.TestRes{"HE HE HE HAW"})
	}
}

func listenRequests(server ServerArena) {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		go handleClient(conn, server)
	}
}

func Host(players []common.PlayerData) {
	server := ServerArena{}
	listenRequests(server)
}
