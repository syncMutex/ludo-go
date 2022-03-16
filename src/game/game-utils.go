package game

import (
	board "ludo/src/ludo-board"
	tbu "ludo/src/termbox-utils"
	"time"

	"github.com/nsf/termbox-go"
)

func (a *Arena) curPawn() board.Pawn {
	return a.curPlayer().Pawns[a.board.CurPawnIdx]
}

func (a *Arena) curPlayer() board.Player {
	return a.board.Players[a.curTurn]
}

func (a *Arena) makeMove() (hasDestroyed bool, hasReachedDest bool) {
	curPlayerColor := a.curPlayer().Color
	curPawn := a.curPawn()

	for i := 0; i < a.Dice.Value; i++ {
		curPawn["curNode"].Cell.Bg = termbox.ColorDefault

		if toDest := curPawn["curNode"].Next["toDest"]; toDest != nil && toDest.Cell.Fg == curPlayerColor {
			curPawn.MoveToNext("toDest", curPlayerColor)
		} else if curPawn["curNode"].Next["common"] != nil {
			curPawn.MoveToNext("common", curPlayerColor)
		} else {
			break
		}
		a.render()
		time.Sleep(time.Millisecond * 0)
	}

	hasDestroyed = a.checkDestroy()
	hasReachedDest = curPawn.IsAtDest()
	a.render()

	return
}

func (a *Arena) checkDestroy() (hasDestroyed bool) {
	curCell := a.curPawn()["curNode"].Cell

	for i, p := range a.board.Players {
		if i == a.curTurn || !p.IsParticipant() {
			continue
		}

		for j, _pawn := range p.Pawns {
			c := _pawn["curNode"].Cell
			if c.X == curCell.X && c.Y == curCell.Y {
				hasDestroyed = true
				if a.curPlayer().IsBot() {
					a.resetBotPawn(i, j)
				}
				_pawn["curNode"] = _pawn["homeNode"]
			}
		}
	}

	return
}

func (a *Arena) isGameOver() bool {
	return a.nextWinningPos >= a.participantsCount-1
}

func (a *Arena) repaintCurPawn() {
	curCell := a.curPawn()["curNode"].Cell
	tbu.SetBg(curCell.X, curCell.Y, a.board.Players[a.curTurn].Color)
}
