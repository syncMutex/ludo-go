package network

import (
	"ludo/src/common"
	"ludo/src/game/arena"
)

type Player struct {
	common.PlayerData
	isConnected bool
	isReserved  bool
	gh          *GobHandler
}

type Server struct {
	players []Player
	arena   arena.Arena
}

func (s Server) isAllReserved() bool {
	for _, p := range s.players {
		if !p.isReserved && p.Type == "Player" {
			return false
		}
	}
	return true
}

func (s *Server) getNameAndJoinGame(gh *GobHandler) *Player {
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

func (s *Server) updateJoinedPlayers() {
	joinedPlayers := []common.PlayerData{}
	for _, p := range s.players {
		if p.isReserved {
			jp := common.PlayerData{Name: p.Name, Color: p.Color}
			joinedPlayers = append(joinedPlayers, jp)
		}
	}
	s.broadcastResponse(common.JOINED_PLAYERS, joinedPlayers)
}

func (s *Server) boardState() (brdSt common.BoardState) {
	brdSt.CurTurn = s.arena.CurPlayer().Color
	brdSt.DiceValue = s.arena.Dice.Value
	return
}

func (s *Server) setupBoard() {
	s.arena.SetupBoard()
	s.arena.CurTurn = 1
	s.arena.ChangePlayerTurnAndValidate()
	s.arena.Board.SetCurPawn(0)
}
