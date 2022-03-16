package cellmap

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

type Cell struct {
	X, Y int
	Ch   rune
	Bg   termbox.Attribute
	Fg   termbox.Attribute
}

type Pos struct{ X, Y int }
type CellMap map[string]Cell
type ElementGroup []interface{}

type Box struct {
	Pos         Pos
	BorderColor termbox.Attribute
	FillColor   termbox.Attribute
	L, W        int
}

type Point struct {
	Pos   Pos
	Color termbox.Attribute
	Ch    rune
}

type Fill struct {
	Pos   Pos
	L, W  int
	Ch    rune
	Color termbox.Attribute
}

type Line struct {
	From   Pos
	Axis   rune // 'x'(horizontal) or 'y'(vertical)
	Length int
	Color  termbox.Attribute
}

func (c Cell) MapKey() string {
	return strconv.Itoa(c.X) + strconv.Itoa(c.Y)
}

func (cm CellMap) setCells(args ...Cell) {
	for _, c := range args {
		cm[c.MapKey()] = c
	}
}

func boxCellMap(b Box) CellMap {
	cm := CellMap{}

	// edges
	edges := []Cell{
		{X: b.Pos.X, Y: b.Pos.Y, Ch: '┌', Fg: b.BorderColor, Bg: b.FillColor},
		{X: b.Pos.X + b.W, Y: b.Pos.Y, Ch: '┐', Fg: b.BorderColor, Bg: b.FillColor},
		{X: b.Pos.X, Y: b.Pos.Y + b.L + 1, Ch: '└', Fg: b.BorderColor, Bg: b.FillColor},
		{X: b.Pos.X + b.W, Y: b.Pos.Y + b.L + 1, Ch: '┘', Fg: b.BorderColor, Bg: b.FillColor},
	}

	cm.setCells(edges...)

	// along x-axis
	for i, x := 0, b.Pos.X+1; i < b.W-1; i, x = i+1, x+1 {
		cm.setCells(
			Cell{X: x, Y: b.Pos.Y + b.L + 1, Ch: '─', Fg: b.BorderColor, Bg: b.FillColor},
			Cell{X: x, Y: b.Pos.Y, Ch: '─', Fg: b.BorderColor, Bg: b.FillColor},
		)
	}

	// along y-axiss
	for i, y := 0, b.Pos.Y+1; i < b.L; i, y = i+1, y+1 {
		cm.setCells(
			Cell{X: b.Pos.X + b.W, Y: y, Ch: '│', Fg: b.BorderColor, Bg: b.FillColor},
			Cell{X: b.Pos.X, Y: y, Ch: '│', Fg: b.BorderColor, Bg: b.FillColor},
		)
	}

	return cm
}

func PointCellMap(pt Point) CellMap {
	cm := CellMap{}
	cm.setCells(
		Cell{X: pt.Pos.X, Y: pt.Pos.Y, Ch: pt.Ch, Fg: termbox.ColorDefault, Bg: pt.Color},
		Cell{X: pt.Pos.X + 1, Y: pt.Pos.Y, Ch: pt.Ch, Fg: termbox.ColorDefault, Bg: pt.Color},
	)
	return cm
}

func LineCellMap(ln Line) CellMap {
	cm := CellMap{}

	switch ln.Axis {
	case 'x':
		for i, x := 0, ln.From.X; i < ln.Length; i, x = i+1, x+1 {
			cm.setCells(Cell{X: x, Y: ln.From.Y, Ch: '─', Fg: ln.Color})
		}
	case 'y':
		for i, y := 0, ln.From.Y+1; i < ln.Length; i, y = i+1, y+1 {
			cm.setCells(Cell{X: ln.From.X, Y: y, Ch: '│', Fg: ln.Color})
		}
	}

	return cm
}

func FillCellMap(fl Fill) CellMap {
	cm := CellMap{}

	y := fl.Pos.Y

	for i := 0; i < fl.L; i++ {
		for j, x := 0, fl.Pos.X; j < fl.W; j++ {
			cm.setCells(Cell{X: x, Y: y, Ch: fl.Ch, Fg: fl.Color})
			x++
		}
		y++
	}

	return cm
}

func (cm CellMap) MergeCellMap(args ...CellMap) {
	for _, update := range args {
		for Cell := range update {
			cm[update[Cell].MapKey()] = update[Cell]
		}
	}
}

func (e ElementGroup) ToCellMap() CellMap {
	cm := CellMap{}
	for _, ele := range e {
		switch e := ele.(type) {
		case Box:
			cm.MergeCellMap(boxCellMap(e))
		case Point:
			cm.MergeCellMap(PointCellMap(e))
		case Line:
			cm.MergeCellMap(LineCellMap(e))
		case Fill:
			cm.MergeCellMap(FillCellMap(e))
		}
	}

	return cm
}
