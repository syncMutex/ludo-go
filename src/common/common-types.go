package common

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

type PlayerData struct {
	Type  string
	Color termbox.Attribute
	Name  string
}

type Dice struct {
	Value int
}

func (d *Dice) Roll() {
	d.Value = rand.Intn(6) + 1
}
