package game

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

func (a *arena) render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	a.board.render()
	a.renderBottomSection()
	termbox.Flush()
}

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

func (b *ludoBoard) renderPawns() {
	for _, p := range b.players {
		for _, pawn := range p.pawns {
			c := pawn["curNode"].cell
			setCell(c.x, c.y, ' ', termbox.ColorDefault, c.bg)
		}
	}
}

func (a *arena) renderBottomSection() {
	x, y := 10, 22
	renderWhoseTurn(a.players[a.curTurn].Color, x, y)
	renderText(x+2, y, "'s turn")
	renderText(x+20, y, "dice: "+strconv.Itoa(a.dice.value))
}

func renderWhoseTurn(bg termbox.Attribute, x, y int) {
	setCell(x, y, ' ', termbox.ColorDefault, bg)
}

func renderText(x, y int, text string) {
	for i := range text {
		termbox.SetChar(x, y, rune(text[i]))
		x++
	}
}
