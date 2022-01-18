package game

import "github.com/nsf/termbox-go"

func createBoardSkeleton(lx, rx, ty, by, boxLen, boxWid int, boardPos pos, players []player) cellMap {
	borderBox := box{
		pos:         pos{boardPos.x, boardPos.y},
		borderColor: termbox.ColorWhite,
		l:           17, w: 39,
	}

	relBoxPos := map[string]pos{
		"lt": {lx, ty},
		"rt": {rx, ty},
		"lb": {lx, by},
		"rb": {rx, by},
	}

	homeBorders := elementGroup{
		box{pos: relBoxPos["lt"], borderColor: players[0].color, l: boxLen, w: boxWid},
		box{pos: relBoxPos["rt"], borderColor: players[1].color, l: boxLen, w: boxWid},
		box{pos: relBoxPos["rb"], borderColor: players[3].color, l: boxLen, w: boxWid},
		box{pos: relBoxPos["lb"], borderColor: players[2].color, l: boxLen, w: boxWid},
		// middle
		box{pos: pos{relBoxPos["lt"].x + boxWid + 4, relBoxPos["lt"].y + boxLen + 3}, borderColor: termbox.ColorWhite, l: boxLen, w: boxWid},
	}

	cm := cellMap{}
	cm.mergeCellMap(
		homeBorders.toCellMap(),
		elementGroup{borderBox}.toCellMap(),
	)

	return cm
}

func createPawns(lx, rx, ty, by, boxLen, boxWid int, boardPos pos) []player {
	colors := []termbox.Attribute{
		termbox.ColorBlue, termbox.ColorRed, termbox.ColorYellow, termbox.ColorGreen,
	}

	var players []player

	relBoxPos := []pos{
		{lx, ty},
		{rx, ty},
		{lx, by},
		{rx, by},
	}

	for idx, color := range colors {
		pawns := [4]pawn{}

		x, y := relBoxPos[idx].x, relBoxPos[idx].y

		c := cell{x: x + 2, y: y + 1, bg: color, ch: ' '}
		homeNode := &node{cell: c}
		pawns[0] = pawn{
			"homeNode": homeNode,
			"curNode":  homeNode,
		}

		c = cell{x: x + 6, y: y + 1, bg: color, ch: ' '}
		homeNode = &node{cell: c}
		pawns[1] = pawn{
			"homeNode": homeNode,
			"curNode":  homeNode,
		}

		c = cell{x: x + 2, y: y + 3, bg: color, ch: ' '}
		homeNode = &node{cell: c}
		pawns[2] = pawn{
			"homeNode": homeNode,
			"curNode":  homeNode,
		}

		c = cell{x: x + 6, y: y + 3, bg: color, ch: ' '}
		homeNode = &node{cell: c}
		pawns[3] = pawn{
			"homeNode": homeNode,
			"curNode":  homeNode,
		}

		players = append(players, player{color: color, pawns: pawns, winningPos: -1})
	}

	return players
}
