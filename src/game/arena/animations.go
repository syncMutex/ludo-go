package arena

import (
	"time"

	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

func (a *Arena) BlinkCurPawn(stopBlink <-chan bool) {
	prevColor := termbox.ColorDefault
	curCell := a.CurPawn()["curNode"].Cell

	toggleBg := func() {
		if prevColor == termbox.ColorDefault {
			tbu.SetBg(curCell.X, curCell.Y, a.Board.Players[a.CurTurn].Color)
			prevColor = a.Board.Players[a.CurTurn].Color
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

func (a *Arena) StartBlinkCurPawn() {
	a.IsBlinkChOpen = true
	go a.BlinkCurPawn(a.BlinkCh)
}

func (a *Arena) StopBlinkCurPawn() {
	if a.IsBlinkChOpen {
		a.BlinkCh <- true
		a.IsBlinkChOpen = false
	}
}
