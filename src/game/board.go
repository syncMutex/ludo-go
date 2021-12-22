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

type cell struct {
	x, y int
	ch   rune
	bg   termbox.Attribute
	fg   termbox.Attribute
}

type colorMap map[termbox.Attribute][]cell

type box struct {
	pos         pos
	borderColor termbox.Attribute
	fillColor   termbox.Attribute
	l, b        int
}

type point struct {
	pos   pos
	color termbox.Attribute
}

func (board *ludoBoard) renderBoard() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for color := range board.boardData {
		for _, p := range board.boardData[color] {
			termbox.SetCell(p.x, p.y, p.ch, p.fg, p.bg)
		}
	}
	termbox.Flush()
}

func boxColorMap(b box) colorMap {
	cm := colorMap{}
	// edges start
	cm[b.borderColor] = append(cm[b.borderColor], cell{
		x: b.pos.x, y: b.pos.y, ch: '┌', fg: b.borderColor, bg: b.fillColor,
	})
	cm[b.borderColor] = append(cm[b.borderColor], cell{
		x: b.pos.x + b.b, y: b.pos.y, ch: '┐', fg: b.borderColor, bg: b.fillColor,
	})
	cm[b.borderColor] = append(cm[b.borderColor], cell{
		x: b.pos.x, y: b.pos.y + b.l + 1, ch: '└', fg: b.borderColor, bg: b.fillColor,
	})
	cm[b.borderColor] = append(cm[b.borderColor], cell{
		x: b.pos.x + b.b, y: b.pos.y + b.l + 1, ch: '┘', fg: b.borderColor, bg: b.fillColor,
	})
	// edges end

	// along x-axis
	for i, x := 0, b.pos.x+1; i < b.b-1; i, x = i+1, x+1 {
		cm[b.borderColor] = append(cm[b.borderColor], cell{
			x: x, y: b.pos.y + b.l + 1, ch: '─', fg: b.borderColor, bg: b.fillColor,
		})
		cm[b.borderColor] = append(cm[b.borderColor], cell{
			x: x, y: b.pos.y, ch: '─', fg: b.borderColor, bg: b.fillColor,
		})
	}

	// along y-axiss
	for i, y := 0, b.pos.y+1; i < b.l; i, y = i+1, y+1 {
		cm[b.borderColor] = append(cm[b.borderColor], cell{
			x: b.pos.x + b.b, y: y, ch: '│', fg: b.borderColor, bg: b.fillColor,
		})
		cm[b.borderColor] = append(cm[b.borderColor], cell{
			x: b.pos.x, y: y, ch: '│', fg: b.borderColor, bg: b.fillColor,
		})
	}

	return cm
}

func pointColorMap(pt point) colorMap {
	cm := colorMap{}
	cm[pt.color] = append(cm[pt.color], cell{x: pt.pos.x, y: pt.pos.y, ch: ' ', fg: termbox.ColorDefault, bg: pt.color})
	cm[pt.color] = append(cm[pt.color], cell{x: pt.pos.x + 1, y: pt.pos.y, ch: ' ', fg: termbox.ColorDefault, bg: pt.color})
	return cm
}

func getColorMap(elements []interface{}) colorMap {
	cm := colorMap{}
	for _, ele := range elements {
		switch e := ele.(type) {
		case box:
			cm.updateColorMap(boxColorMap(e))
		case point:
			cm.updateColorMap(pointColorMap(e))
		}
	}

	return cm
}

func (cm colorMap) updateColorMap(update colorMap) {
	for col := range update {
		cm[col] = append(cm[col], update[col]...)
	}
}

func createBoardData() colorMap {
	bx, by := 1, 1

	borderBox := box{
		pos:         pos{bx, by},
		borderColor: termbox.ColorWhite,
		l:           11, b: 26,
	}

	boxLen, boxWid := 3, 8

	homeBoxes := []interface{}{
		box{pos: pos{bx + 2, by + 1}, borderColor: termbox.ColorBlue, l: boxLen, b: boxWid},
		box{pos: pos{bx + 16, by + 1}, borderColor: termbox.ColorGreen, l: boxLen, b: boxWid},
		box{pos: pos{bx + 16, by + 7}, borderColor: termbox.ColorRed, l: boxLen, b: boxWid},
		box{pos: pos{bx + 2, by + 7}, borderColor: termbox.ColorYellow, l: boxLen, b: boxWid},
	}

	pawns := []interface{}{
		point{color: termbox.ColorBlue, pos: pos{5, 3}},
		point{color: termbox.ColorBlue, pos: pos{8, 3}},
		point{color: termbox.ColorBlue, pos: pos{5, 5}},
		point{color: termbox.ColorBlue, pos: pos{8, 5}},

		point{color: termbox.ColorGreen, pos: pos{19, 3}},
		point{color: termbox.ColorGreen, pos: pos{22, 3}},
		point{color: termbox.ColorGreen, pos: pos{19, 5}},
		point{color: termbox.ColorGreen, pos: pos{22, 5}},

		point{color: termbox.ColorYellow, pos: pos{5, 9}},
		point{color: termbox.ColorYellow, pos: pos{8, 9}},
		point{color: termbox.ColorYellow, pos: pos{5, 11}},
		point{color: termbox.ColorYellow, pos: pos{8, 11}},

		point{color: termbox.ColorRed, pos: pos{19, 9}},
		point{color: termbox.ColorRed, pos: pos{22, 9}},
		point{color: termbox.ColorRed, pos: pos{19, 11}},
		point{color: termbox.ColorRed, pos: pos{22, 11}},
	}

	cm := colorMap{}

	c := getColorMap(
		append(
			append(
				append(
					[]interface{}{}, homeBoxes...,
				), borderBox,
			), pawns...,
		),
	)

	cm.updateColorMap(c)

	return cm
}

func (board *ludoBoard) setupBoard() {
	board.boardData = createBoardData()

	board.renderBoard()
}
