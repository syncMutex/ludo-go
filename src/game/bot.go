package game

import (
	"ludo/src/keyboard"
	ll "ludo/src/linked-list"
	"time"

	"github.com/nsf/termbox-go"
)

func (a *Arena) playBot() {
	time.Sleep(time.Millisecond * 100)
	a.chooseBestPawn()
	a.handleKeyboard(keyboard.KeyboardEvent{Key: termbox.KeyEnter, Ch: ' '})
}

func (a Arena) isEnemyPawnAt(x, y int) bool {
	for i, p := range a.board.Players {
		if i == a.curTurn || !p.IsParticipant() {
			continue
		}

		for _, _pawn := range p.Pawns {
			c := _pawn["curNode"].Cell
			if c.X == x && c.Y == y {
				return true
			}
		}
	}

	return false
}

func getPawnAfter(curNode *ll.Node, step int) *ll.Node {
	temp := curNode

	for i := 0; i < step; i++ {
		if temp == nil {
			break
		}
		temp = temp.Next["common"]
	}

	return temp
}

func (a *Arena) chooseBestPawn() {
	temp := a.board.CurPawnIdx
	bestPawn := temp
	maxMoveProgressed := a.bots[a.curTurn][bestPawn]

	for {
		n := getPawnAfter(a.curPawn()["curNode"], a.Dice.Value)
		if n != nil && a.isEnemyPawnAt(n.Cell.X, n.Cell.Y) {
			b := a.bots[a.curTurn]
			b[bestPawn] += a.Dice.Value
			return
		}
		if maxMoveProgressed < a.bots[a.curTurn][temp] {
			maxMoveProgressed = a.bots[a.curTurn][temp]
			bestPawn = a.board.CurPawnIdx
		}
		a.setNextCurPawnAndValidate(1)
		if a.board.CurPawnIdx == temp {
			break
		}
	}
	b := a.bots[a.curTurn]
	b[bestPawn] += a.Dice.Value
	a.board.SetCurPawn(bestPawn)
}

func (a *Arena) resetBotPawn(botIdx, pawnIdx int) {
	b := a.bots[botIdx]
	b[pawnIdx] = 0
}

func (a *Arena) botsInit() {
	for idx, p := range a.board.Players {
		if p.IsBot() {
			a.bots[idx] = [4]int{}
		}
	}
}
