package game

import (
	"time"

	"github.com/nsf/termbox-go"
)

func (a *arena) makeMove() {
	curPlayer := a.board.players[a.curTurn]
	for i := 0; i < a.dice.value; i++ {
		curPlayer.pawns[a.board.curPawn]["curNode"].cell.bg = termbox.ColorDefault
		curPlayer.pawns[a.board.curPawn].moveNext("common", curPlayer.color)
		a.render()
		termbox.Flush()
		time.Sleep(time.Millisecond * 100)
	}
}

func (p pawn) moveNext(pathName string, bg termbox.Attribute) {
	p["curNode"] = p["curNode"].next[pathName]
	p["curNode"].cell.bg = bg
}
