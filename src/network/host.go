package network

import (
	"ludo/src/common"
	"net"
	"strconv"

	"github.com/nsf/termbox-go"
)

type Player struct {
	common.PlayerData
	isConnected bool
	isReserved  bool
	gh          *GobHandler
}

type ServerArena struct {
	players           []Player
	dice              common.Dice
	curTurn           int
	nextWinningPos    int
	participantsCount int
	bots              map[int][4]int
	availableColors   []termbox.Attribute
}

func (s *ServerArena) joinGame(gh *GobHandler) *Player {
	var playerName string
	gh.Decode(&playerName)

	for i := range s.players {
		p := &s.players[i]
		if !p.isReserved && p.Type == "Player" {
			p.gh = gh
			p.Name = playerName
			p.isReserved = true
			return p
		}
	}
	return nil
}

func (s *ServerArena) sendJoinedPlayers(sendResponse func(int, interface{}) error) {
	joinedPlayers := []common.PlayerData{}
	for _, p := range s.players {
		jp := common.PlayerData{Name: p.Name, Color: p.Color}
		joinedPlayers = append(joinedPlayers, jp)
	}
	sendResponse(common.JOINED_PLAYERS, joinedPlayers)
}

func handleClient(conn net.Conn, server ServerArena) {
	defer conn.Close()
	gh := NewGobHandler(conn)

	gh.SendResponse(common.CONN_RES, common.Res{Ok: true, Msg: "Connected Successfully."})

	for {
		instruc, _ := gh.ReceiveInstruc()
		switch instruc {
		case common.JOIN_GAME:
			playerInfo := server.joinGame(gh)
			if playerInfo == nil {
				gh.SendResponse(common.ERROR, common.Res{Ok: false, Msg: "Game full."})
				return
			}
			gh.SendResponse(common.PLAYER_DATA, common.PlayerData{Type: playerInfo.Type, Color: playerInfo.Color})
		case common.JOINED_PLAYERS:
			server.sendJoinedPlayers(gh.SendResponse)
		}
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
	playersList := []Player{}

	for i, p := range players {
		pl := Player{}
		pl.Color = p.Color
		pl.Type = p.Type
		pl.Name = "Player-" + strconv.Itoa(i)

		playersList = append(playersList, pl)
	}

	server := ServerArena{players: playersList}
	listenRequests(server)
}
