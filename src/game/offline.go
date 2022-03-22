package game

import (
	"ludo/src/common"
	"ludo/src/game/arena"
	"ludo/src/keyboard"
	board "ludo/src/ludo-board"
	"time"

	"github.com/nsf/termbox-go"
)

const DO_RENDER = true

func handleKeyboardOffline(a *arena.Arena, k keyboard.KeyboardEvent) bool {
	if a.IsGameOver() {
		return k.Key == termbox.KeyEsc
	}

	a.StopBlinkCurPawn()
	a.RepaintCurPawn()
	switch k.Key {
	case termbox.KeyArrowRight:
		a.SetNextCurPawnAndValidate(1)
	case termbox.KeyArrowLeft:
		a.SetNextCurPawnAndValidate(-1)
	case termbox.KeyEnter:
		fallthrough
	case termbox.KeySpace:
		hasDestroyed, hasReachedDest := a.MakeMove(time.Millisecond*0, true)
		a.Dice.Roll()
		a.Render()
		if !hasDestroyed && !hasReachedDest {
			a.ChangePlayerTurnAndValidate(DO_RENDER, nil)
		} else if hasReachedDest {
			if a.CurPlayer().IsAllPawnsAtDest() {
				a.SetCurPlayerWin()
				if a.IsGameOver() {
					a.ChangePlayerTurn()
					a.SetCurPlayerWin()
					a.RenderGameOver(a.LeaderBoard())
					return false
				}
			}
			if ok := a.SetNextCurPawnAndValidate(1); !ok {
				a.ChangePlayerTurnAndValidate(DO_RENDER, nil)
			}
		}
	case termbox.KeyEsc:
		return true
	}
	a.Render()
	a.StartBlinkCurPawn()
	return false
}

func botFuncOffline(a *arena.Arena) {
	handleKeyboardOffline(a, keyboard.KeyboardEvent{Key: termbox.KeyEnter, Ch: ' '})
}

func runGameLoopOffline(a *arena.Arena) {
	kChan := a.KChan

	go keyboard.ListenToKeyboard(kChan)
	a.ChangePlayerTurn(1)
	a.ChangePlayerTurnAndValidate(DO_RENDER, nil)
	a.Board.SetCurPawn(0)
	common.SetRandSeed()
	a.Dice.Roll()

	a.Render()
	a.StartBlinkCurPawn()

	for a.CurPlayer().IsBot() {
		a.PlayBot(botFuncOffline)
		if a.IsGameOver() {
			break
		}
	}
mainloop:
	for {
		ev := <-kChan.EvChan
		kChan.Pause()
		if stop := handleKeyboardOffline(a, ev); stop {
			kChan.Stop()
			break mainloop
		}
		for a.CurPlayer().IsBot() {
			a.PlayBot(botFuncOffline)
			if a.IsGameOver() {
				break
			}
		}
		kChan.Resume()
	}
}

func StartGameOffline(players []common.PlayerData) {

	gameDice := common.Dice{}
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	a := arena.Arena{
		Board:          board.LudoBoard{},
		Players:        players,
		BlinkCh:        make(chan bool),
		NextWinningPos: 0,
		Dice:           gameDice,
		Bots:           make(map[int][4]int),
		KChan:          &kChan,
	}
	a.SetupBoard()
	a.BotsInit()
	runGameLoopOffline(&a)
}
