package menu

import (
	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

func (m *menuPagesType) renderMenu() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	curMenu := m.menus[m.curIdx]

	y := m.displayPos.y

	var fg, bg termbox.Attribute

	for oidx, op := range curMenu.opts {
		fg, bg = termbox.ColorDefault, termbox.ColorDefault
		if curMenu.curIdx == oidx {
			bg = termbox.ColorYellow
			fg = termbox.ColorBlack
		}
		tbu.RenderText(tbu.Text{Text: op.label, X: m.displayPos.x, Y: y, Fg: fg, Bg: bg, InlinePadding: 3})
		y += 2
	}

	termbox.Flush()
}
