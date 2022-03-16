package board

import (
	"ludo/src/common"
	ll "ludo/src/linked-list"
	"ludo/src/ludo-board/cellmap"

	"github.com/nsf/termbox-go"
)

func createBoardSkeleton(lx, rx, ty, by, boxLen, boxWid int, boardPos cellmap.Pos, players [4]Player) cellmap.CellMap {
	borderBox := cellmap.Box{
		Pos:         cellmap.Pos{boardPos.X, boardPos.Y},
		BorderColor: termbox.ColorWhite,
		L:           17, W: 39,
	}

	relBoxPos := map[string]cellmap.Pos{
		"lt": {lx, ty},
		"rt": {rx, ty},
		"lb": {lx, by},
		"rb": {rx, by},
	}

	homeBorders := cellmap.ElementGroup{
		cellmap.Box{Pos: relBoxPos["lt"], BorderColor: players[0].Color, L: boxLen, W: boxWid},
		cellmap.Box{Pos: relBoxPos["rt"], BorderColor: players[1].Color, L: boxLen, W: boxWid},
		cellmap.Box{Pos: relBoxPos["rb"], BorderColor: players[3].Color, L: boxLen, W: boxWid},
		cellmap.Box{Pos: relBoxPos["lb"], BorderColor: players[2].Color, L: boxLen, W: boxWid},
		// middle
		cellmap.Box{Pos: cellmap.Pos{relBoxPos["lt"].X + boxWid + 4, relBoxPos["lt"].Y + boxLen + 3}, BorderColor: termbox.ColorWhite, L: boxLen, W: boxWid},
	}

	cm := cellmap.CellMap{}
	cm.MergeCellMap(
		homeBorders.ToCellMap(),
		cellmap.ElementGroup{borderBox}.ToCellMap(),
	)

	return cm
}

func createPawns(lx, rx, ty, by, boxLen, boxWid int, boardPos cellmap.Pos, playersData []common.PlayerData) [4]Player {
	var players [4]Player

	relBoxPos := []cellmap.Pos{
		{lx, ty},
		{rx, ty},
		{lx, by},
		{rx, by},
	}

	for idx, p := range playersData {
		if p.Type == "-" {
			players[idx].PlayerType = p.Type
			players[idx].Color = p.Color
			continue
		}

		pawns := [4]Pawn{}

		x, y := relBoxPos[idx].X, relBoxPos[idx].Y

		c := cellmap.Cell{X: x + 2, Y: y + 1, Bg: p.Color, Ch: ' '}
		homeNode := &ll.Node{Cell: c}
		pawns[0] = Pawn{
			"homeNode": homeNode,
			"curNode":  homeNode,
		}

		c = cellmap.Cell{X: x + 6, Y: y + 1, Bg: p.Color, Ch: ' '}
		homeNode = &ll.Node{Cell: c}
		pawns[1] = Pawn{
			"homeNode": homeNode,
			"curNode":  homeNode,
		}

		c = cellmap.Cell{X: x + 2, Y: y + 3, Bg: p.Color, Ch: ' '}
		homeNode = &ll.Node{Cell: c}
		pawns[2] = Pawn{
			"homeNode": homeNode,
			"curNode":  homeNode,
		}

		c = cellmap.Cell{X: x + 6, Y: y + 3, Bg: p.Color, Ch: ' '}
		homeNode = &ll.Node{Cell: c}
		pawns[3] = Pawn{
			"homeNode": homeNode,
			"curNode":  homeNode,
		}

		players[idx] = Player{Color: p.Color, Pawns: pawns, WinningPos: -1, PlayerType: p.Type}
	}

	return players
}
