package network

import (
	"ludo/src/common"
	"ludo/src/game/arena"
	"ludo/src/network/schema"
	"time"
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
	playerName := DecodeData[string](gh)

	for i := range s.players {
		p := &(s.players[i])
		if !p.isReserved && p.Type == "Player" {
			p.gh = gh
			p.Name = playerName
			p.isReserved = true
			a := &(s.arena)
			a.Players[i].Type = p.Type
			a.Players[i].Color = p.Color
			a.Players[i].Name = p.Name
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
	s.broadcastResponse(schema.JOINED_PLAYERS, joinedPlayers)
}

func (s *Server) boardState() (brdSt schema.BoardState) {
	brdSt.CurTurn = s.arena.CurPlayer().Color
	brdSt.DiceValue = s.arena.Dice.Value
	return
}

func (s *Server) setupBoard() {
	a := &(s.arena)
	a.SetupBoard()
	a.CurTurn = 1
	a.ChangePlayerTurnAndValidate(DONT_RENDER, nil)
	a.Board.SetCurPawn(0)
}

func (s *Server) broadcastBoardState() {
	s.broadcastResponse(schema.BOARD_STATE, s.boardState())
}

func (s *Server) broadcastLoop() {
	s.broadcastResponse(schema.LOOPED, s.boardState())
}

func (s *Server) onMove(pawnIdx int, playerInfo common.PlayerData) {
	a := &(s.arena)
	a.SetCurPlayerAndPawn(playerInfo.Color, pawnIdx)

	hasDestroyed, hasReachedDest := a.MakeMove(time.Millisecond*0, DONT_RENDER)
	s.broadcastResponseExcept(schema.MOVE_BY, schema.MoveBy{Color: playerInfo.Color, PawnIdx: pawnIdx}, playerInfo.Color)
	a.Dice.Roll()

	if !hasDestroyed && !hasReachedDest {
		a.ChangePlayerTurnAndValidate(DONT_RENDER, s.broadcastLoop)
	} else if hasReachedDest {
		if a.CurPlayer().IsAllPawnsAtDest() {
			a.SetCurPlayerWin()
			if a.IsGameOver() {
				a.ChangePlayerTurn()
				a.SetCurPlayerWin()
				s.broadcastResponse(schema.GAME_OVER, s.arena.LeaderBoard())
			}
		}
		if ok := a.SetNextCurPawnAndValidate(1); !ok {
			a.ChangePlayerTurnAndValidate(DONT_RENDER, nil)
		}
	}
	s.broadcastBoardState()
}
