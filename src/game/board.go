package game

import (
	"github.com/nsf/termbox-go"
)

type pos struct{ x, y int }

type ludoBoard struct {
	players    []player
	boardLayer cellMap
	pathLayer  path
}

type player struct {
	color termbox.Attribute
	pawns [4]map[string]*node
}

type elementGroup []interface{}

func (b *ludoBoard) render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	b.renderBoardLayer()
	b.renderPathLayer()
	b.renderHomeNodes()
	termbox.Flush()
}

func boardLayerCellMap(lx, rx, ty, by, boxLen, boxWid int, boardPos pos, players []player) cellMap {
	cm := cellMap{}
	cm.mergeCellMap(
		createBoardSkeleton(lx, rx, ty, by, boxLen, boxWid, boardPos, players),
	)
	return cm
}

func (board *ludoBoard) setupBoard() {
	boardPos := pos{5, 2}
	lx, rx, ty, by := boardPos.x+2, boardPos.x+27, boardPos.y+1, boardPos.y+13

	boxLen, boxWid := 3, 9

	board.players = createPawns(lx, rx, ty, by, boxLen, boxWid, boardPos)
	board.boardLayer = boardLayerCellMap(lx, rx, ty, by, boxLen, boxWid, boardPos, board.players)
	board.pathLayer = createPathsLL(lx, rx, ty, by, boxLen, boxWid, boardPos, board.players)
	board.renderBoardLayer()
}
