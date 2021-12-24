package game

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

type cell struct {
	x, y int
	ch   rune
	bg   termbox.Attribute
	fg   termbox.Attribute
}

type colorMap map[string]cell

type box struct {
	pos         pos
	borderColor termbox.Attribute
	fillColor   termbox.Attribute
	l, w        int
}

type point struct {
	pos   pos
	color termbox.Attribute
	ch    rune
}

type fill struct {
	pos   pos
	l, w  int
	ch    rune
	color termbox.Attribute
}

type line struct {
	from   pos
	axis   rune // 'x'(horizontal) or 'y'(vertical)
	length int
	color  termbox.Attribute
}

func (c cell) mapKey() string {
	return strconv.Itoa(c.x) + strconv.Itoa(c.y)
}

func (cm colorMap) setCells(args ...cell) {
	for _, c := range args {
		cm[c.mapKey()] = c
	}
}

func boxColorMap(b box) colorMap {
	cm := colorMap{}

	// edges
	edges := []cell{
		{x: b.pos.x, y: b.pos.y, ch: '┌', fg: b.borderColor, bg: b.fillColor},
		{x: b.pos.x + b.w, y: b.pos.y, ch: '┐', fg: b.borderColor, bg: b.fillColor},
		{x: b.pos.x, y: b.pos.y + b.l + 1, ch: '└', fg: b.borderColor, bg: b.fillColor},
		{x: b.pos.x + b.w, y: b.pos.y + b.l + 1, ch: '┘', fg: b.borderColor, bg: b.fillColor},
	}

	cm.setCells(edges...)

	// along x-axis
	for i, x := 0, b.pos.x+1; i < b.w-1; i, x = i+1, x+1 {
		cm.setCells(
			cell{x: x, y: b.pos.y + b.l + 1, ch: '─', fg: b.borderColor, bg: b.fillColor},
			cell{x: x, y: b.pos.y, ch: '─', fg: b.borderColor, bg: b.fillColor},
		)
	}

	// along y-axiss
	for i, y := 0, b.pos.y+1; i < b.l; i, y = i+1, y+1 {
		cm.setCells(
			cell{x: b.pos.x + b.w, y: y, ch: '│', fg: b.borderColor, bg: b.fillColor},
			cell{x: b.pos.x, y: y, ch: '│', fg: b.borderColor, bg: b.fillColor},
		)
	}

	return cm
}

func pointColorMap(pt point) colorMap {
	cm := colorMap{}
	cm.setCells(
		cell{x: pt.pos.x, y: pt.pos.y, ch: pt.ch, fg: termbox.ColorDefault, bg: pt.color},
		cell{x: pt.pos.x + 1, y: pt.pos.y, ch: pt.ch, fg: termbox.ColorDefault, bg: pt.color},
	)
	return cm
}

func lineColorMap(ln line) colorMap {
	cm := colorMap{}

	switch ln.axis {
	case 'x':
		for i, x := 0, ln.from.x; i < ln.length; i, x = i+1, x+1 {
			cm.setCells(cell{x: x, y: ln.from.y, ch: '─', fg: ln.color})
		}
	case 'y':
		for i, y := 0, ln.from.y+1; i < ln.length; i, y = i+1, y+1 {
			cm.setCells(cell{x: ln.from.x, y: y, ch: '│', fg: ln.color})
		}
	}

	return cm
}

func fillColorMap(fl fill) colorMap {
	cm := colorMap{}

	y := fl.pos.y

	for i := 0; i < fl.l; i++ {
		for j, x := 0, fl.pos.x; j < fl.w; j++ {
			cm.setCells(cell{x: x, y: y, ch: fl.ch, fg: fl.color})
			x++
		}
		y++
	}

	return cm
}

func (cm colorMap) mergeColorMap(args ...colorMap) {
	for _, update := range args {
		for cell := range update {
			cm[update[cell].mapKey()] = update[cell]
		}
	}
}

func (e elementGroup) toColorMap() colorMap {
	cm := colorMap{}
	for _, ele := range e {
		switch e := ele.(type) {
		case box:
			cm.mergeColorMap(boxColorMap(e))
		case point:
			cm.mergeColorMap(pointColorMap(e))
		case line:
			cm.mergeColorMap(lineColorMap(e))
		case fill:
			cm.mergeColorMap(fillColorMap(e))
		}
	}

	return cm
}
