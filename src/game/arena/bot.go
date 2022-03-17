package arena

import (
	ll "ludo/src/linked-list"
	"time"
)

func (a *Arena) PlayBot(botFunc func(a *Arena)) {
	time.Sleep(time.Millisecond * 100)
	a.ChooseBestPawn()
	botFunc(a)
}

func (a Arena) IsEnemyPawnAt(x, y int) bool {
	for i, p := range a.Board.Players {
		if i == a.CurTurn || !p.IsParticipant() {
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

func GetPawnAfter(curNode *ll.Node, step int) *ll.Node {
	temp := curNode

	for i := 0; i < step; i++ {
		if temp == nil {
			break
		}
		temp = temp.Next["common"]
	}

	return temp
}

func (a *Arena) ChooseBestPawn() {
	temp := a.Board.CurPawnIdx
	bestPawn := temp
	maxMoveProgressed := a.Bots[a.CurTurn][bestPawn]

	for {
		n := GetPawnAfter(a.CurPawn()["curNode"], a.Dice.Value)
		if n != nil && a.IsEnemyPawnAt(n.Cell.X, n.Cell.Y) {
			b := a.Bots[a.CurTurn]
			b[bestPawn] += a.Dice.Value
			return
		}
		if maxMoveProgressed < a.Bots[a.CurTurn][temp] {
			maxMoveProgressed = a.Bots[a.CurTurn][temp]
			bestPawn = a.Board.CurPawnIdx
		}
		a.SetNextCurPawnAndValidate(1)
		if a.Board.CurPawnIdx == temp {
			break
		}
	}
	b := a.Bots[a.CurTurn]
	b[bestPawn] += a.Dice.Value
	a.Board.SetCurPawn(bestPawn)
}

func (a *Arena) ResetBotPawn(botIdx, pawnIdx int) {
	b := a.Bots[botIdx]
	b[pawnIdx] = 0
}

func (a *Arena) BotsInit() {
	for idx, p := range a.Board.Players {
		if p.IsBot() {
			a.Bots[idx] = [4]int{}
		}
	}
}
