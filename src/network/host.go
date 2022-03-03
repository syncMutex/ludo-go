package network

import (
	"encoding/gob"
	gameUtils "ludo/src/game-utils"
	"net"
)

type Player struct {
	gameUtils.PlayerData
	conn net.Conn
}

type ServerArena struct {
	players           []Player
	dice              gameUtils.Dice
	curTurn           int
	nextWinningPos    int
	participantsCount int
	bots              map[int][4]int
}

func handleClient(conn net.Conn, server ServerArena) {
	enc := gob.NewEncoder(conn)
	//	dec := gob.NewDecoder(conn)

	enc.Encode(gameUtils.ConnRes{Ok: true, Msg: "Connected successfully."})
}

func listenRequests(server ServerArena) {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		go handleClient(conn, server)
	}
}

func Host(players []gameUtils.PlayerData) {
	server := ServerArena{}
	listenRequests(server)
}
