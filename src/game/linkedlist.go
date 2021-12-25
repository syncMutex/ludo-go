package game

type linkedlist struct {
	head *node
	tail *node
	size int
}

func (l *linkedlist) addEnd(c cell) {
	newNode := &node{cell: c, next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	l.tail = newNode

	l.size++
}

func (l *linkedlist) addStart(c cell) {
	newNode := &node{cell: c, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}

	l.size++
}
