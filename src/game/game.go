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
	players        []PlayerData
	dice           dice
	board          ludoBoard
	curTurn        int
	blinkCh        chan bool
	isBlinkChOpen  bool
	nextWinningPos int
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
		if ok := a.setNextCurPawnAndValidate(1); !ok {
			a.dice.value = a.dice.roll()
			a.render()
			if !a.curPlayer().isAllPawnsAtDest() {
				time.Sleep(time.Second)
			}
			a.changePlayerTurn()
		}
	}
}

func (a *arena) setNextCurPawnAndValidate(mag int) bool {
	temp := a.board.curPawnIdx
	for a.board.setNextCurPawn(a.curTurn, mag, a.dice.value); a.curPawn().isAtDest(); {
		a.board.setNextCurPawn(a.curTurn, mag, a.dice.value)
		if a.board.curPawnIdx == temp {
			break
		}
	}

	temp = a.board.curPawnIdx
	for !a.curPawn().hasNPathsAhead(a.dice.value) {
		a.board.setNextCurPawn(a.curTurn, mag, a.dice.value)
		if a.board.curPawnIdx == temp {
			return false
		}
	}

	return true
}

func (a *arena) setCurPlayerWin() {
	a.board.players[a.curTurn].winningPos = a.nextWinningPos + 1
	a.nextWinningPos++
}

func (a *arena) handleKeyboard(k keyboard.KeyboardEvent) bool {
	a.stopBlinkCurPawn()
	a.repaintCurPawn()
	switch k.Key {
	case termbox.KeyArrowRight:
		a.setNextCurPawnAndValidate(1)
	case termbox.KeyArrowLeft:
		a.setNextCurPawnAndValidate(-1)
	case termbox.KeyEnter:
		fallthrough
	case termbox.KeySpace:
		hasDestroyed, hasReachedDest := a.makeMove()

		if !hasDestroyed && !hasReachedDest {
			a.changePlayerTurn()
		} else if hasReachedDest {
			if a.board.players[a.curTurn].isAllPawnsAtDest() {
				a.setCurPlayerWin()
				a.changePlayerTurn()
				if a.isGameOver() {
					a.setCurPlayerWin()
					return true
				}
			} else if ok := a.setNextCurPawnAndValidate(1); !ok {
				a.changePlayerTurn()
			}
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
