package game

import (
	"time"

	"github.com/nsf/termbox-go"
)

func (a *arena) makeMove() {
	curPlayer := a.board.players[a.curTurn]
	curPawn := curPlayer.pawns[a.board.curPawn]

	for i := 0; i < a.dice.value; i++ {
		curPawn["curNode"].cell.bg = termbox.ColorDefault

		if toDest := curPawn["curNode"].next["toDest"]; toDest != nil && toDest.cell.fg == curPlayer.color {
			curPawn.moveToNext("toDest", curPlayer.color)
		} else if curPawn["curNode"].next["common"] != nil {
			curPawn.moveToNext("common", curPlayer.color)
		} else {
			break
		}
		a.render()
		time.Sleep(time.Millisecond * 100)
	}

	a.checkDestroy()
	a.render()
}

func (a *arena) checkDestroy() bool {
	curPlayer := a.board.players[a.curTurn]
	curCell := curPlayer.pawns[a.board.curPawn]["curNode"].cell

	hasDestroyed := false

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

	return hasDestroyed
}

func (p pawn) moveToNext(pathName string, bg termbox.Attribute) {
	p["curNode"] = p["curNode"].next[pathName]
	p["curNode"].cell.bg = bg
}
