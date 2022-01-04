package game

import (
	"time"

	"github.com/nsf/termbox-go"
)

func (a *arena) blinkCurPawn(stopBlink <-chan bool) {
blinkloop:
	for {
		select {
		case <-stopBlink:
			break blinkloop
		case <-time.After(time.Millisecond * 300):
			curPawn := a.board.players[a.curTurn].pawns[a.board.curPawn]["curNode"]
			if curPawn.cell.bg == termbox.ColorDefault {
				curPawn.cell.bg = a.board.players[a.curTurn].color
				continue
			}
			curPawn.cell.bg = termbox.ColorDefault
		}
	}
}

func (a *arena) startBlinkCurPawn() {
	a.isBlinkChOpen = true
	go a.blinkCurPawn(a.blinkCh)
}

func (a *arena) stopBlinkCurPawn() {
	if a.isBlinkChOpen {
		a.blinkCh <- true
		a.isBlinkChOpen = false
	}
}
