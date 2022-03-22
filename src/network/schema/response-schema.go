package schema

import (
	"ludo/src/common"
	"ludo/src/game/arena"

	"github.com/nsf/termbox-go"
)

const (
	UNKNOWN_ERR = iota - 2
	KNOWN_ERR   = iota - 1
	CONN_RES    = iota
	JOINED_PLAYERS
	PLAYER_COLOR
	START_GAME
	GAME_OVER

	BOARD_STATE
	MOVE
	MOVE_BY
	LOOPED
)

type Res struct {
	Ok  bool
	Msg string
}

type BoardState struct {
	CurTurn   termbox.Attribute
	DiceValue int
}

type MoveBy struct {
	Color   termbox.Attribute
	PawnIdx int
}

type NetworkSchemas interface {
	Res | BoardState | MoveBy | int | string | []common.PlayerData | arena.Arena | []termbox.Attribute
}
