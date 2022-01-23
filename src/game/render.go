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

func setBg(x, y int, bg termbox.Attribute) {
	termbox.SetBg(x, y, bg)
	termbox.SetBg(x+1, y, bg)
}

func (b *ludoBoard) renderBoardLayer() {
	for _, p := range b.boardLayer {
		termbox.SetCell(p.x, p.y, p.ch, p.fg, p.bg)
	}
}

func (b *ludoBoard) renderPathLayer() {
	temp := b.pathLayer.ll.head
	setCell(temp.cell.x, temp.cell.y, temp.cell.ch, temp.cell.fg, temp.cell.bg)
	temp = temp.next["common"]

	for temp != b.pathLayer.ll.head {
		setCell(temp.cell.x, temp.cell.y, temp.cell.ch, temp.cell.fg, temp.cell.bg)
		if temp.next["toDest"] != nil {
			temp2 := temp.next["toDest"]
			for temp2 != nil {
				setCell(temp2.cell.x, temp2.cell.y, temp2.cell.ch, temp2.cell.fg, temp2.cell.bg)
				temp2 = temp2.next["toDest"]
			}
		}
		temp = temp.next["common"]
	}
}

func (b *ludoBoard) renderPawns() {
	for _, p := range b.players {
		if p.playerType == "-" {
			continue
		}
		pkeys := make(map[string]int)
		for _, pawn := range p.pawns {
			c := pawn["curNode"].cell
			setCell(c.x, c.y, ' ', termbox.ColorDefault, p.color)

			if _, has := pkeys[c.mapKey()]; has {
				renderText(c.x, c.y, strconv.Itoa(pkeys[c.mapKey()]+1), termbox.ColorBlack)
			}
			pkeys[c.mapKey()]++
		}
	}
}

func (a *arena) renderBottomSection() {
	x, y := 10, 22
	renderWhoseTurn(a.players[a.curTurn].Color, x, y)
	renderText(x+2, y, "'s turn", termbox.ColorDefault)
	renderText(x+20, y, "dice: "+strconv.Itoa(a.dice.value), termbox.ColorDefault)
}

func renderWhoseTurn(bg termbox.Attribute, x, y int) {
	setCell(x, y, ' ', termbox.ColorDefault, bg)
}

func renderText(x, y int, text string, textColor termbox.Attribute) {
	for i := range text {
		termbox.SetChar(x, y, rune(text[i]))
		termbox.SetFg(x, y, textColor)
		x++
	}
}

func (a *arena) renderGameOver() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	leaderBoard := make([]termbox.Attribute, a.participantsCount)
	for _, p := range a.board.players {
		if p.isParticipant() {
			leaderBoard[p.winningPos-1] = p.color
		}
	}

	x, y := 5, 5

	renderText(x+5, y-3, "Game Over!", termbox.ColorGreen)

	for i, p := range leaderBoard {
		x = 5
		renderText(x, y, strconv.Itoa(i+1), termbox.ColorDefault)
		setBg(x+3, y, p)
		y += 2
	}

	renderText(x+2, y+2, "press esc to exit.", termbox.ColorGreen)

	termbox.Flush()
}
