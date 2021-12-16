package main

import (
	"ludo/src/menu"

	"github.com/nsf/termbox-go"
)

func initTermboxStuff() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
}

func uninitTermboxStuff() {
	termbox.Close()
}

func main() {
	defer uninitTermboxStuff()
	initTermboxStuff()
	menu.StartMainMenu()
}
