package main

import (
	"fmt"

	"github.com/sagargaikwad0/gostl/stack"
)

func main() {
	s := stack.New[int]()

	// s.Push(1)
	// s.Push(2)
	s.Print()
	s.Pop()
	// s.Push(3)
	fmt.Println()
	top, ok := s.Top()
	if !ok {
		fmt.Println("Stack is empty")
	}
	fmt.Println(top)

	// s.Print()

}
