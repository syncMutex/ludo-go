package game

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
}

func (cs *cellSlice) mergeCellSlice(args ...cellSlice) {
	for _, update := range args {
		*cs = append(*cs, update...)
	}
}

func (pc pathCell) toPathSlice() cellSlice {
	cs := cellSlice{}

	switch pc.axis {
	case 'x':
		var x, dest int

		if pc.dir == 1 {
			x = pc.pos.x - 1
			dest = pc.pos.x + pc.length - 1
			for ; x < dest; x += 2 {
				cs = append(cs, cell{x: x, y: pc.pos.y, ch: 0x2591})
			}
		} else {
			x = (pc.pos.x - 1) + pc.length - 1
			dest = pc.pos.x - 1
			for ; x > dest; x -= 2 {
				cs = append(cs, cell{x: x, y: pc.pos.y, ch: 0x2591})
			}
		}
	case 'y':
		var y, dest int

		if pc.dir == 1 {
			y = pc.pos.y - 1
			dest = pc.pos.y + pc.length - 1
			for ; y < dest; y += 1 {
				cs = append(cs, cell{x: pc.pos.x, y: y, ch: 0x2591})
			}
		} else {
			y = (pc.pos.y - 1) + pc.length
			dest = pc.pos.y - 1
			for ; y > dest; y -= 1 {
				cs = append(cs, cell{x: pc.pos.x, y: y, ch: 0x2591})
			}
		}
	}

	return cs
}

func (p *path) extVertic(from pos, l, dir int) {
	pathSlice := pathCell{pos: from, axis: 'y', ch: 0x2591, length: l, dir: dir}.toPathSlice()

	for _, ele := range pathSlice {
		p.ll.addEnd(ele)
	}
}

func (p *path) extHoriz(from pos, w, dir int) {
	pathSlice := pathCell{pos: from, axis: 'x', ch: 0x2591, length: w, dir: dir}.toPathSlice()

	for _, ele := range pathSlice {
		p.ll.addEnd(ele)
	}
}

func createPathsLL(lx, rx, ty, by, boxLen, boxWid int, boardPos pos) path {
	paths := path{}

	paths.extVertic(pos{lx + boxWid + 6, by - 1}, 6, -1)

	paths.extHoriz(pos{lx, ty + boxLen + 6}, 12, -1)
	paths.extHoriz(pos{lx + 1, ty + boxLen + 5}, 2, 1)
	paths.extHoriz(pos{lx + 1, ty + boxLen + 4}, 12, 1)

	paths.extVertic(pos{lx + boxWid + 6, ty}, 6, -1)
	paths.extHoriz(pos{lx + boxWid + 8, ty}, 2, -1)
	paths.extVertic(pos{lx + boxWid + 10, ty + 1}, 6, 1)

	paths.extHoriz(pos{rx, ty + boxLen + 4}, 12, 1)
	paths.extHoriz(pos{rx - 1 + 10, ty + boxLen + 5}, 2, -1)
	paths.extHoriz(pos{rx - 1, ty + boxLen + 6}, 12, -1)

	paths.extVertic(pos{lx + boxWid + 10, by}, 6, 1)
	paths.extHoriz(pos{lx + boxWid + 8, by - 1 + 5}, 2, -1)

	return paths
}
