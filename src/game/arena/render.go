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
	pName := a.Players[a.CurTurn].Name
	renderWhoseTurn(a.Players[a.CurTurn].Color, x, y)
	tbu.RenderString(x+3, y, pName, termbox.ColorDefault)
	tbu.RenderString(x+3+len(pName), y, "'s turn", termbox.ColorDefault)
	tbu.RenderString(x+3+len(pName)+12, y, "Dice: "+strconv.Itoa(a.Dice.Value), termbox.ColorDefault)
}

func renderWhoseTurn(bg termbox.Attribute, x, y int) {
	tbu.SetCell(x, y, ' ', termbox.ColorDefault, bg)
}

func (a *Arena) LeaderBoard() []termbox.Attribute {
	lb := make([]termbox.Attribute, a.ParticipantsCount)
	for _, p := range a.Board.Players {
		if p.IsParticipant() {
			lb[p.WinningPos-1] = p.Color
		}
	}
	return lb
}

func (a *Arena) getPlayerName(color termbox.Attribute) string {
	for _, p := range a.Players {
		if p.Color == color {
			return p.Name
		}
	}
	return ""
}

func (a *Arena) RenderGameOver(lb []termbox.Attribute) {
	tbu.Clear()

	x, y := 5, 5

	tbu.RenderString(x+5, y-3, "Game Over!", termbox.ColorGreen)

	for i, p := range lb {
		x = 5
		tbu.RenderString(x, y, strconv.Itoa(i+1), termbox.ColorDefault)
		tbu.SetBg(x+3, y, p)
		tbu.RenderString(x+6, y, a.getPlayerName(p), termbox.ColorDefault)
		y += 2
	}

	tbu.RenderString(x+2, y+2, "press esc to exit.", termbox.ColorGreen)

	termbox.Flush()
}
