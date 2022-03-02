package game

import (
	gameUtils "ludo/src/game-utils"

	"github.com/nsf/termbox-go"
)

type pos struct{ x, y int }

type ludoBoard struct {
	players    [4]player
	boardLayer cellMap
	pathLayer  path
	curPawnIdx int
}

type pawn map[string]*node

type player struct {
	playerType string
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
		if !p.isParticipant() {
			continue
		}
		for _, pawn := range p.pawns {
			pawn["curNode"].cell.bg = p.color
		}
	}
}

func (p player) isParticipant() bool {
	return p.playerType != "-"
}

func (p player) isBot() bool {
	return p.playerType == "Bot"
}

func boardLayerCellMap(lx, rx, ty, by, boxLen, boxWid int, boardPos pos, players [4]player) cellMap {
	cm := cellMap{}
	cm.mergeCellMap(
		createBoardSkeleton(lx, rx, ty, by, boxLen, boxWid, boardPos, players),
	)
	return cm
}

func (board *ludoBoard) setupBoard(players []gameUtils.PlayerData) {
	boardPos := pos{5, 2}
	lx, rx, ty, by := boardPos.x+2, boardPos.x+27, boardPos.y+1, boardPos.y+13

	boxLen, boxWid := 3, 9

	board.players = createPawns(lx, rx, ty, by, boxLen, boxWid, boardPos, players)
	board.boardLayer = boardLayerCellMap(lx, rx, ty, by, boxLen, boxWid, boardPos, board.players)
	board.pathLayer = createPathsLL(lx, rx, ty, by, boxLen, boxWid, boardPos, board.players)
	board.connectPawnsPosToPath()
	board.setHomeBg()
}
