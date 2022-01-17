package keyboard

import (
	"github.com/nsf/termbox-go"
)

type KeyboardEvent struct {
	Key termbox.Key
	Ch  rune
}

type KeyboardProps struct {
	EvChan   chan KeyboardEvent
	stopSig  chan bool
	isPaused bool
}

func (k *KeyboardProps) Stop() {
	k.Resume()
	k.stopSig <- true
}

func (k *KeyboardProps) Pause() {
	k.isPaused = true
}

func (k *KeyboardProps) Resume() {
	k.isPaused = false
}

func ListenToKeyboard(kbChans *KeyboardProps) {
	termbox.SetInputMode(termbox.InputEsc)

	kbChans.stopSig = make(chan bool, 1)

keyboardLoop:
	for {
		select {
		case v := <-kbChans.stopSig:
			if v {
				break keyboardLoop
			}
		default:
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				if kbChans.isPaused {
					break
				}
				kbChans.EvChan <- KeyboardEvent{ev.Key, ev.Ch}
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}
}
