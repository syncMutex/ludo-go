package game

import (
	"github.com/nsf/termbox-go"
)

type pos struct{ x, y int }

type ludoBoard struct {
	pawnsLocations []pawnData
	boardLayer     cellMap
	pathLayer      path
}

type pawnData struct {
	color    termbox.Attribute
	pawnsPos [4]pos
}

type elementGroup []interface{}

func (b *ludoBoard) renderBoardLayer() {
	for _, p := range b.boardLayer {
		termbox.SetCell(p.x, p.y, p.ch, p.fg, p.bg)
	}
}

func (b *ludoBoard) render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	b.renderBoardLayer()
	b.renderPathLayer()
	termbox.Flush()
}

func setPathCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
	termbox.SetCell(x+1, y, ch, fg, bg)
}

func (b *ludoBoard) renderPathLayer() {
	temp := b.pathLayer.ll.head

	for temp != nil {
		if temp.pawn != nil {
			setPathCell(temp.cell.x, temp.cell.y, ' ', termbox.ColorDefault, temp.pawn.bg)
		} else {
			setPathCell(temp.cell.x, temp.cell.y, temp.cell.ch, temp.cell.fg, temp.cell.bg)
		}
		if temp.next["toHome"] != nil {
			temp2 := temp.next["toHome"]
			for temp2 != nil {
				if temp2.pawn != nil {
					setPathCell(temp2.cell.x, temp2.cell.y, ' ', termbox.ColorDefault, temp2.pawn.bg)
				} else {
					setPathCell(temp2.cell.x, temp2.cell.y, temp2.cell.ch, temp2.cell.fg, temp2.cell.bg)
				}
				temp2 = temp2.next["toHome"]
			}
		}
		temp = temp.next["common"]
	}
}

func boardLayerCellMap(lx, rx, ty, by, boxLen, boxWid int, boardPos pos) cellMap {
	cm := cellMap{}
	cm.mergeCellMap(
		createBoardSkeleton(lx, rx, ty, by, boxLen, boxWid, boardPos),
	)
	return cm
}

func (board *ludoBoard) setupBoard() {
	boardPos := pos{5, 2}
	lx, rx, ty, by := boardPos.x+2, boardPos.x+27, boardPos.y+1, boardPos.y+13

	boxLen, boxWid := 3, 9

	board.boardLayer = boardLayerCellMap(lx, rx, ty, by, boxLen, boxWid, boardPos)
	board.pathLayer = createPathsLL(lx, rx, ty, by, boxLen, boxWid, boardPos)
	board.renderBoardLayer()
}
