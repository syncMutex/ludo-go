package common

import "github.com/nsf/termbox-go"

const (
	UNKNOWN_ERR = iota - 2
	KNOWN_ERR   = iota - 1
	CONN_RES    = iota
	JOINED_PLAYERS
	PLAYER_COLOR
	START_GAME

	BOARD_STATE
)

type Res struct {
	Ok  bool
	Msg string
}

type BoardState struct {
	CurTurn   termbox.Attribute
	DiceValue int
}
