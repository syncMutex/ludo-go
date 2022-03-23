package board

import (
	"ludo/src/common"
	ll "ludo/src/linked-list"
	"ludo/src/ludo-board/cellmap"

	"ludo/src/common/paths"

	"github.com/nsf/termbox-go"
)

type LudoBoard struct {
	Players    [4]Player
	boardLayer cellmap.CellMap
	pathLayer  paths.Path
	CurPawnIdx int
}

type Pawn map[string]*ll.Node

type Player struct {
	PlayerType string
	WinningPos int
	Color      termbox.Attribute
	Pawns      [4]Pawn
}

func (b *LudoBoard) setHomeBg() {
	for _, p := range b.Players {
		if !p.IsParticipant() {
			continue
		}
		for _, pawn := range p.Pawns {
			pawn["curNode"].Cell.Bg = p.Color
		}
	}
}

func (p Player) IsParticipant() bool {
	return p.PlayerType != "-"
}

func (p Player) IsBot() bool {
	return p.PlayerType == "Bot"
}

func boardLayerCellMap(lx, rx, ty, by, boxLen, boxWid int, boardPos cellmap.Pos, players [4]Player) cellmap.CellMap {
	cm := cellmap.CellMap{}
	cm.MergeCellMap(
		createBoardSkeleton(lx, rx, ty, by, boxLen, boxWid, boardPos, players),
	)
	return cm
}

func (p Player) IsAllPawnsAtDest() bool {
	for _, p := range p.Pawns {
		if !p.IsAtDest() {
			return false
		}
	}
	return true
}

func (board *LudoBoard) SetupBoard(players []common.PlayerData) {
	boardPos := cellmap.Pos{5, 2}
	lx, rx, ty, by := boardPos.X+2, boardPos.X+27, boardPos.Y+1, boardPos.Y+13

	boxLen, boxWid := 3, 9

	colorsList := [4]termbox.Attribute{}

	for i, p := range players {
		colorsList[i] = p.Color
	}

	board.Players = createPawns(lx, rx, ty, by, boxLen, boxWid, boardPos, players)
	board.boardLayer = boardLayerCellMap(lx, rx, ty, by, boxLen, boxWid, boardPos, board.Players)
	board.pathLayer = paths.CreatePathsLL(lx, rx, ty, by, boxLen, boxWid, boardPos, colorsList)
	board.connectPawnsPosToPath()
	board.setHomeBg()
}
