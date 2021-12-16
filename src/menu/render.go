package menu

import (
	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

func renderMenu() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	y := menuPos.y
	for idx, opt := range menuOps {
		if curOpt == idx {
			tbu.RenderText(tbu.Text{X: menuPos.x, Y: y, Text: opt, Bg: termbox.ColorYellow, Fg: termbox.ColorBlack, InlinePadding: 3})
		} else {
			tbu.RenderText(tbu.Text{X: menuPos.x, Y: y, Text: opt})
		}
		y++
	}
	termbox.Flush()
}
