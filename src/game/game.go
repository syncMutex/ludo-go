package game

import (
	"ludo/src/keyboard"
	"time"

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

func (l *ludoBoard) test() {
	paths := l.pathLayer.ll
	temp := paths.head

	pawn := cell{bg: termbox.ColorBlue, ch: ' '}

	for {
		temp.pawn = &pawn
		time.Sleep(time.Millisecond * 50)
		temp.pawn = nil
		temp = temp.next

		if temp == nil {
			temp = paths.head
		}
	}
}

func (a *arena) runGameLoop() {
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	go keyboard.ListenToKeyboard(&kChan)
	go a.board.test()

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
			a.board.render()
		}
	}
}

func StartGameOffline(players []PlayerData) {
	ar := arena{board: ludoBoard{}}
	ar.board.setupBoard()
	ar.runGameLoop()
}
