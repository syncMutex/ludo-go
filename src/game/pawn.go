package game

func (l *ludoBoard) connectPawnsPosToPath() {
	openingNodes := map[int]int{14: 0, 27: 1, 40: 3, 1: 2}

	for i, j := range openingNodes {
		n := l.pathLayer.ll.getNodeAt(i - 1)

		n.next["common"].cell.fg = l.players[j].color

		if !l.players[j].isParticipant() {
			continue
		}

		for nidx := range &l.players[j].pawns {
			l.players[j].pawns[nidx]["homeNode"].next = n.next
		}
	}
}

func (b *ludoBoard) setCurPawn(idx int) {
	b.curPawnIdx = idx
}

func (p pawn) hasNPathsAhead(n int) bool {
	temp := p["curNode"]

	if temp.next["common"] != nil && temp.next["toDest"] == nil {
		return true
	}

	for i := 0; i < n; i++ {
		temp = temp.next["toDest"]
		if temp == nil {
			return false
		}
	}

	return true
}

func (b *ludoBoard) setNextCurPawn(curTurn, mag int) {
	b.setCurPawn(b.curPawnIdx + mag)

	if b.curPawnIdx < 0 {
		b.curPawnIdx = len(b.players[curTurn].pawns) - 1
	} else if b.curPawnIdx >= len(b.players[curTurn].pawns) {
		b.curPawnIdx = 0
	}
}

func (a *Arena) repaintCurPawn() {
	curCell := a.curPawn()["curNode"].cell
	setBg(curCell.x, curCell.y, a.board.players[a.curTurn].color)
}

func (p pawn) isAtDest() bool {
	if p["curNode"].next["toDest"] == nil && p["curNode"].next["common"] == nil {
		return true
	}
	return false
}
