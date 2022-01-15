package game

import (
	"ludo/src/keyboard"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type PlayerData struct {
	Type  string
	Color termbox.Attribute
}

type dice struct {
	value int
}

type arena struct {
	players       []PlayerData
	dice          dice
	board         ludoBoard
	curTurn       int
	blinkCh       chan bool
	isBlinkChOpen bool
}

func (d *dice) roll() int {
	return 6
}

func (a *arena) changePlayerTurn() {
	a.curTurn++
	if a.curTurn >= len(a.players) {
		a.curTurn = 0
	}

	a.board.setCurPawn(0)

	for a.curPawn().isAtDest() {
		a.board.setNextCurPawn(a.curTurn, 1)
	}
}

func (a *arena) handleKeyboard(k keyboard.KeyboardEvent) bool {
	a.stopBlinkCurPawn()
	a.repaintCurPawn()
	switch k.Key {
	case termbox.KeyArrowRight:
		for a.board.setNextCurPawn(a.curTurn, 1); a.curPawn().isAtDest(); {
			a.board.setNextCurPawn(a.curTurn, 1)
		}
	case termbox.KeyArrowLeft:
		for a.board.setNextCurPawn(a.curTurn, -1); a.curPawn().isAtDest(); {
			a.board.setNextCurPawn(a.curTurn, -1)
		}
	case termbox.KeyEnter:
		fallthrough
	case termbox.KeySpace:
		if hasDestroyed := a.makeMove(); !hasDestroyed {
			a.changePlayerTurn()
		}
		a.dice.value = a.dice.roll()
	case termbox.KeyEsc:
		return true
	}
	a.render()
	a.startBlinkCurPawn()
	return false
}

func setRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func (a *arena) runGameLoop() {
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	go keyboard.ListenToKeyboard(&kChan)
	a.board.setCurPawn(0)
	setRandSeed()
	a.dice.value = a.dice.roll()

	a.render()
	a.startBlinkCurPawn()

mainloop:
	for {
		ev := <-kChan.EvChan
		if stop := a.handleKeyboard(ev); stop {
			kChan.Stop()
			break mainloop
		}
		kChan.Done()
		os.Stdin.Sync()
	}
}

func StartGameOffline(players []PlayerData) {
	ar := arena{board: ludoBoard{}, players: players, blinkCh: make(chan bool)}
	ar.board.setupBoard()
	ar.runGameLoop()
}
