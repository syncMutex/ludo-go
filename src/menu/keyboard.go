package menu

import (
	"ludo/src/keyboard"

	"github.com/nsf/termbox-go"
)

func listenMenuKeyboard() {
	kb := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}
	go keyboard.ListenToKeyboard(kb)
	for {
		e := <-kb.EvChan
		switch e.Key {
		case 's':
			fallthrough
		case 'S':
			fallthrough
		case termbox.KeyArrowDown:
			handleMenuNav(1)
		case 'w':
			fallthrough
		case 'W':
			fallthrough
		case termbox.KeyArrowUp:
			handleMenuNav(-1)
		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			if quit := handleOptSelect(); quit {
				kb.StopKeyboardListen()
				return
			}
		case termbox.KeyEsc:
			kb.StopKeyboardListen()
			return
		}
		renderMenu()
	}
}
