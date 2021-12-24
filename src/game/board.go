package game

import (
	"github.com/nsf/termbox-go"
)

type pos struct{ x, y int }

type ludoBoard struct {
	pawnsLocations []pawnData
	boardData      colorMap
}

type pawnData struct {
	color    termbox.Attribute
	pawnsPos [4]pos
}

type elementGroup []interface{}

func (board *ludoBoard) renderBoard() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, p := range board.boardData {
		termbox.SetCell(p.x, p.y, p.ch, p.fg, p.bg)
	}
	termbox.Flush()
}

func createBoardData() colorMap {
	boardPos := pos{5, 2}

	cm := colorMap{}
	cm.mergeColorMap(
		createBoardSkeleton(boardPos),
	)
	return cm
}

func (board *ludoBoard) setupBoard() {
	board.boardData = createBoardData()

	board.renderBoard()
}
