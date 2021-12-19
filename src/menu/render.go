package menu

import (
	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

var (
	paddingMag = 3
)

func (m *menuPagesType) renderMenu() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	curMenu := m.menus[m.curIdx]

	y := m.displayPos.y

	var fg, bg termbox.Attribute

	for oidx, op := range curMenu.opts {
		fg, bg = termbox.ColorDefault, termbox.ColorDefault

		if op.subOpts != nil {
			tbu.RenderText(tbu.Text{Text: op.label, X: m.displayPos.x, Y: y, Fg: fg, Bg: bg, InlinePadding: paddingMag})

			x := m.displayPos.x + len(op.label) + paddingMag + 2

			for sidx, subOpt := range op.subOpts {
				if sidx == op.curIdx {
					fg, bg = termbox.ColorBlack, termbox.ColorYellow
					if curMenu.curIdx != oidx {
						fg, bg = termbox.ColorBlack, termbox.ColorCyan
					}
				} else {
					fg, bg = termbox.ColorDefault, termbox.ColorDefault
				}
				tbu.RenderText(tbu.Text{Text: subOpt, X: x, Y: y, Fg: fg, Bg: bg, InlinePadding: 1})
				x += 3
			}
		} else {
			if curMenu.curIdx == oidx {
				bg = termbox.ColorYellow
				fg = termbox.ColorBlack
			}
			tbu.RenderText(tbu.Text{Text: op.label, X: m.displayPos.x, Y: y, Fg: fg, Bg: bg, InlinePadding: paddingMag})
		}
		y += 2
	}

	termbox.Flush()
}
