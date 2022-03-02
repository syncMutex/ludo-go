package network

import (
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

type Server struct {
	arena ServerArena
}

func handleClient(conn net.Conn) {
	for {

	}
}

func listenRequests(server ServerArena) {
	for {
		ln, _ := net.Listen("tcp", ":8080")
		conn, _ := ln.Accept()

		if server.participantsCount < 4 {
			conn.Write([]byte("Waiting for other players..."))
		}
		go handleClient(conn)
	}
}

func Host(players []gameUtils.PlayerData) {
	server := ServerArena{}
	listenRequests(server)
}
