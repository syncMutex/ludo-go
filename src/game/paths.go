package game

type node struct {
	cell cell
	pawn *cell
	next *node
}

type path struct {
	ll linkedlist
}

func (p *path) extVertic(from pos, l, dir int) {
	pathSlice := elementGroup{fill{pos: from, w: 2, l: l, ch: 0x2591}}.toCellSlice()

	if dir == -1 {
		for i, j := 0, len(pathSlice)-1; i < j; i, j = i+1, j-1 {
			pathSlice[i], pathSlice[j] = pathSlice[j], pathSlice[i]
		}
	}

	for _, ele := range pathSlice {
		p.ll.addEnd(ele)
	}
}

func (p *path) extHoriz(from pos, w, dir int) {
	pathSlice := elementGroup{fill{pos: from, w: w, l: 1, ch: 0x2591}}.toCellSlice()

	if dir == -1 {
		for i, j := 0, len(pathSlice)-1; i < j; i, j = i+1, j-1 {
			pathSlice[i], pathSlice[j] = pathSlice[j], pathSlice[i]
		}
	}

	for _, ele := range pathSlice {
		p.ll.addEnd(ele)
	}
}

func createPathsLL(lx, rx, ty, by, boxLen, boxWid int, boardPos pos) path {
	paths := path{}

	paths.extVertic(pos{lx + boxWid + 6, by - 1}, 6, -1)

	paths.extHoriz(pos{lx, ty + boxLen + 6}, 12, -1)
	paths.extVertic(pos{lx, ty + boxLen + 5}, 2, -1)
	paths.extHoriz(pos{lx, ty + boxLen + 4}, 12, 1)

	paths.extVertic(pos{lx + boxWid + 6, ty}, 6, -1)
	paths.extHoriz(pos{lx + boxWid + 8, ty}, 2, -1)
	paths.extVertic(pos{lx + boxWid + 10, ty}, 6, 1)

	paths.extHoriz(pos{rx - 1, ty + boxLen + 4}, 12, 1)
	paths.extVertic(pos{rx - 1 + 10, ty + boxLen + 5}, 2, -1)
	paths.extHoriz(pos{rx - 1, ty + boxLen + 6}, 12, -1)

	paths.extVertic(pos{lx + boxWid + 10, by - 1}, 6, 1)
	paths.extHoriz(pos{lx + boxWid + 8, by - 1 + 5}, 2, -1)

	return paths
}
