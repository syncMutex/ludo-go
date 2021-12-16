package keyboard

import (
	"github.com/nsf/termbox-go"
)

type KeyboardEvent struct {
	Key termbox.Key
	Ch  rune
}

type KeyboardProps struct {
	EvChan chan KeyboardEvent
	stop   bool
}

func (kb *KeyboardProps) StopKeyboardListen() {
	kb.stop = true
}

func ListenToKeyboard(kbChans KeyboardProps) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		if kbChans.stop {
			break
		} else {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				kbChans.EvChan <- KeyboardEvent{ev.Key, ev.Ch}
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}
}
