package ll

import "ludo/src/ludo-board/cellmap"

type Node struct {
	Cell cellmap.Cell
	Next map[string]*Node
}

type Linkedlist struct {
	Head *Node
}

func (l *Linkedlist) AddEnd(c cellmap.Cell, pathName string, temp *Node) *Node {
	newNode := &Node{Cell: c, Next: map[string]*Node{}}

	if l.Head == nil {
		l.Head = newNode
	} else {
		if temp == nil {
			temp = l.Head
		}
		for temp.Next[pathName] != nil {
			temp = temp.Next[pathName]
		}

		temp.Next[pathName] = newNode
	}
	return newNode
}

func (l *Linkedlist) GetNodeAt(idx int) *Node {
	n := l.Head

	for idx > 0 {
		n = n.Next["common"]
		idx--
	}

	return n
}
