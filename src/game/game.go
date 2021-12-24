package game

import (
	"ludo/src/keyboard"

	"github.com/nsf/termbox-go"
)

type PlayerData struct {
	Name  string
	Type  string
	Color termbox.Attribute
}

type arena struct {
	players []PlayerData
	board   ludoBoard
	curTurn int
}

func handleKeyboard(k keyboard.KeyboardEvent) bool {
	switch k.Key {
	case termbox.KeyEsc:
		return true
	}
	return false
}

func (a *arena) runGameLoop() {
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	go keyboard.ListenToKeyboard(&kChan)

mainloop:
	for {
		select {
		case ev := <-kChan.EvChan:
			if handleKeyboard(ev) {
				kChan.Stop()
				break mainloop
			}
			kChan.Done()
		default:
			a.board.renderBoardLayer()
		}
	}
}

func StartGameOffline(players []PlayerData) {
	ar := arena{board: ludoBoard{}}
	ar.board.setupBoard()
	ar.runGameLoop()
}
