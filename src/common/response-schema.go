package common

const (
	ERROR    = iota - 1
	CONN_RES = iota
	JOINED_PLAYERS
	PLAYER_COLOR
)

type Res struct {
	Ok  bool
	Msg string
}
