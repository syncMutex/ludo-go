package arena

import (
	"ludo/src/common"
	"ludo/src/keyboard"
	board "ludo/src/ludo-board"
	"time"

	"github.com/nsf/termbox-go"
)

type Arena struct {
	Players           []common.PlayerData
	Dice              common.Dice
	Board             board.LudoBoard
	CurTurn           int
	BlinkCh           chan bool
	IsBlinkChOpen     bool
	NextWinningPos    int
	ParticipantsCount int
	Bots              map[int][4]int
	KChan             *keyboard.KeyboardProps
}

func (a *Arena) SetupBoard() {
	participantsCount := 0
	for _, p := range a.Players {
		if p.Type != "-" {
			participantsCount++
		}
	}
	a.ParticipantsCount = participantsCount
	a.Board.SetupBoard(a.Players)
}

func (a *Arena) SetCurPlayerAndPawn(color termbox.Attribute, pidx int) {
	for i, p := range a.Players {
		if p.Color == color {
			a.CurTurn = i
			a.Board.CurPawnIdx = pidx
			return
		}
	}
}

func (a *Arena) ChangePlayerTurn(idx ...int) bool {
	if len(idx) == 1 {
		a.CurTurn = idx[0]
	} else {
		a.CurTurn++
		if a.CurTurn >= len(a.Players) {
			a.CurTurn = 0
		}
		for !a.CurPlayer().IsParticipant() {
			a.CurTurn++
			if a.CurTurn >= len(a.Players) {
				a.CurTurn = 0
			}
		}
	}

	if a.CurPlayer().IsAllPawnsAtDest() {
		a.ChangePlayerTurn()
	}

	a.Board.SetCurPawn(0)
	if a.CurPawn().IsAtDest() || !a.CurPawn().HasNPathsAhead(a.Dice.Value) {
		return a.SetNextCurPawnAndValidate(1)
	}

	return true
}

func (a *Arena) SetNextCurPawnAndValidate(mag int) bool {
	temp := a.Board.CurPawnIdx
	for a.Board.SetNextCurPawn(a.CurTurn, mag); a.CurPawn().IsAtDest(); {
		a.Board.SetNextCurPawn(a.CurTurn, mag)
		if a.Board.CurPawnIdx == temp {
			break
		}
	}

	temp = a.Board.CurPawnIdx
	for !a.CurPawn().HasNPathsAhead(a.Dice.Value) {
		a.Board.SetNextCurPawn(a.CurTurn, mag)
		if a.Board.CurPawnIdx == temp {
			return false
		}
	}

	return true
}

func (a *Arena) SetCurPlayerWin() {
	a.Board.Players[a.CurTurn].WinningPos = a.NextWinningPos + 1
	a.NextWinningPos++
}

func (a *Arena) ChangePlayerTurnAndValidate(doRender bool) {
	ok := a.ChangePlayerTurn()
	if doRender {
		a.Render()
	}
	for !ok {
		if doRender {
			a.Render()
		}
		time.Sleep(time.Millisecond * 1500)
		a.Dice.Roll()
		if doRender {
			a.Render()
		}
		ok = a.ChangePlayerTurn()
	}
}
