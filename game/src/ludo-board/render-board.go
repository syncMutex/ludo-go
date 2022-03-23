package board

import (
	tbu "ludo/src/termbox-utils"
	"strconv"

	"github.com/nsf/termbox-go"
)

func (b *LudoBoard) RenderBoardLayer() {
	for _, p := range b.boardLayer {
		termbox.SetCell(p.X, p.Y, p.Ch, p.Fg, p.Bg)
	}
}

func (b *LudoBoard) RenderPathLayer() {
	temp := b.pathLayer.LL.Head
	tbu.SetCell(temp.Cell.X, temp.Cell.Y, temp.Cell.Ch, temp.Cell.Fg, temp.Cell.Bg)
	temp = temp.Next["common"]

	for temp != b.pathLayer.LL.Head {
		tbu.SetCell(temp.Cell.X, temp.Cell.Y, temp.Cell.Ch, temp.Cell.Fg, temp.Cell.Bg)
		if temp.Next["toDest"] != nil {
			temp2 := temp.Next["toDest"]
			for temp2 != nil {
				tbu.SetCell(temp2.Cell.X, temp2.Cell.Y, temp2.Cell.Ch, temp2.Cell.Fg, temp2.Cell.Bg)
				temp2 = temp2.Next["toDest"]
			}
		}
		temp = temp.Next["common"]
	}
}

func (b *LudoBoard) RenderPawns() {
	for _, p := range b.Players {
		if p.PlayerType == "-" {
			continue
		}
		pkeys := make(map[string]int)
		for _, pawn := range p.Pawns {
			c := pawn["curNode"].Cell
			tbu.SetCell(c.X, c.Y, ' ', termbox.ColorDefault, p.Color)

			if _, has := pkeys[c.MapKey()]; has {
				tbu.RenderString(c.X, c.Y, strconv.Itoa(pkeys[c.MapKey()]+1), termbox.ColorBlack)
			}
			pkeys[c.MapKey()]++
		}
	}
}

func (b *LudoBoard) Render() {
	b.RenderBoardLayer()
	b.RenderPathLayer()
	b.RenderPawns()
}
