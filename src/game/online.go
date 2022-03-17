package game

import (
	"ludo/src/common"
	"ludo/src/keyboard"
	"ludo/src/network"
	tbu "ludo/src/termbox-utils"
	"time"

	"github.com/nsf/termbox-go"
)

func renderJoinedPlayers(players []common.PlayerData) {
	tbu.Clear()
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
			case common.KNOWN_ERR:
				tbu.Clear()
				tbu.RenderText(tbu.Text{X: 5, Y: 3, Text: gh.GetRes().Msg, Fg: termbox.ColorRed})
				termbox.Flush()
				time.Sleep(time.Second * 3)
				nh.Continue(false)
				kChan.Stop()
				return -1
			}
			nh.Continue(true)
		case ev := <-kChan.EvChan:
			if ev.Key == termbox.KeyEsc {
				nh.Continue(false)
				kChan.Stop()
				return 0
			}
		}
	}
}
