package game

import (
	"github.com/nsf/termbox-go"
)

type pos struct{ x, y int }

type ludoBoard struct {
	players    []player
	boardLayer cellMap
	pathLayer  path
	curPawnIdx int
}

type pawn map[string]*node

type player struct {
	winningPos int
	color      termbox.Attribute
	pawns      [4]pawn
}

type elementGroup []interface{}

func (b *ludoBoard) render() {
	b.renderBoardLayer()
	b.renderPathLayer()
	b.renderPawns()
}

func (b *ludoBoard) setHomeBg() {
	for _, p := range b.players {
		for _, pawn := range p.pawns {
			pawn["curNode"].cell.bg = p.color
		}
	}
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
	board.connectPawnsPosToPath()
	board.setHomeBg()
}
