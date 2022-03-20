package network

import (
	"sync"
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
