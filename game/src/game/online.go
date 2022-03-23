package game

import (
	"fmt"
	"ludo/src/common"
	"ludo/src/game/arena"
	"ludo/src/keyboard"
	"ludo/src/network"
	"ludo/src/network/schema"
	tbu "ludo/src/termbox-utils"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	animDuration = time.Millisecond * 0
)

func handleKeyboardOnline(a *arena.Arena, k keyboard.KeyboardEvent, gh *network.GobHandler) bool {
	if a.IsGameOver() {
		return k.Key == termbox.KeyEsc
	}

	a.StopBlinkCurPawn()
	a.RepaintCurPawn()
	switch k.Key {
	case termbox.KeyArrowRight:
		a.StartBlinkCurPawn()
		a.SetNextCurPawnAndValidate(1)
	case termbox.KeyArrowLeft:
		a.StartBlinkCurPawn()
		a.SetNextCurPawnAndValidate(-1)
	case termbox.KeyEnter:
		fallthrough
	case termbox.KeySpace:
		gh.SendResponse(schema.MOVE, a.Board.CurPawnIdx)
		a.MakeMove(animDuration, DO_RENDER)
		a.Render()
	case termbox.KeyEsc:
		return true
	}
	a.Render()
	return false
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
			a.Board.SetCurPawn(3)
			a.SetNextCurPawnAndValidate(1)
			kChan.Resume()
		}
	}

	for {
		select {
		case instruc := <-nh.NetChan:
			switch instruc {
			case schema.BOARD_STATE:
				brdSt := network.DecodeData[schema.BoardState](gh)
				a.Dice.Value = brdSt.DiceValue
				a.SetCurPlayerAndPawn(brdSt.CurTurn, 0)
				a.Render()
				curTurnFunc()
			case schema.LOOPED:
				brdSt := network.DecodeData[schema.BoardState](gh)
				a.Dice.Value = brdSt.DiceValue
				a.SetCurPlayerAndPawn(brdSt.CurTurn, 0)
				a.Render()
			case schema.MOVE_BY:
				movedBy := network.DecodeData[schema.MoveBy](gh)
				a.SetCurPlayerAndPawn(movedBy.Color, movedBy.PawnIdx)
				a.MakeMove(animDuration, DO_RENDER)
				a.Render()
			case schema.KNOWN_ERR:
				tbu.Clear()
				tbu.RenderText(tbu.Text{X: 5, Y: 3, Text: network.DecodeData[schema.Res](gh).Msg, Fg: termbox.ColorRed})
				termbox.Flush()
				time.Sleep(time.Second * 3)
				return -1
			case schema.GAME_OVER:
				lb := network.DecodeData[[]termbox.Attribute](gh)
				a.RenderGameOver(lb)
				kChan.Resume()
				for {
					ev := <-kChan.EvChan
					kChan.Pause()
					if ev.Key == termbox.KeyEsc {
						return 0
					}
					kChan.Resume()
				}
			}
			nh.Continue(true)
		case ev := <-kChan.EvChan:
			kChan.Pause()
			if stop := handleKeyboardOnline(a, ev, gh); stop {
				return 0
			}
			if ev.Key != termbox.KeySpace && ev.Key != termbox.KeyEnter {
				kChan.Resume()
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

func renderErr(errText string) {
	tbu.Clear()
	tbu.RenderText(tbu.Text{X: 5, Y: 3, Text: errText, Fg: termbox.ColorRed})
	termbox.Flush()
	time.Sleep(time.Second * 3)
}

func StartGameOnline(url string) int {
	var playerData common.PlayerData

	playerData.Name = tbu.PromptText(3, 3, termbox.ColorDefault, "Your Name: ", 15)

	if url == "" {
		tbu.Clear()
		termbox.Close()
		fmt.Print("paste the game url: ")
		fmt.Scanln(&url)
		termbox.Init()
	}

	gh, err := network.Join(url)

	if err != nil {
		renderErr("unable to connect. Maybe check the url.")
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
			case schema.CONN_RES:
				if network.DecodeData[schema.Res](gh).Ok {
					gh.Encode(playerData.Name)
				}
			case schema.PLAYER_COLOR:
				gh.Decode(&playerData.Color)
			case schema.JOINED_PLAYERS:
				players := network.DecodeData[[]common.PlayerData](gh)
				renderJoinedPlayers(players)
			case schema.START_GAME:
				a := network.DecodeData[arena.Arena](gh)
				return onlineGameLoop(&a, &kChan, gh, nh, playerData)
			case schema.KNOWN_ERR:
				renderErr(network.DecodeData[schema.Res](gh).Msg)
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
