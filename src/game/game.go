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
	players       []PlayerData
	dice          dice
	board         ludoBoard
	pauseKeyboard bool
	curTurn       int
	blinkCh       chan bool
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

func (a *arena) repaintCurPawn() {
	a.board.players[a.curTurn].pawns[a.board.curPawn]["curNode"].cell.bg = a.board.players[a.curTurn].color
}

func (a *arena) handleKeyboard(k keyboard.KeyboardEvent) bool {
	a.stopBlinkCurPawn()
	a.repaintCurPawn()
	switch k.Key {
	case termbox.KeyArrowRight:
		a.board.setNextCurPawn(a.curTurn, 1)
		a.startBlinkCurPawn()
	case termbox.KeyArrowLeft:
		a.board.setNextCurPawn(a.curTurn, -1)
		a.startBlinkCurPawn()
	case termbox.KeySpace:
		a.makeMove()
	case termbox.KeyEsc:
		return true
	}
	return false
}

func (a *arena) runGameLoop() {
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	go keyboard.ListenToKeyboard(&kChan)
	a.board.setCurPawn(0)
	a.startBlinkCurPawn()

mainloop:
	for {
		select {
		case ev := <-kChan.EvChan:
			if a.handleKeyboard(ev) {
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
	ar := arena{board: ludoBoard{}, players: players, blinkCh: make(chan bool)}
	ar.board.setupBoard()
	ar.runGameLoop()
}
