package paths

import (
	ll "ludo/src/linked-list"
	"ludo/src/ludo-board/cellmap"

	"github.com/nsf/termbox-go"
)

type Path struct {
	LL ll.Linkedlist
}

type CellSlice []cellmap.Cell

type PathCell struct {
	Pos    cellmap.Pos
	Axis   rune
	Length int
	Dir    int
	Ch     rune
	Fg     termbox.Attribute
}

// 0x2591
func (pc PathCell) toPathSlice() CellSlice {
	cs := CellSlice{}

	switch pc.Axis {
	case 'x':
		var x, dest int

		if pc.Dir == 1 {
			x = pc.Pos.X - 1
			dest = pc.Pos.X + pc.Length - 1
			for ; x < dest; x += 2 {
				cs = append(cs, cellmap.Cell{X: x, Y: pc.Pos.Y, Ch: 0x2591, Fg: pc.Fg})
			}
		} else {
			x = (pc.Pos.X - 1) + pc.Length - 1
			dest = pc.Pos.X - 1
			for ; x > dest; x -= 2 {
				cs = append(cs, cellmap.Cell{X: x, Y: pc.Pos.Y, Ch: 0x2591, Fg: pc.Fg})
			}
		}
	case 'y':
		var y, dest int

		if pc.Dir == 1 {
			y = pc.Pos.Y - 1
			dest = pc.Pos.Y + pc.Length - 1
			for ; y < dest; y += 1 {
				cs = append(cs, cellmap.Cell{X: pc.Pos.X, Y: y, Ch: 0x2591, Fg: pc.Fg})
			}
		} else {
			y = (pc.Pos.Y - 1) + pc.Length
			dest = pc.Pos.Y - 1
			for ; y > dest; y -= 1 {
				cs = append(cs, cellmap.Cell{X: pc.Pos.X, Y: y, Ch: 0x2591, Fg: pc.Fg})
			}
		}
	}

	return cs
}

func (p *Path) Extend(from cellmap.Pos, axis rune, length, dir int) *ll.Node {
	PathSlice := PathCell{Pos: from, Axis: axis, Ch: 0x2591, Length: length, Dir: dir}.toPathSlice()
	lastNode := p.LL.Head
	for _, ele := range PathSlice {
		lastNode = p.LL.AddEnd(ele, "common", lastNode)
	}
	return lastNode
}

func (p *Path) extToDest(from cellmap.Pos, axis rune, l, dir int, color termbox.Attribute, fromNode *ll.Node) {
	PathSlice := PathCell{Pos: from, Axis: axis, Ch: 0x2591, Length: l, Dir: dir, Fg: color}.toPathSlice()
	for _, ele := range PathSlice {
		fromNode = p.LL.AddEnd(ele, "toDest", fromNode)
	}
}

func CreatePathsLL(lx, rx, ty, by, boxLen, boxWid int, boardPos cellmap.Pos, colors [4]termbox.Attribute) Path {
	paths := Path{}

	paths.Extend(cellmap.Pos{lx + boxWid + 6, by - 1}, 'y', 6, -1)

	paths.Extend(cellmap.Pos{lx, ty + boxLen + 6}, 'x', 12, -1)

	lastNode := paths.Extend(cellmap.Pos{lx + 1, ty + boxLen + 5}, 'x', 2, 1)
	paths.extToDest(cellmap.Pos{lx + 3, ty + boxLen + 5}, 'x', 10, 1, colors[0], lastNode)
	paths.extToDest(cellmap.Pos{lx + 15, ty + boxLen + 5}, 'x', 2, 1, colors[0], lastNode)
	paths.Extend(cellmap.Pos{lx + 1, ty + boxLen + 4}, 'x', 12, 1)

	paths.Extend(cellmap.Pos{lx + boxWid + 6, ty}, 'y', 6, -1)

	lastNode = paths.Extend(cellmap.Pos{lx + boxWid + 8, ty}, 'x', 2, -1)
	paths.extToDest(cellmap.Pos{lx + boxWid + 8, ty + 2}, 'y', 5, 1, colors[1], lastNode)
	paths.extToDest(cellmap.Pos{lx + boxWid + 8, ty + 8}, 'y', 1, 1, colors[1], lastNode)
	paths.Extend(cellmap.Pos{lx + boxWid + 10, ty + 1}, 'y', 6, 1)

	paths.Extend(cellmap.Pos{rx, ty + boxLen + 4}, 'x', 12, 1)

	lastNode = paths.Extend(cellmap.Pos{rx - 1 + 10, ty + boxLen + 5}, 'x', 2, -1)
	paths.extToDest(cellmap.Pos{rx - 1, ty + boxLen + 5}, 'x', 10, -1, colors[3], lastNode)
	paths.extToDest(cellmap.Pos{rx - 5, ty + boxLen + 5}, 'x', 2, -1, colors[3], lastNode)
	paths.Extend(cellmap.Pos{rx - 1, ty + boxLen + 6}, 'x', 12, -1)

	paths.Extend(cellmap.Pos{lx + boxWid + 10, by}, 'y', 6, 1)
	lastNode = paths.Extend(cellmap.Pos{lx + boxWid + 8, by - 1 + 5}, 'x', 2, -1)
	paths.extToDest(cellmap.Pos{lx + boxWid + 8, by - 1}, 'y', 5, -1, colors[2], lastNode)
	paths.extToDest(cellmap.Pos{lx + boxWid + 8, by - 3}, 'y', 1, -1, colors[2], lastNode)

	lastNode.Next["common"] = paths.LL.Head

	return paths
}
