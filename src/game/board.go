package game

import "github.com/nsf/termbox-go"

type pos struct{ x, y int }

type ludoBoard struct {
	pawnsLocations []pawnData
	boardData      boardStruct
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

type boardStruct map[termbox.Attribute][]cell

type box struct {
	pos         pos
	borderColor termbox.Attribute
	fillColor   termbox.Attribute
	l, b        int
}

func (board *ludoBoard) renderBoard() {
	for color := range board.boardData {
		for _, p := range board.boardData[color] {
			termbox.SetCell(p.x, p.y, p.ch, p.fg, p.bg)
		}
	}
	termbox.Flush()
}

func getColorMap(boxes []box) boardStruct {
	colorMap := boardStruct{}
	for _, b := range boxes {
		// edges start
		colorMap[b.borderColor] = append(colorMap[b.borderColor], cell{
			x: b.pos.x, y: b.pos.y, ch: '┌', fg: b.borderColor, bg: b.fillColor,
		})
		colorMap[b.borderColor] = append(colorMap[b.borderColor], cell{
			x: b.pos.x + b.b, y: b.pos.y, ch: '┐', fg: b.borderColor, bg: b.fillColor,
		})
		colorMap[b.borderColor] = append(colorMap[b.borderColor], cell{
			x: b.pos.x, y: b.pos.y + b.l + 1, ch: '└', fg: b.borderColor, bg: b.fillColor,
		})
		colorMap[b.borderColor] = append(colorMap[b.borderColor], cell{
			x: b.pos.x + b.b, y: b.pos.y + b.l + 1, ch: '┘', fg: b.borderColor, bg: b.fillColor,
		})
		// edges end

		// along x-axis
		for i, x := 0, b.pos.x+1; i < b.b-1; i, x = i+1, x+1 {
			colorMap[b.borderColor] = append(colorMap[b.borderColor], cell{
				x: x, y: b.pos.y + b.l + 1, ch: '─', fg: b.borderColor, bg: b.fillColor,
			})
			colorMap[b.borderColor] = append(colorMap[b.borderColor], cell{
				x: x, y: b.pos.y, ch: '─', fg: b.borderColor, bg: b.fillColor,
			})
		}

		// along y-axiss
		for i, y := 0, b.pos.y+1; i < b.l; i, y = i+1, y+1 {
			colorMap[b.borderColor] = append(colorMap[b.borderColor], cell{
				x: b.pos.x + b.b, y: y, ch: '│', fg: b.borderColor, bg: b.fillColor,
			})
			colorMap[b.borderColor] = append(colorMap[b.borderColor], cell{
				x: b.pos.x, y: y, ch: '│', fg: b.borderColor, bg: b.fillColor,
			})
		}
	}

	return colorMap
}

func createBoardData() boardStruct {
	homeBoxes := []box{
		{
			pos:         pos{2, 1},
			borderColor: termbox.ColorBlue,
			l:           2, b: 7,
		},
		{
			pos:         pos{13, 1},
			borderColor: termbox.ColorGreen,
			l:           2, b: 7,
		},
		{
			pos:         pos{13, 6},
			borderColor: termbox.ColorRed,
			l:           2, b: 7,
		},
		{
			pos:         pos{2, 6},
			borderColor: termbox.ColorYellow,
			l:           2, b: 7,
		},
	}

	borderBox := []box{
		{
			pos:         pos{0, 0},
			borderColor: termbox.ColorWhite,
			l:           9, b: 22,
		},
	}

	colorMap := boardStruct{}

	cm := getColorMap(append(homeBoxes, borderBox...))

	for col := range cm {
		colorMap[col] = append(colorMap[col], cm[col]...)
	}

	return colorMap
}

func (board *ludoBoard) setupBoard() {
	board.boardData = createBoardData()

	board.renderBoard()
}
