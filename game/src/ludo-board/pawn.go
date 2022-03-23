package board

import "github.com/nsf/termbox-go"

func (l *LudoBoard) connectPawnsPosToPath() {
	openingNodes := map[int]int{14: 0, 27: 1, 40: 3, 1: 2}

	for i, j := range openingNodes {
		n := l.pathLayer.LL.GetNodeAt(i - 1)

		n.Next["common"].Cell.Fg = l.Players[j].Color

		if !l.Players[j].IsParticipant() {
			continue
		}

		for nidx := range &l.Players[j].Pawns {
			l.Players[j].Pawns[nidx]["homeNode"].Next = n.Next
		}
	}
}

func (b *LudoBoard) SetCurPawn(idx int) {
	b.CurPawnIdx = idx
}

func (p Pawn) HasNPathsAhead(n int) bool {
	temp := p["curNode"]

	if temp.Next["common"] != nil && temp.Next["toDest"] == nil {
		return true
	}

	for i := 0; i < n; i++ {
		temp = temp.Next["toDest"]
		if temp == nil {
			return false
		}
	}

	return true
}

func (p Pawn) MoveToNext(pathName string, bg termbox.Attribute) {
	p["curNode"] = p["curNode"].Next[pathName]
	p["curNode"].Cell.Bg = bg
}

func (b *LudoBoard) SetNextCurPawn(curTurn, mag int) {
	b.SetCurPawn(b.CurPawnIdx + mag)

	if b.CurPawnIdx < 0 {
		b.CurPawnIdx = len(b.Players[curTurn].Pawns) - 1
	} else if b.CurPawnIdx >= len(b.Players[curTurn].Pawns) {
		b.CurPawnIdx = 0
	}
}

func (p Pawn) IsAtDest() bool {
	return p["curNode"].Next["toDest"] == nil && p["curNode"].Next["common"] == nil
}
