package tbutils

import "github.com/nsf/termbox-go"

func SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
	termbox.SetCell(x+1, y, ch, fg, bg)
}

func SetBg(x, y int, bg termbox.Attribute) {
	termbox.SetBg(x, y, bg)
	termbox.SetBg(x+1, y, bg)
}

func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}
