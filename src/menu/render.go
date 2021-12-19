package menu

import (
	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

var (
	paddingMag = 3
)

func renderSubOpts(op opt, x, y, oidx, menuIdx int) {
	var fg, bg termbox.Attribute

	if menuIdx == oidx {
		bg, fg = termbox.ColorYellow, termbox.ColorBlack
	}
	tbu.RenderText(tbu.Text{Text: op.label, X: x, Y: y, Fg: fg, Bg: bg, InlinePadding: paddingMag})

	x = x + len(op.label) + paddingMag + 5

	switch t := op.subOpts[op.curIdx].(type) {
	case string:
		if menuIdx == oidx {
			fg, bg = termbox.ColorGreen, termbox.ColorDefault
		} else {
			fg, bg = termbox.ColorDefault, termbox.ColorDefault
		}
		tbu.RenderText(tbu.Text{Text: t, X: x, Y: y, Fg: fg, Bg: bg})
	case termbox.Attribute:
		tbu.RenderText(tbu.Text{Text: "", X: x, Y: y, Fg: fg, Bg: t, InlinePadding: 1})
	}

}

func (m *menuPagesType) renderMenu() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	curMenu := m.menus[m.curIdx]

	y := m.displayPos.y

	var fg, bg termbox.Attribute

	for oidx, op := range curMenu.opts {
		fg, bg = termbox.ColorDefault, termbox.ColorDefault

		if op.subOpts != nil {
			renderSubOpts(op, m.displayPos.x, y, oidx, curMenu.curIdx)
		} else {
			if curMenu.curIdx == oidx {
				bg, fg = termbox.ColorYellow, termbox.ColorBlack
			}
			tbu.RenderText(tbu.Text{Text: op.label, X: m.displayPos.x, Y: y, Fg: fg, Bg: bg, InlinePadding: paddingMag})
		}
		y += 2
	}

	termbox.Flush()
}
