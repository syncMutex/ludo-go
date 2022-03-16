package game

import (
	"time"

	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

func (a *Arena) blinkCurPawn(stopBlink <-chan bool) {
	prevColor := termbox.ColorDefault
	curCell := a.curPawn()["curNode"].Cell

	toggleBg := func() {
		if prevColor == termbox.ColorDefault {
			tbu.SetBg(curCell.X, curCell.Y, a.board.Players[a.curTurn].Color)
			prevColor = a.board.Players[a.curTurn].Color
		} else {
			tbu.SetBg(curCell.X, curCell.Y, termbox.ColorDefault)
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
