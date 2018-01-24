package linked_list

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	Value int
	Next  *Node
}

func (l *LinkedList) Empty() bool {
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

func (l *LinkedList) Lenght() int {
	return l.length
}

func (l *LinkedList) InsertBefore(in *Node, before *Node) *Node {
	if !l.Empty() {
		n, prev := l.head, l.head
		for n != nil {
			if n.Value == before.Value {
				if n == l.head {
					return l.Push(in)
				}

				prev.Next = in
				in.Next = n
				l.length++

				return in
			}
			prev = n
			n = n.Next
		}
		return nil
	}
	return nil
}

func (l *LinkedList) InsertAfter(in *Node, after *Node) *Node {
	if !l.Empty() {
		n := l.head
		for n != nil {
			if n.Value == after.Value {
				if n == l.tail {
					return l.InsertLast(in)
				}
				in.Next = n.Next
				n.Next = in
				l.length++

				return in
			}
			n = n.Next
		}
		return nil
	}
	return nil
}

func (l *LinkedList) InsertLast(n *Node) *Node {
	if l.Empty() {
		return l.Push(n)
	}
	l.tail.Next = n
	l.tail = n
	l.length++
	return n
}

func (l *LinkedList) Reverse() {
	n := l.head.Next
	prev := l.head
	for n != nil {
		next := n.Next
		n.Next = prev
		prev = n
		n = next
	}
	l.head, l.tail = l.tail, l.head
	l.tail.Next = nil
}

func (l *LinkedList) Remove(s *Node) *Node {
	if !l.Empty() {
		n, prev := l.head, l.head
		for n != nil {
			if n.Value == s.Value {
				l.length--
				if n == l.head {
					l.head = n.Next
					return n
				}
				if n == l.tail {
					l.tail = prev
					l.tail.Next = nil
					return n
				}
				prev.Next = n.Next
				return n
			}
			prev = n
			n = n.Next
		}
		return nil
	}
	return nil
}

func (l *LinkedList) Push(n *Node) *Node {
	if l.Empty() {
		l.head, l.tail = n, n
	} else {
		prev := l.head
		n.Next = prev
		l.head = n
	}
	l.length++
	return n
}

func (l *LinkedList) Pop() *Node {
	if !l.Empty() {
		n := l.head
		l.head = l.head.Next
		l.length--

		return n
	}
	return nil
}
