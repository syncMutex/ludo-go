package network

import (
	"bufio"
	"encoding/gob"
	"ludo/src/common"
	"net"
	"sync"
)

type GobHandler struct {
	enc    *gob.Encoder
	dec    *gob.Decoder
	reader *bufio.Reader
	Conn   net.Conn
}

type InstrucLoopHandler struct {
	NetChan               chan int
	receiverFunc          func() (int, error)
	IsRunning             bool
	processNextInstucChan chan bool
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
		return common.UNKNOWN_ERR, err
	}
	return instruc, err
}

func (h *GobHandler) SendInstruc(instruc int, data interface{}, wg ...*sync.WaitGroup) error {
	err := h.Encode(instruc)

	if len(wg) != 0 && wg[0] != nil {
		wg[0].Done()
		return nil
	}
	return err
}

func (h *GobHandler) SendResponse(instruc int, data interface{}, wg ...*sync.WaitGroup) error {
	h.Encode(instruc)
	err := h.Encode(data)

	if len(wg) != 0 && wg[0] != nil {
		wg[0].Done()
		return nil
	}
	return err
}

func (h *GobHandler) GetRes() (res common.Res) {
	h.Decode(&res)
	return
}

func NewGobHandler(conn net.Conn) *GobHandler {
	return &GobHandler{gob.NewEncoder(conn), gob.NewDecoder(conn), bufio.NewReader(conn), conn}
}

func NewInstrucLoopHandler(receiverFunc func() (int, error)) *InstrucLoopHandler {
	return &InstrucLoopHandler{
		NetChan:               make(chan int),
		receiverFunc:          receiverFunc,
		IsRunning:             false,
		processNextInstucChan: make(chan bool),
	}
}

func (n *InstrucLoopHandler) Continue(isContinue bool) {
	n.processNextInstucChan <- isContinue
}

func (n *InstrucLoopHandler) RunLoop() {
	for {
		n.IsRunning = true
		instruc, _ := n.receiverFunc()
		n.NetChan <- instruc
		n.IsRunning = false
		if yes := <-n.processNextInstucChan; !yes {
			return
		}
	}
}
