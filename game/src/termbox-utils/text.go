package tbutils

import (
	"ludo/src/keyboard"

	"github.com/nsf/termbox-go"
)

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

func RenderString(x, y int, text string, textColor termbox.Attribute) {
	for i := range text {
		termbox.SetChar(x, y, rune(text[i]))
		termbox.SetFg(x, y, textColor)
		x++
	}
}

func PromptText(x, y int, fg termbox.Attribute, label string, maxLen int) (inp string) {
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	go keyboard.ListenToKeyboard(&kChan)
	Clear()
	termbox.SetCursor(x+len(label)+len(inp), y)
	RenderText(Text{x, y, label, fg, termbox.ColorDefault, 0})
	termbox.Flush()

	for {
		ev := <-kChan.EvChan
		kChan.Pause()
		if (len(inp) < maxLen) && (ev.Ch >= 'a' && ev.Ch <= 'z') ||
			(ev.Ch >= 'A' && ev.Ch <= 'Z') ||
			(ev.Ch >= '0' && ev.Ch <= '9') ||
			(ev.Ch == '_') ||
			(ev.Ch == '-') {
			inp += string(ev.Ch)
		} else if ev.Key == termbox.KeyBackspace && len(inp) != 0 {
			inp = inp[0 : len(inp)-1]
		} else if ev.Key == termbox.KeyEnter {
			break
		}

		Clear()
		RenderText(Text{x, y, label, termbox.ColorDefault, termbox.ColorDefault, 0})
		RenderText(Text{x + len(label), y, inp, fg, termbox.ColorDefault, 0})
		termbox.Flush()
		kChan.Resume()
		termbox.SetCursor(x+len(label)+len(inp), y)
	}
	kChan.Stop()
	termbox.HideCursor()
	Clear()
	termbox.Flush()
	return
}
