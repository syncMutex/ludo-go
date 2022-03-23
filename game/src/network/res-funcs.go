package network

import (
	"sync"

	"github.com/nsf/termbox-go"
)

func (s *Server) broadcastResponse(instruc int, data interface{}) {
	var wg sync.WaitGroup

	for _, p := range s.players {
		if p.gh != nil {
			wg.Add(1)
			go p.gh.SendResponse(instruc, data, &wg)
		}
	}
	wg.Wait()
}

func inArray[T termbox.Attribute](val T, arr ...T) bool {
	for _, e := range arr {
		if e == val {
			return true
		}
	}
	return false
}

func (s *Server) broadcastResponseExcept(instruc int, data interface{}, except ...termbox.Attribute) {
	var wg sync.WaitGroup

	for _, p := range s.players {
		if p.gh != nil && !inArray(p.Color, except...) {
			wg.Add(1)
			go p.gh.SendResponse(instruc, data, &wg)
		}
	}
	wg.Wait()
}

func (s *Server) broadcastInstruc(instruc int) {
	var wg sync.WaitGroup

	for _, p := range s.players {
		if p.gh != nil {
			wg.Add(1)
			go p.gh.SendInstruc(instruc, &wg)
		}
	}
	wg.Wait()
}
