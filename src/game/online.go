package game

import (
	"ludo/src/common"
	"ludo/src/game/arena"
	"ludo/src/keyboard"
	"ludo/src/network"
	tbu "ludo/src/termbox-utils"
	"time"

	"github.com/nsf/termbox-go"
)

func getBoardState(gh *network.GobHandler) (brdSt common.BoardState) {
	gh.Decode(&brdSt)
	return
}

func onlineGameLoop(
	a *arena.Arena,
	kChan *keyboard.KeyboardProps,
	gh *network.GobHandler,
	nh *network.InstrucLoopHandler,
	playerData common.PlayerData,
) int {

	nh.Continue(true)

	kChan.Pause()

	a.KChan = kChan
	a.BlinkCh = make(chan bool)

	a.SetupBoard()

	curTurnFunc := func() {
		if playerData.Color == a.CurPlayer().Color {
			a.StartBlinkCurPawn()
			kChan.Resume()
		}
	}

	for {
		select {
		case instruc := <-nh.NetChan:
			switch instruc {
			case common.BOARD_STATE:
				brdSt := getBoardState(gh)
				a.Dice.Value = brdSt.DiceValue
				a.SetCurPlayerAndPawn(brdSt.CurTurn, 0)
				curTurnFunc()
				a.Render()
			case common.KNOWN_ERR:
				tbu.Clear()
				tbu.RenderText(tbu.Text{X: 5, Y: 3, Text: gh.GetRes().Msg, Fg: termbox.ColorRed})
				termbox.Flush()
				time.Sleep(time.Second * 3)
				return -1
			}
			nh.Continue(true)
		case ev := <-kChan.EvChan:
			if ev.Key == termbox.KeyEsc {
				return 0
			}
		}
	}
}

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

	nh := network.NewInstrucLoopHandler(gh.ReceiveInstruc)
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	go nh.RunLoop()
	go keyboard.ListenToKeyboard(&kChan)

	defer func() {
		kChan.Stop()
		gh.Conn.Close()
	}()

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
			case common.START_GAME:
				var a arena.Arena
				gh.Decode(&a)
				return onlineGameLoop(&a, &kChan, gh, nh, playerData)
			case common.KNOWN_ERR:
				tbu.Clear()
				tbu.RenderText(tbu.Text{X: 5, Y: 3, Text: gh.GetRes().Msg, Fg: termbox.ColorRed})
				termbox.Flush()
				time.Sleep(time.Second * 3)
				return -1
			}
			nh.Continue(true)
		case ev := <-kChan.EvChan:
			if ev.Key == termbox.KeyEsc {
				return 0
			}
		}
	}
}
