package game

import (
	"ludo/src/keyboard"
	"math/rand"

	"github.com/nsf/termbox-go"
)

type PlayerData struct {
	Name  string
	Type  string
	Color termbox.Attribute
}

type dice struct {
	value int
}

func (d *dice) roll() {
	d.value = rand.Intn(5) + 1
}

func (d *dice) reset() {
	d.value = 0
}

type arena struct {
	players []PlayerData
	dice    dice
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
			a.board.render()
		}
	}
}

func StartGameOffline(players []PlayerData) {
	ar := arena{board: ludoBoard{}, players: players}
	ar.board.setupBoard()
	ar.runGameLoop()
}
