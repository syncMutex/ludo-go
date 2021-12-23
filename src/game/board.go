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
	boardPos := pos{5, 2}

	borderBox := box{
		pos:         pos{boardPos.x, boardPos.y},
		borderColor: termbox.ColorWhite,
		l:           17, w: 38,
	}

	boxLen, boxWid := 3, 9

	lx, rx, ty, by := boardPos.x+2, boardPos.x+26, boardPos.y+1, boardPos.y+13

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
		box{pos: pos{relBoxPos["lt"].x + boxWid + 3, relBoxPos["lt"].y + boxLen + 3}, borderColor: termbox.ColorWhite, l: boxLen, w: boxWid},
	}

	cm := colorMap{}

	cm.mergeColorMap(
		homeBorders.toColorMap(),
		elementGroup{borderBox}.toColorMap(),
	)

	return cm
}

func (board *ludoBoard) setupBoard() {
	board.boardData = createBoardData()

	board.renderBoard()
}
