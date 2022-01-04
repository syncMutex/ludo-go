package game

type node struct {
	cell cell
	next map[string]*node
}

type linkedlist struct {
	head *node
}

func (l *linkedlist) addEnd(c cell, pathName string, temp *node) *node {
	newNode := &node{cell: c, next: map[string]*node{}}

	if l.head == nil {
		l.head = newNode
	} else {
		if temp == nil {
			temp = l.head
		}
		for temp.next[pathName] != nil {
			temp = temp.next[pathName]
		}

		temp.next[pathName] = newNode
	}
	return newNode
}

func (l *linkedlist) getNodeAt(idx int) *node {
	n := l.head

	for idx > 0 {
		n = n.next["common"]
		idx--
	}

	return n
}
