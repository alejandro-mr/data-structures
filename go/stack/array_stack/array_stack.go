package array_stack

import "fmt"

func main() {
	s := New(5)
	fmt.Println(s)
	fmt.Println("Is empty? ", s.IsEmpty())
	for i := 1; i <= 10; i++ {
		s.Push(i)
	}
	pop := s.Pop()
	fmt.Println("Sacando: ", pop)
	fmt.Println(s)
}

type ArrayStack struct {
	data     []int
	top      *int
	capacity int
}

func New(len int) *ArrayStack {
	if len > 0 {
		return &ArrayStack{capacity: len}
	}
	return nil
}

func (s *ArrayStack) isFull() bool {
	return (len(s.data) - 1) == s.capacity-1
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *ArrayStack) Push(v int) *int {
	if !s.isFull() {
		s.data = append(s.data, v)
		return &s.data[len(s.data)-1]
	}
	return nil
}

func (s *ArrayStack) Peek() int {
	if !s.IsEmpty() {
		return s.data[len(s.data)-1]
	}
	return -1
}

func (s *ArrayStack) Pop() int {
	size := len(s.data) - 1
	if size >= 0 {
		pop := s.data[size]
		s.data = s.data[:size]

		return pop
	}
	return -1
}
