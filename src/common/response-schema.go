package common

const (
	ERROR    = iota - 1
	CONN_RES = iota
	TEST_RES
)

type ConnRes struct {
	Ok  bool
	Msg string
}

type TestRes struct {
	Msg string
}
