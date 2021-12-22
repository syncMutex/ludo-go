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
	for color := range board.boardData {
		for _, p := range board.boardData[color] {
			termbox.SetCell(p.x, p.y, p.ch, p.fg, p.bg)
		}
	}
	termbox.Flush()
}

func createBoardData() colorMap {
	bx, by := 1, 1

	borderBox := box{
		pos:         pos{bx, by},
		borderColor: termbox.ColorWhite,
		l:           11, w: 27,
	}

	boxLen, boxWid := 3, 9

	homeBorders := elementGroup{
		box{pos: pos{bx + 2, by + 1}, borderColor: termbox.ColorBlue, l: boxLen, w: boxWid},
		box{pos: pos{bx + 16, by + 1}, borderColor: termbox.ColorGreen, l: boxLen, w: boxWid},
		box{pos: pos{bx + 16, by + 7}, borderColor: termbox.ColorRed, l: boxLen, w: boxWid},
		box{pos: pos{bx + 2, by + 7}, borderColor: termbox.ColorYellow, l: boxLen, w: boxWid},
	}

	pawns := elementGroup{
		point{color: termbox.ColorBlue, pos: pos{5, 3}},
		point{color: termbox.ColorBlue, pos: pos{9, 3}},
		point{color: termbox.ColorBlue, pos: pos{5, 5}},
		point{color: termbox.ColorBlue, pos: pos{9, 5}},

		point{color: termbox.ColorGreen, pos: pos{19, 3}},
		point{color: termbox.ColorGreen, pos: pos{23, 3}},
		point{color: termbox.ColorGreen, pos: pos{19, 5}},
		point{color: termbox.ColorGreen, pos: pos{23, 5}},

		point{color: termbox.ColorYellow, pos: pos{5, 9}},
		point{color: termbox.ColorYellow, pos: pos{9, 9}},
		point{color: termbox.ColorYellow, pos: pos{5, 11}},
		point{color: termbox.ColorYellow, pos: pos{9, 11}},

		point{color: termbox.ColorRed, pos: pos{19, 9}},
		point{color: termbox.ColorRed, pos: pos{23, 9}},
		point{color: termbox.ColorRed, pos: pos{19, 11}},
		point{color: termbox.ColorRed, pos: pos{23, 11}},
	}

	cm := colorMap{}

	cm.mergeColorMap(homeBorders.toColorMap(), elementGroup{borderBox}.toColorMap(), pawns.toColorMap())

	return cm
}

func (board *ludoBoard) setupBoard() {
	board.boardData = createBoardData()

	board.renderBoard()
}
