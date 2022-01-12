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

func (b *ludoBoard) setCurPawn(idx int) {
	b.curPawn = idx
}

func (b *ludoBoard) setNextCurPawn(curTurn, mag int) {
	b.setCurPawn(b.curPawn + mag)

	if b.curPawn < 0 {
		b.curPawn = len(b.players[curTurn].pawns) - 1
	} else if b.curPawn >= len(b.players[curTurn].pawns) {
		b.curPawn = 0
	}
}

func (a *arena) changePlayerTurn() {
	a.curTurn++
	if a.curTurn >= len(a.players) {
		a.curTurn = 0
	}
}

func (a *arena) repaintCurPawn() {
	curCell := a.board.players[a.curTurn].pawns[a.board.curPawn]["curNode"].cell
	setBg(curCell.x, curCell.y, a.board.players[a.curTurn].color)
}

func (a *arena) handleKeyboard(k keyboard.KeyboardEvent) bool {
	a.stopBlinkCurPawn()
	a.repaintCurPawn()
	switch k.Key {
	case termbox.KeyArrowRight:
		a.board.setNextCurPawn(a.curTurn, 1)
	case termbox.KeyArrowLeft:
		a.board.setNextCurPawn(a.curTurn, -1)
	case termbox.KeyEnter:
		fallthrough
	case termbox.KeySpace:
		a.makeMove()
		a.changePlayerTurn()
	case termbox.KeyEsc:
		return true
	default:

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
		if a.handleKeyboard(ev) {
			kChan.Stop()
			break mainloop
		}
		kChan.Done()
	}
}

func StartGameOffline(players []PlayerData) {
	ar := arena{board: ludoBoard{}, players: players, blinkCh: make(chan bool)}
	ar.board.setupBoard()
	ar.runGameLoop()
}
