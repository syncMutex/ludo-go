package keyboard

import (
	"sync"

	"github.com/nsf/termbox-go"
)

type KeyboardEvent struct {
	Key termbox.Key
	Ch  rune
}

type KeyboardProps struct {
	EvChan    chan KeyboardEvent
	stopSig   chan bool
	waitGroup sync.WaitGroup
}

func (k *KeyboardProps) Stop() {
	k.Done()
	k.stopSig <- true
}

func (k *KeyboardProps) Done() {
	k.waitGroup.Done()
}

func ListenToKeyboard(kbChans *KeyboardProps) {
	termbox.SetInputMode(termbox.InputEsc)

	kbChans.stopSig = make(chan bool)
	kbChans.waitGroup = *new(sync.WaitGroup)

keyboardLoop:
	for {
		select {
		case v := <-kbChans.stopSig:
			if v {
				break keyboardLoop
			}
		default:
			kbChans.waitGroup.Add(1)
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				kbChans.EvChan <- KeyboardEvent{ev.Key, ev.Ch}
			case termbox.EventError:
				panic(ev.Err)
			}
			kbChans.waitGroup.Wait()
		}
	}
}
