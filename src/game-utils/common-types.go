package gameUtils

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

type PlayerData struct {
	Type  string
	Color termbox.Attribute
}

type Dice struct {
	Value int
}

func (d *Dice) Roll() {
	d.Value = rand.Intn(6) + 1
}
