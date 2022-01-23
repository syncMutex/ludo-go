package game

import (
	"time"

	"github.com/nsf/termbox-go"
)

func (a *arena) curPawn(field ...string) pawn {
	return a.curPlayer().pawns[a.board.curPawnIdx]
}

func (a *arena) curPlayer() player {
	return a.board.players[a.curTurn]
}

func (a *arena) makeMove() (hasDestroyed bool, hasReachedDest bool) {
	curPlayerColor := a.curPlayer().color
	curPawn := a.curPawn()

	for i := 0; i < a.dice.value; i++ {
		curPawn["curNode"].cell.bg = termbox.ColorDefault

		if toDest := curPawn["curNode"].next["toDest"]; toDest != nil && toDest.cell.fg == curPlayerColor {
			curPawn.moveToNext("toDest", curPlayerColor)
		} else if curPawn["curNode"].next["common"] != nil {
			curPawn.moveToNext("common", curPlayerColor)
		} else {
			break
		}
		a.render()
		time.Sleep(time.Millisecond * 0)
	}

	hasDestroyed = a.checkDestroy()
	hasReachedDest = curPawn.isAtDest()
	a.render()

	return
}

func (p pawn) moveToNext(pathName string, bg termbox.Attribute) {
	p["curNode"] = p["curNode"].next[pathName]
	p["curNode"].cell.bg = bg
}

func (a *arena) checkDestroy() (hasDestroyed bool) {
	curCell := a.curPawn()["curNode"].cell

	for i, p := range a.board.players {
		if i == a.curTurn {
			continue
		}

		for _, _pawn := range p.pawns {
			c := _pawn["curNode"].cell
			if c.x == curCell.x && c.y == curCell.y {
				hasDestroyed = true
				_pawn["curNode"] = _pawn["homeNode"]
			}
		}
	}

	return
}

func (p player) isAllPawnsAtDest() bool {
	for _, p := range p.pawns {
		if !p.isAtDest() {
			return false
		}
	}
	return true
}

func (a *arena) isGameOver() bool {
	return a.nextWinningPos >= 3
}
