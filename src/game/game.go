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
	players           []PlayerData
	dice              dice
	board             ludoBoard
	curTurn           int
	blinkCh           chan bool
	isBlinkChOpen     bool
	nextWinningPos    int
	participantsCount int
	bots              map[int][4]int
}

func (d *dice) roll() {
	d.value = rand.Intn(6) + 1
}

func (a *arena) changePlayerTurn(idx ...int) bool {
	if len(idx) == 1 {
		a.curTurn = idx[0]
	} else {
		a.curTurn++
		if a.curTurn >= len(a.players) {
			a.curTurn = 0
		}
		for !a.curPlayer().isParticipant() {
			a.curTurn++
			if a.curTurn >= len(a.players) {
				a.curTurn = 0
			}
		}
	}

	if a.curPlayer().isAllPawnsAtDest() {
		a.changePlayerTurn()
	}

	a.board.setCurPawn(0)
	if a.curPawn().isAtDest() || !a.curPawn().hasNPathsAhead(a.dice.value) {
		return a.setNextCurPawnAndValidate(1)
	}

	return true
}

func (a *arena) setNextCurPawnAndValidate(mag int) bool {
	temp := a.board.curPawnIdx
	for a.board.setNextCurPawn(a.curTurn, mag); a.curPawn().isAtDest(); {
		a.board.setNextCurPawn(a.curTurn, mag)
		if a.board.curPawnIdx == temp {
			break
		}
	}

	temp = a.board.curPawnIdx
	for !a.curPawn().hasNPathsAhead(a.dice.value) {
		a.board.setNextCurPawn(a.curTurn, mag)
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

func (a *arena) changePlayerTurnAndValidate() {
	ok := a.changePlayerTurn()
	a.render()
	for !ok {
		a.render()
		time.Sleep(time.Millisecond * 1500)
		a.dice.roll()
		a.render()
		ok = a.changePlayerTurn()
	}
}

func (a *arena) handleKeyboard(k keyboard.KeyboardEvent) bool {
	if a.isGameOver() {
		return k.Key == termbox.KeyEsc
	}

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
		a.dice.roll()
		a.render()
		if !hasDestroyed && !hasReachedDest {
			a.changePlayerTurnAndValidate()
		} else if hasReachedDest {
			if a.curPlayer().isAllPawnsAtDest() {
				a.setCurPlayerWin()
				if a.isGameOver() {
					a.changePlayerTurn()
					a.setCurPlayerWin()
					a.renderGameOver()
					return false
				}
			}
			if ok := a.setNextCurPawnAndValidate(1); !ok {
				a.changePlayerTurnAndValidate()
			}
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
	a.changePlayerTurn(1)
	a.changePlayerTurnAndValidate()
	a.board.setCurPawn(0)
	setRandSeed()
	a.dice.roll()

	a.render()
	a.startBlinkCurPawn()

	for a.curPlayer().isBot() {
		a.playBot()
		if a.isGameOver() {
			break
		}
	}
mainloop:
	for {
		ev := <-kChan.EvChan
		kChan.Pause()
		if stop := a.handleKeyboard(ev); stop {
			kChan.Stop()
			break mainloop
		}
		for a.curPlayer().isBot() {
			a.playBot()
			if a.isGameOver() {
				break
			}
		}
		kChan.Resume()
	}
}

func StartGameOffline(players []PlayerData) {
	participantsCount := 0
	for _, p := range players {
		if p.Type != "-" {
			participantsCount++
		}
	}
	a := arena{
		participantsCount: participantsCount,
		board:             ludoBoard{},
		players:           players,
		blinkCh:           make(chan bool),
		nextWinningPos:    0,
		bots:              make(map[int][4]int),
	}
	a.board.setupBoard(players)
	a.botsInit()
	a.runGameLoop()
}
