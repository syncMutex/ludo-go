package menu

import (
	"ludo/src/keyboard"

	"github.com/nsf/termbox-go"
)

func (m *menuPagesType) keyboardLoop() {
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
			m.menus[m.curIdx].handleOptNav(1)
		case 'w':
			fallthrough
		case 'W':
			fallthrough
		case termbox.KeyArrowUp:
			m.menus[m.curIdx].handleOptNav(-1)
		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			if quit := m.handleOptSelect(); quit {
				kb.StopKeyboardListen()
				return
			}
		case termbox.KeyEsc:
			kb.StopKeyboardListen()
			return
		}
		m.renderMenu()
	}
}
