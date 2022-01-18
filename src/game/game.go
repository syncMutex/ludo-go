package game

import (
	"ludo/src/keyboard"
	"math/rand"
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
	return rand.Intn(6) + 1
}

func (a *arena) changePlayerTurn() {
	a.curTurn++
	if a.curTurn >= len(a.players) {
		a.curTurn = 0
	}

	a.board.setCurPawn(0)

	if a.curPawn().isAtDest() || !a.curPawn().hasNPathsAhead(a.dice.value) {
		a.validateAndSetNextCurPawn(1)
	}
}

func (a *arena) validateAndSetNextCurPawn(mag int) {
	for a.board.setNextCurPawn(a.curTurn, mag, a.dice.value); a.curPawn().isAtDest(); {
		a.board.setNextCurPawn(a.curTurn, mag, a.dice.value)
	}

	for !a.curPawn().hasNPathsAhead(a.dice.value) {
		a.board.setNextCurPawn(a.curTurn, mag, a.dice.value)
	}
}

func (a *arena) handleKeyboard(k keyboard.KeyboardEvent) bool {
	a.stopBlinkCurPawn()
	a.repaintCurPawn()
	switch k.Key {
	case termbox.KeyArrowRight:
		a.validateAndSetNextCurPawn(1)
	case termbox.KeyArrowLeft:
		a.validateAndSetNextCurPawn(-1)
	case termbox.KeyEnter:
		fallthrough
	case termbox.KeySpace:
		if hasDestroyed := a.makeMove(); !hasDestroyed {
			a.dice.value = a.dice.roll()
			a.changePlayerTurn()
		} else {
			a.dice.value = a.dice.roll()
		}
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
		kChan.Pause()
		if stop := a.handleKeyboard(ev); stop {
			kChan.Stop()
			break mainloop
		}
		kChan.Resume()
	}
}

func StartGameOffline(players []PlayerData) {
	ar := arena{board: ludoBoard{}, players: players, blinkCh: make(chan bool)}
	ar.board.setupBoard()
	ar.runGameLoop()
}
