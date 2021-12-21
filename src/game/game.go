package game

import (
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

func StartGameOffline(players []PlayerData) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
	ar := arena{board: ludoBoard{}}
	ar.board.setupBoard()
	time.Sleep(time.Second * 5)
}
