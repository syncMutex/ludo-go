package arena

import (
	"strconv"

	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

func (a *Arena) Render() {
	tbu.Clear()
	a.Board.Render()
	a.RenderBottomSection()
	termbox.Flush()
}

func (a *Arena) RenderBottomSection() {
	x, y := 10, 22
	renderWhoseTurn(a.Players[a.CurTurn].Color, x, y)
	tbu.RenderString(x+2, y, "'s turn", termbox.ColorDefault)
	tbu.RenderString(x+20, y, "Dice: "+strconv.Itoa(a.Dice.Value), termbox.ColorDefault)
}

func renderWhoseTurn(bg termbox.Attribute, x, y int) {
	tbu.SetCell(x, y, ' ', termbox.ColorDefault, bg)
}

func (a *Arena) RenderGameOver() {
	tbu.Clear()
	leaderBoard := make([]termbox.Attribute, a.ParticipantsCount)
	for _, p := range a.Board.Players {
		if p.IsParticipant() {
			leaderBoard[p.WinningPos-1] = p.Color
		}
	}

	x, y := 5, 5

	tbu.RenderString(x+5, y-3, "Game Over!", termbox.ColorGreen)

	for i, p := range leaderBoard {
		x = 5
		tbu.RenderString(x, y, strconv.Itoa(i+1), termbox.ColorDefault)
		tbu.SetBg(x+3, y, p)
		y += 2
	}

	tbu.RenderString(x+2, y+2, "press esc to exit.", termbox.ColorGreen)

	termbox.Flush()
}
