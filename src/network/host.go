package network

import (
	"ludo/src/common"
	"ludo/src/game/arena"
	board "ludo/src/ludo-board"
	"net"
	"strconv"
)

func handleClient(conn net.Conn, server Server) {
	defer conn.Close()
	gh := NewGobHandler(conn)

	gh.SendResponse(common.CONN_RES, common.Res{Ok: true, Msg: "Connected Successfully."})
	playerInfo := server.getNameAndJoinGame(gh)
	if playerInfo == nil {
		gh.SendResponse(common.KNOWN_ERR, common.Res{Ok: false, Msg: "Game full."})
		return
	}
	gh.SendResponse(common.PLAYER_COLOR, playerInfo.Color)
	server.updateJoinedPlayers()

	if server.isAllReserved() {
		server.broadcastResponse(common.START_GAME, server.arena)
		server.setupBoard()
		brdSt := server.boardState()
		server.broadcastResponse(common.BOARD_STATE, brdSt)
	}

	for {
		instruc, _ := gh.ReceiveInstruc()
		switch instruc {
		}
	}
}

func listenRequests(server Server) {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		go handleClient(conn, server)
	}
}

func Host(players []common.PlayerData) {
	playersList := []Player{}

	for i, p := range players {
		pl := Player{}
		pl.Color = p.Color
		pl.Type = p.Type
		pl.Name = "Player-" + strconv.Itoa(i)

		playersList = append(playersList, pl)
	}

	common.SetRandSeed()
	gameDice := common.Dice{}
	gameDice.Roll()

	server := Server{
		players: playersList,
		arena: arena.Arena{
			Board:          board.LudoBoard{},
			Players:        players,
			NextWinningPos: 0,
			Dice:           gameDice,
			Bots:           make(map[int][4]int),
			// will be init on client
			KChan:   nil,
			BlinkCh: nil,
		},
	}
	listenRequests(server)
}
