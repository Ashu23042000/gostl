package linkedlist

import "fmt"

// custom datatype
type DataType interface {
	~int | ~string | ~float64
}

// node struct
type Node[T DataType] struct {
	Data T
	Next *Node[T]
}

// list struct
type List[T DataType] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

// create new list
func New[T DataType](data T) *List[T] {
	return &List[T]{}
}

// create new node
func newNode[T DataType](data T) *Node[T] {
	return &Node[T]{}
}

// insert at begening of list
func (l *List[T]) InsertAtHead(data T) {
	if l.Head != nil {
		temp := l.Head
		l.Head = newNode[T](data)
		l.Head.Next = temp
		l.Length++
		return
	}

	l.Head = newNode[T](data)
	l.Tail = l.Head
	l.Length++

}

// insert at end of list
func (l *List[T]) InsertAtTail(data T) {
	l.Tail.Next = newNode[T](data)
	l.Tail = l.Tail.Next
	l.Length++

}

// print data of each node from list
func (l *List[T]) Print() {
	temp := l.Head

	for temp != nil {
		fmt.Printf("%v ", temp.Data)
		temp = temp.Next
	}

	fmt.Println()
}

// func (l *List[T]) Push(data T) {
// 	if l.Head == nil {
// 		l.InsertAtHead(data)
// 		return
// 	}

// 	l.Tail.Next = newNode[T](data)
// 	l.Tail = l.Tail.Next
// 	l.Length++

// }

// insert new node at given position
func (l *List[T]) InsertAtPosition(pos int, data T) {
	if pos == 0 {
		l.InsertAtHead(data)
		return
	}

	if pos > l.Length {
		fmt.Println("Invalid Position")
		return
	}

	cnt := 0
	temp := l.Head

	for cnt != pos && cnt <= l.Length {
		cnt++
		temp = temp.Next
	}

	new := newNode[T](data)
	new.Next = temp.Next

	if temp.Next == nil {
		temp.Next = new
		l.Tail = temp.Next
	} else {
		temp.Next = new
	}

	l.Length++

}

// delete node at given position
func (l *List[T]) DeleteAtPosition(pos int, data T) {
	if pos == 0 {
		fmt.Println("Invalid Position, Head is at position 1st")
		return
	}

	if pos > l.Length {
		fmt.Println("Invalid Position")
		return
	}

	cnt := 0
	temp := l.Head

	var prvs *Node[T]

	for cnt != pos && cnt <= l.Length {
		cnt++
		prvs = temp
		temp = temp.Next
	}

	if temp.Next != nil {
		prvs = temp
		prvs.Next = temp.Next
	} else {
		l.Tail = prvs
	}

	l.Length--
}
