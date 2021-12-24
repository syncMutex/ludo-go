package game

import (
	"github.com/nsf/termbox-go"
)

type pos struct{ x, y int }

type ludoBoard struct {
	pawnsLocations []pawnData
	boardLayer     cellMap
}

type pawnData struct {
	color    termbox.Attribute
	pawnsPos [4]pos
}

type elementGroup []interface{}

func (board *ludoBoard) renderBoardLayer() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, p := range board.boardLayer {
		termbox.SetCell(p.x, p.y, p.ch, p.fg, p.bg)
	}
	termbox.Flush()
}

func boardLayerCellMap() cellMap {
	boardPos := pos{5, 2}

	cm := cellMap{}
	cm.mergeCellMap(
		createBoardSkeleton(boardPos),
	)
	return cm
}

func (board *ludoBoard) setupBoard() {
	board.boardLayer = boardLayerCellMap()
	board.renderBoardLayer()
}
