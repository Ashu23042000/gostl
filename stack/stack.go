package stack

import "fmt"

// custom dataType
type dataType interface{}

// stack struct
type Stack[T dataType] struct {
	data []T
}

// create new stack
func New[T dataType]() *Stack[T] {
	return &Stack[T]{}
}

// push element into stack
func (s *Stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

// remove top element from stack
func (s *Stack[T]) Pop() {
	if len(s.data) <= 0 {
		fmt.Println("stack is empty")
		return
	}
	s.data = s.data[0 : len(s.data)-1]
}

// return length of stack
func (s *Stack[T]) Length() int {
	return len(s.data)
}

// return top element of stack
func (s *Stack[T]) Top() (T, bool) {
	if len(s.data) <= 0 {
		var empty T
		return empty, false
	}
	return s.data[len(s.data)-1], true
}

// print all elements of stack
func (s *Stack[T]) Print() {
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Println(s.data[i])
	}
}
