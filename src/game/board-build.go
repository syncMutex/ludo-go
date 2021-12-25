package game

import "github.com/nsf/termbox-go"

func createBoardSkeleton(lx, rx, ty, by, boxLen, boxWid int, boardPos pos) cellMap {
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
		box{pos: relBoxPos["lt"], borderColor: termbox.ColorBlue, l: boxLen, w: boxWid},
		box{pos: relBoxPos["rt"], borderColor: termbox.ColorGreen, l: boxLen, w: boxWid},
		box{pos: relBoxPos["rb"], borderColor: termbox.ColorRed, l: boxLen, w: boxWid},
		box{pos: relBoxPos["lb"], borderColor: termbox.ColorYellow, l: boxLen, w: boxWid},
		box{pos: pos{relBoxPos["lt"].x + boxWid + 4, relBoxPos["lt"].y + boxLen + 3}, borderColor: termbox.ColorWhite, l: boxLen, w: boxWid},
	}

	// 0x259

	// paths := elementGroup{
	// 	fill{pos: pos{lx + boxWid + 6, ty}, l: 6, w: 6, ch: 0x2591},
	// 	fill{pos: pos{lx + boxWid + 8, ty}, l: 6, w: 2, ch: 0x2591, color: termbox.ColorGreen},
	// 	fill{pos: pos{lx, ty + boxLen + 4}, l: 3, w: 12, ch: 0x2591},
	// 	fill{pos: pos{lx, ty + boxLen + 5}, l: 1, w: 12, ch: 0x2591, color: termbox.ColorBlue},
	// 	fill{pos: pos{rx - 1, ty + boxLen + 4}, l: 3, w: 12, ch: 0x2591},
	// 	fill{pos: pos{rx - 1, ty + boxLen + 5}, l: 1, w: 12, ch: 0x2591, color: termbox.ColorRed},
	// 	fill{pos: pos{lx + boxWid + 6, by - 1}, l: 6, w: 6, ch: 0x2591},
	// 	fill{pos: pos{lx + boxWid + 8, by - 1}, l: 6, w: 2, ch: 0x2591, color: termbox.ColorYellow},
	// }

	cm := cellMap{}
	cm.mergeCellMap(
		homeBorders.toCellMap(),
		elementGroup{borderBox}.toCellMap(),
	)

	return cm
}
