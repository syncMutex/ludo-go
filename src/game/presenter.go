package game

import "github.com/nsf/termbox-go"

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
	l, w        int
}

type point struct {
	pos   pos
	color termbox.Attribute
}

type line struct {
	from   pos
	axis   rune // 'x'(horizontal) or 'y'(vertical)
	length int
	color  termbox.Attribute
}

func boxColorMap(b box) colorMap {
	cm := colorMap{}

	// edges
	cm[b.borderColor] = append(cm[b.borderColor], cell{
		x: b.pos.x, y: b.pos.y, ch: '┌', fg: b.borderColor, bg: b.fillColor,
	})
	cm[b.borderColor] = append(cm[b.borderColor], cell{
		x: b.pos.x + b.w, y: b.pos.y, ch: '┐', fg: b.borderColor, bg: b.fillColor,
	})
	cm[b.borderColor] = append(cm[b.borderColor], cell{
		x: b.pos.x, y: b.pos.y + b.l + 1, ch: '└', fg: b.borderColor, bg: b.fillColor,
	})
	cm[b.borderColor] = append(cm[b.borderColor], cell{
		x: b.pos.x + b.w, y: b.pos.y + b.l + 1, ch: '┘', fg: b.borderColor, bg: b.fillColor,
	})

	// along x-axis
	for i, x := 0, b.pos.x+1; i < b.w-1; i, x = i+1, x+1 {
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
			x: b.pos.x + b.w, y: y, ch: '│', fg: b.borderColor, bg: b.fillColor,
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

func lineColorMap(ln line) colorMap {
	cm := colorMap{}

	switch ln.axis {
	case 'x':
		for i, x := 0, ln.from.x; i < ln.length; i, x = i+1, x+1 {
			cm[ln.color] = append(cm[ln.color], cell{
				x: x, y: ln.from.y, ch: '─', fg: ln.color,
			})
		}
	case 'y':
		for i, y := 0, ln.from.y+1; i < ln.length; i, y = i+1, y+1 {
			cm[ln.color] = append(cm[ln.color], cell{
				x: ln.from.x, y: y, ch: '│', fg: ln.color,
			})
		}
	}

	return cm
}

func (cm colorMap) mergeColorMap(args ...colorMap) {
	for _, update := range args {
		for col := range update {
			cm[col] = append(cm[col], update[col]...)
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
		}
	}

	return cm
}
