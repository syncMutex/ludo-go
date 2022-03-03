package game

import (
	"ludo/src/keyboard"
	"time"

	"github.com/nsf/termbox-go"
)

func (a *Arena) playBot() {
	time.Sleep(time.Millisecond * 100)
	a.chooseBestPawn()
	a.handleKeyboard(keyboard.KeyboardEvent{Key: termbox.KeyEnter, Ch: ' '})
}

func (a Arena) isEnemyPawnAt(x, y int) bool {
	for i, p := range a.board.players {
		if i == a.curTurn || !p.isParticipant() {
			continue
		}

		for _, _pawn := range p.pawns {
			c := _pawn["curNode"].cell
			if c.x == x && c.y == y {
				return true
			}
		}
	}

	return false
}

func getPawnAfter(curNode *node, step int) *node {
	temp := curNode

	for i := 0; i < step; i++ {
		if temp == nil {
			break
		}
		temp = temp.next["common"]
	}

	return temp
}

func (a *Arena) chooseBestPawn() {
	temp := a.board.curPawnIdx
	bestPawn := temp
	maxMoveProgressed := a.bots[a.curTurn][bestPawn]

	for {
		n := getPawnAfter(a.curPawn()["curNode"], a.Dice.Value)
		if n != nil && a.isEnemyPawnAt(n.cell.x, n.cell.y) {
			b := a.bots[a.curTurn]
			b[bestPawn] += a.Dice.Value
			return
		}
		if maxMoveProgressed < a.bots[a.curTurn][temp] {
			maxMoveProgressed = a.bots[a.curTurn][temp]
			bestPawn = a.board.curPawnIdx
		}
		a.setNextCurPawnAndValidate(1)
		if a.board.curPawnIdx == temp {
			break
		}
	}
	b := a.bots[a.curTurn]
	b[bestPawn] += a.Dice.Value
	a.board.setCurPawn(bestPawn)
}

func (a *Arena) resetBotPawn(botIdx, pawnIdx int) {
	b := a.bots[botIdx]
	b[pawnIdx] = 0
}

func (a *Arena) botsInit() {
	for idx, p := range a.board.players {
		if p.isBot() {
			a.bots[idx] = [4]int{}
		}
	}
}