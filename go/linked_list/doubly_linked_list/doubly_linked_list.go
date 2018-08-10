package main

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

func (l *LinkedList) IsEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Tail() *Node {
	return l.tail
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Insert(n Node) int {
	if l.IsEmpty() {
		l.head, l.tail = &n, &n
		l.size++
		return 0
	}
	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	n.Prev = current
	current.Next = &n
	l.tail = &n

	l.size++

	return 0
}

func (l *LinkedList) InsertFirst(n Node) int {
	if l.IsEmpty() {
		l.head, l.tail = &n, &n
	} else {
		n.Next = l.head
		l.head.Prev = &n
		l.head = &n
	}
	l.size++

	return 0
}

func (l *LinkedList) InsertLast(n Node) int {
	if l.IsEmpty() {
		return 1
	}
	l.tail.Next = &n
	n.Prev = l.tail
	l.tail = &n

	l.size++

	return 0
}

func (l *LinkedList) InsertBefore(n Node, before Node) int {
	if l.IsEmpty() {
		return 1
	}
	current := l.head
	// We are inserting before `head`
	if current.Value == before.Value {
		n.Next = l.head
		l.head.Prev = &n
		l.head = &n

		l.size++
		return 0
	}
	for current.Next != nil {
		if current.Value == before.Value {
			n.Next, n.Prev = current, current.Prev
			current.Prev.Next, current.Prev = &n, &n

			l.size++
			return 0
		}
		current = current.Next
	}
	// Current is equal to `tail`
	if l.tail.Value == before.Value {
		n.Next, n.Prev = l.tail, l.tail.Prev
		l.tail.Prev.Next, l.tail.Prev = &n, &n

		l.size++
		return 0
	}
	return 1
}

func (l *LinkedList) Remove(n Node) int {
	if l.IsEmpty() {
		return 1
	}
	current := l.head
	if current.Value == n.Value {
		l.head = l.head.Next
		l.head.Prev = nil

		l.size--
		return 0
	}
	current = current.Next
	for current.Next != nil {
		if current.Value == n.Value {
			current.Prev.Next, current.Next.Prev = current.Next, current.Prev
			l.size--
			return 0
		}
		current = current.Next
	}
	if l.tail.Value == n.Value {
		l.tail = l.tail.Prev
		l.tail.Next = nil

		l.size--
		return 0
	}
	return 1
}

func (l *LinkedList) RemoveFirst() int {
	if l.IsEmpty() {
		return 1
	}
	if l.size > 1 {
		l.head = l.head.Next
		l.head.Prev = nil
	} else {
		l.head = nil
		l.tail = nil
	}

	l.size--
	return 0
}

func (l *LinkedList) RemoveLast() int {
	if l.IsEmpty() {
		return 1
	}
	if l.size > 1 {
		l.tail = l.tail.Prev
		l.tail.Next = nil
	} else {
		l.tail = nil
		l.head = nil
	}

	l.size--
	return 0
}

func (l *LinkedList) Reverse() {
	if !l.IsEmpty() {
		current := l.head
		for current != nil {
			current.Next, current.Prev = current.Prev, current.Next
			current = current.Prev
		}
		l.tail, l.head = l.head, l.tail
	}
}

func (l *LinkedList) clear() int {
	current := l.head
	for current.Next != nil {
		pointer := current
		current.Prev = nil
		pointer.Next = nil
		current = current.Next
	}
	l.size = 0
	return 0
}
