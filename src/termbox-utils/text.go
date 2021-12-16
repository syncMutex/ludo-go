package tbutils

import "github.com/nsf/termbox-go"

type Text struct {
	X, Y          int
	Text          string
	Fg, Bg        termbox.Attribute
	InlinePadding int
}

func RenderText(text Text) {
	for text.InlinePadding != 0 {
		text.Text = " " + text.Text + " "
		text.InlinePadding--
	}
	for i := range text.Text {
		termbox.SetCell(text.X, text.Y, rune(text.Text[i]), text.Fg, text.Bg)
		text.X++
	}
}
