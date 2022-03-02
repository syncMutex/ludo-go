package game

import (
	"time"

	"github.com/nsf/termbox-go"
)

func (a *Arena) blinkCurPawn(stopBlink <-chan bool) {
	prevColor := termbox.ColorDefault
	curCell := a.curPawn()["curNode"].cell

	toggleBg := func() {
		if prevColor == termbox.ColorDefault {
			setBg(curCell.x, curCell.y, a.board.players[a.curTurn].color)
			prevColor = a.board.players[a.curTurn].color
		} else {
			setBg(curCell.x, curCell.y, termbox.ColorDefault)
			prevColor = termbox.ColorDefault
		}
		termbox.Flush()
	}

	toggleBg()

blinkloop:
	for {
		select {
		case <-stopBlink:
			break blinkloop
		case <-time.After(time.Millisecond * 300):
			toggleBg()
		}
	}
}

func (a *Arena) startBlinkCurPawn() {
	a.isBlinkChOpen = true
	go a.blinkCurPawn(a.blinkCh)
}

func (a *Arena) stopBlinkCurPawn() {
	if a.isBlinkChOpen {
		a.blinkCh <- true
		a.isBlinkChOpen = false
	}
}
