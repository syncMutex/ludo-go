package menu

import (
	"ludo/src/keyboard"

	"github.com/nsf/termbox-go"
)

type callback func()

func (m *menuPagesType) keyboardLoop() callback {
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
		case 'a':
			fallthrough
		case 'A':
			fallthrough
		case termbox.KeyArrowLeft:
			m.menus[m.curIdx].handleSubOptNav(-1)
		case 'd':
			fallthrough
		case 'D':
			fallthrough
		case termbox.KeyArrowRight:
			m.menus[m.curIdx].handleSubOptNav(1)
		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			if quit, callback := m.handleOptSelect(); quit {
				kb.StopKeyboardListen()
				return callback
			}
		case termbox.KeyEsc:
			kb.StopKeyboardListen()
			return nil
		}
		m.renderMenu()
	}
}
