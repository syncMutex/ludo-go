package game

import (
	"github.com/nsf/termbox-go"
)

func setCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
	termbox.SetCell(x+1, y, ch, fg, bg)
}

func (b *ludoBoard) renderBoardLayer() {
	for _, p := range b.boardLayer {
		termbox.SetCell(p.x, p.y, p.ch, p.fg, p.bg)
	}
}

func (b *ludoBoard) renderPathLayer() {
	temp := b.pathLayer.ll.head

	for temp != nil {
		setCell(temp.cell.x, temp.cell.y, temp.cell.ch, temp.cell.fg, temp.cell.bg)
		if temp.next["toHome"] != nil {
			temp2 := temp.next["toHome"]
			for temp2 != nil {
				setCell(temp2.cell.x, temp2.cell.y, temp2.cell.ch, temp2.cell.fg, temp2.cell.bg)
				temp2 = temp2.next["toHome"]
			}
		}
		temp = temp.next["common"]
	}
}

func (b *ludoBoard) renderHome() {
	for _, p := range b.players {
		for _, pawn := range p.pawns {
			c := pawn["curNode"].cell
			setCell(c.x, c.y, ' ', termbox.ColorDefault, c.bg)
		}
	}
}
