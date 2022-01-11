package game

import (
	"time"

	"github.com/nsf/termbox-go"
)

func (a *arena) makeMove() {
	curPlayer := a.board.players[a.curTurn]

	for i := 0; i < a.dice.value; i++ {
		curPawn := curPlayer.pawns[a.board.curPawn]
		curPawn["curNode"].cell.bg = termbox.ColorDefault
		curPawn.moveToNext("common", curPlayer.color)
		a.render()
		time.Sleep(time.Millisecond * 100)
	}
}

func (p pawn) moveToNext(pathName string, bg termbox.Attribute) {
	p["curNode"] = p["curNode"].next[pathName]
	p["curNode"].cell.bg = bg
}
