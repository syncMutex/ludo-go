package common

const (
	UNKNOWN_ERR = iota - 2
	KNOWN_ERR   = iota - 1
	CONN_RES    = iota
	JOINED_PLAYERS
	PLAYER_COLOR
	START_GAME
)

type Res struct {
	Ok  bool
	Msg string
}
