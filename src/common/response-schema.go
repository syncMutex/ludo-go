package common

const (
	ERROR    = iota - 1
	CONN_RES = iota
	JOINED_PLAYERS
	PLAYER_DATA
	JOIN_GAME
)

type Res struct {
	Ok  bool
	Msg string
}
