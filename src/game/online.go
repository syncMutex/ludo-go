package game

import (
	"ludo/src/common"
	"ludo/src/keyboard"
	"ludo/src/network"
	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

func renderJoinedPlayers(players []common.PlayerData) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	x, y := 2, 5
	tbu.RenderString(x, y, "Lobby", termbox.ColorBlue)
	y += 2
	for _, p := range players {
		tbu.SetCell(x, y, ' ', termbox.ColorDefault, p.Color)
		tbu.RenderString(x+4, y, p.Name, termbox.ColorDefault)
		y += 2
	}
	termbox.Flush()
}

func StartGameOnline() int {
	var players []common.PlayerData
	var playerData common.PlayerData

	playerData.Name = tbu.PromptText(3, 3, termbox.ColorDefault, "Your Name: ", 15)

	gh, err := network.Join()

	if err != nil {
		termbox.Close()
		return -1
	}
	defer gh.Conn.Close()

	nh := network.NewInstrucLoopHandler(gh.ReceiveInstruc)
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	go nh.RunLoop()
	go keyboard.ListenToKeyboard(&kChan)

	for {
		select {
		case instruc := <-nh.NetChan:
			switch instruc {
			case common.CONN_RES:
				if gh.GetRes().Ok {
					gh.Encode(playerData.Name)
				}
			case common.PLAYER_COLOR:
				gh.Decode(&playerData.Color)
			case common.JOINED_PLAYERS:
				gh.Decode(&players)
				renderJoinedPlayers(players)
			}
			nh.Resume()
		case ev := <-kChan.EvChan:
			if ev.Key == termbox.KeyEsc {
				kChan.Stop()
				return 0
			}
		}
	}
}
