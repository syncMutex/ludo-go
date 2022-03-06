package common

import (
	"bufio"
	"encoding/gob"
	"math/rand"
	"net"

	"github.com/nsf/termbox-go"
)

type PlayerData struct {
	Type  string
	Color termbox.Attribute
}

type Dice struct {
	Value int
}

type GobHandler struct {
	enc    *gob.Encoder
	dec    *gob.Decoder
	reader *bufio.Reader
	Conn   net.Conn
}

func (h *GobHandler) Encode(i interface{}) error {
	return h.enc.Encode(i)
}

func (h *GobHandler) Decode(i interface{}) error {
	return h.dec.Decode(i)
}

func (h *GobHandler) ReceiveInstruc() (int, error) {
	var instruc int

	err := h.Decode(&instruc)

	if err != nil {
		return ERROR, err
	}
	return instruc, err
}

func (h *GobHandler) SendResponse(instruc int, data interface{}) error {
	h.Encode(instruc)
	return h.Encode(data)
}

func NewGobHandler(conn net.Conn) *GobHandler {
	return &GobHandler{gob.NewEncoder(conn), gob.NewDecoder(conn), bufio.NewReader(conn), conn}
}

func (d *Dice) Roll() {
	d.Value = rand.Intn(6) + 1
}
