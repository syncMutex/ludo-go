package game

import (
	"ludo/src/common"
	"ludo/src/keyboard"
	board "ludo/src/ludo-board"
	"ludo/src/network"
	tbu "ludo/src/termbox-utils"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Arena struct {
	players           []common.PlayerData
	Dice              common.Dice
	board             board.LudoBoard
	curTurn           int
	blinkCh           chan bool
	isBlinkChOpen     bool
	nextWinningPos    int
	participantsCount int
	bots              map[int][4]int
	kChan             keyboard.KeyboardProps
}

func (a *Arena) changePlayerTurn(idx ...int) bool {
	if len(idx) == 1 {
		a.curTurn = idx[0]
	} else {
		a.curTurn++
		if a.curTurn >= len(a.players) {
			a.curTurn = 0
		}
		for !a.curPlayer().IsParticipant() {
			a.curTurn++
			if a.curTurn >= len(a.players) {
				a.curTurn = 0
			}
		}
	}

	if a.curPlayer().IsAllPawnsAtDest() {
		a.changePlayerTurn()
	}

	a.board.SetCurPawn(0)
	if a.curPawn().IsAtDest() || !a.curPawn().HasNPathsAhead(a.Dice.Value) {
		return a.setNextCurPawnAndValidate(1)
	}

	return true
}

func (a *Arena) setNextCurPawnAndValidate(mag int) bool {
	temp := a.board.CurPawnIdx
	for a.board.SetNextCurPawn(a.curTurn, mag); a.curPawn().IsAtDest(); {
		a.board.SetNextCurPawn(a.curTurn, mag)
		if a.board.CurPawnIdx == temp {
			break
		}
	}

	temp = a.board.CurPawnIdx
	for !a.curPawn().HasNPathsAhead(a.Dice.Value) {
		a.board.SetNextCurPawn(a.curTurn, mag)
		if a.board.CurPawnIdx == temp {
			return false
		}
	}

	return true
}

func (a *Arena) setCurPlayerWin() {
	a.board.Players[a.curTurn].WinningPos = a.nextWinningPos + 1
	a.nextWinningPos++
}

func (a *Arena) changePlayerTurnAndValidate() {
	ok := a.changePlayerTurn()
	a.render()
	for !ok {
		a.render()
		time.Sleep(time.Millisecond * 1500)
		a.Dice.Roll()
		a.render()
		ok = a.changePlayerTurn()
	}
}

func (a *Arena) handleKeyboard(k keyboard.KeyboardEvent) bool {
	if a.isGameOver() {
		return k.Key == termbox.KeyEsc
	}

	a.stopBlinkCurPawn()
	a.repaintCurPawn()
	switch k.Key {
	case termbox.KeyArrowRight:
		a.setNextCurPawnAndValidate(1)
	case termbox.KeyArrowLeft:
		a.setNextCurPawnAndValidate(-1)
	case termbox.KeyEnter:
		fallthrough
	case termbox.KeySpace:
		hasDestroyed, hasReachedDest := a.makeMove()
		a.Dice.Roll()
		a.render()
		if !hasDestroyed && !hasReachedDest {
			a.changePlayerTurnAndValidate()
		} else if hasReachedDest {
			if a.curPlayer().IsAllPawnsAtDest() {
				a.setCurPlayerWin()
				if a.isGameOver() {
					a.changePlayerTurn()
					a.setCurPlayerWin()
					a.renderGameOver()
					return false
				}
			}
			if ok := a.setNextCurPawnAndValidate(1); !ok {
				a.changePlayerTurnAndValidate()
			}
		}
	case termbox.KeyEsc:
		return true
	}
	a.render()
	a.startBlinkCurPawn()
	return false
}

func setRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func (a *Arena) runGameLoop() {
	kChan := a.kChan

	go keyboard.ListenToKeyboard(&kChan)
	a.changePlayerTurn(1)
	a.changePlayerTurnAndValidate()
	a.board.SetCurPawn(0)
	setRandSeed()
	a.Dice.Roll()

	a.render()
	a.startBlinkCurPawn()

	for a.curPlayer().IsBot() {
		a.playBot()
		if a.isGameOver() {
			break
		}
	}
mainloop:
	for {
		ev := <-kChan.EvChan
		kChan.Pause()
		if stop := a.handleKeyboard(ev); stop {
			kChan.Stop()
			break mainloop
		}
		for a.curPlayer().IsBot() {
			a.playBot()
			if a.isGameOver() {
				break
			}
		}
		kChan.Resume()
	}
}

func StartGameOffline(players []common.PlayerData) {
	participantsCount := 0
	for _, p := range players {
		if p.Type != "-" {
			participantsCount++
		}
	}

	gameDice := common.Dice{}
	kChan := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}

	a := Arena{
		participantsCount: participantsCount,
		board:             board.LudoBoard{},
		players:           players,
		blinkCh:           make(chan bool),
		nextWinningPos:    0,
		Dice:              gameDice,
		bots:              make(map[int][4]int),
		kChan:             kChan,
	}
	a.board.SetupBoard(players)
	a.botsInit()
	a.runGameLoop()
}

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
