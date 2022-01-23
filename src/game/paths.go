package game

import (
	"github.com/nsf/termbox-go"
)

type path struct {
	ll linkedlist
}

type cellSlice []cell

type pathCell struct {
	pos    pos
	axis   rune
	length int
	dir    int
	ch     rune
	fg     termbox.Attribute
}

// 0x2591
func (pc pathCell) toPathSlice() cellSlice {
	cs := cellSlice{}

	switch pc.axis {
	case 'x':
		var x, dest int

		if pc.dir == 1 {
			x = pc.pos.x - 1
			dest = pc.pos.x + pc.length - 1
			for ; x < dest; x += 2 {
				cs = append(cs, cell{x: x, y: pc.pos.y, ch: 0x2591, fg: pc.fg})
			}
		} else {
			x = (pc.pos.x - 1) + pc.length - 1
			dest = pc.pos.x - 1
			for ; x > dest; x -= 2 {
				cs = append(cs, cell{x: x, y: pc.pos.y, ch: 0x2591, fg: pc.fg})
			}
		}
	case 'y':
		var y, dest int

		if pc.dir == 1 {
			y = pc.pos.y - 1
			dest = pc.pos.y + pc.length - 1
			for ; y < dest; y += 1 {
				cs = append(cs, cell{x: pc.pos.x, y: y, ch: 0x2591, fg: pc.fg})
			}
		} else {
			y = (pc.pos.y - 1) + pc.length
			dest = pc.pos.y - 1
			for ; y > dest; y -= 1 {
				cs = append(cs, cell{x: pc.pos.x, y: y, ch: 0x2591, fg: pc.fg})
			}
		}
	}

	return cs
}

func (p *path) extend(from pos, axis rune, length, dir int) *node {
	pathSlice := pathCell{pos: from, axis: axis, ch: 0x2591, length: length, dir: dir}.toPathSlice()
	lastNode := p.ll.head
	for _, ele := range pathSlice {
		lastNode = p.ll.addEnd(ele, "common", lastNode)
	}
	return lastNode
}

func (p *path) extToDest(from pos, axis rune, l, dir int, color termbox.Attribute, fromNode *node) {
	pathSlice := pathCell{pos: from, axis: axis, ch: 0x2591, length: l, dir: dir, fg: color}.toPathSlice()
	for _, ele := range pathSlice {
		fromNode = p.ll.addEnd(ele, "toDest", fromNode)
	}
}

func createPathsLL(lx, rx, ty, by, boxLen, boxWid int, boardPos pos, players [4]player) path {
	paths := path{}

	paths.extend(pos{lx + boxWid + 6, by - 1}, 'y', 6, -1)

	paths.extend(pos{lx, ty + boxLen + 6}, 'x', 12, -1)

	lastNode := paths.extend(pos{lx + 1, ty + boxLen + 5}, 'x', 2, 1)
	paths.extToDest(pos{lx + 3, ty + boxLen + 5}, 'x', 10, 1, players[0].color, lastNode)
	paths.extToDest(pos{lx + 15, ty + boxLen + 5}, 'x', 2, 1, players[0].color, lastNode)
	paths.extend(pos{lx + 1, ty + boxLen + 4}, 'x', 12, 1)

	paths.extend(pos{lx + boxWid + 6, ty}, 'y', 6, -1)

	lastNode = paths.extend(pos{lx + boxWid + 8, ty}, 'x', 2, -1)
	paths.extToDest(pos{lx + boxWid + 8, ty + 2}, 'y', 5, 1, players[1].color, lastNode)
	paths.extToDest(pos{lx + boxWid + 8, ty + 8}, 'y', 1, 1, players[1].color, lastNode)
	paths.extend(pos{lx + boxWid + 10, ty + 1}, 'y', 6, 1)

	paths.extend(pos{rx, ty + boxLen + 4}, 'x', 12, 1)

	lastNode = paths.extend(pos{rx - 1 + 10, ty + boxLen + 5}, 'x', 2, -1)
	paths.extToDest(pos{rx - 1, ty + boxLen + 5}, 'x', 10, -1, players[3].color, lastNode)
	paths.extToDest(pos{rx - 5, ty + boxLen + 5}, 'x', 2, -1, players[3].color, lastNode)
	paths.extend(pos{rx - 1, ty + boxLen + 6}, 'x', 12, -1)

	paths.extend(pos{lx + boxWid + 10, by}, 'y', 6, 1)
	lastNode = paths.extend(pos{lx + boxWid + 8, by - 1 + 5}, 'x', 2, -1)
	paths.extToDest(pos{lx + boxWid + 8, by - 1}, 'y', 5, -1, players[2].color, lastNode)
	paths.extToDest(pos{lx + boxWid + 8, by - 3}, 'y', 1, -1, players[2].color, lastNode)

	lastNode.next["common"] = paths.ll.head

	return paths
}
