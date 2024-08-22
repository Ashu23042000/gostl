package linkedlist

import "fmt"

// custom dataType
type dataType interface {
	// ~int | ~string | ~float64
}

// node struct
type node[T dataType] struct {
	data T
	next *node[T]
}

// list struct
type List[T dataType] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// create new list
func New[T dataType]() *List[T] {
	return &List[T]{}
}

// create new node
func newnode[T dataType](data T) *node[T] {
	return &node[T]{data: data}
}

// insert at begening of list
func (l *List[T]) InsertAtHead(data T) {
	if l.head != nil {
		temp := l.head
		l.head = newnode[T](data)
		l.head.next = temp
		l.length++
		return
	}

	l.head = newnode[T](data)
	l.tail = l.head
	l.length++
}

// insert at end of list
func (l *List[T]) InsertAtTail(data T) {

	if l.head == nil {
		l.InsertAtHead(data)
		return
	}

	l.tail.next = newnode[T](data)
	l.tail = l.tail.next
	l.length++
}

// print data of each node from list
func (l *List[T]) Print() {
	temp := l.head

	for temp != nil {
		fmt.Printf("%v -> ", temp.data)
		temp = temp.next
	}

	fmt.Println("nil")
}

// insert new node at given position
func (l *List[T]) InsertAtPosition(pos int, data T) {
	if pos == 0 {
		l.InsertAtHead(data)
		return
	}

	if pos > l.length {
		fmt.Println("Invalid Position")
		return
	}

	cnt := 1
	temp := l.head

	for cnt != pos && cnt <= l.length {
		cnt++
		temp = temp.next
	}

	new := newnode[T](data)
	new.next = temp.next

	if temp.next == nil {
		temp.next = new
		l.tail = temp.next
	} else {
		temp.next = new
	}

	l.length++
}

// delete node at given position
func (l *List[T]) DeleteAtPosition(pos int) {
	if pos < 0 {
		fmt.Println("Invalid Position, list start from 0th position")
		return
	}

	if pos > l.length-1 {
		fmt.Println("Invalid Position, list start from 0th position")
		return
	}

	if pos == 0 {
		l.DeleteHead()
		return
	}

	if pos == l.length-1 {
		l.DeleteTail()
		return
	}

	cnt := 1
	temp := l.head.next
	prvs := l.head

	for cnt != pos && cnt < l.length {
		cnt++
		prvs = temp
		temp = temp.next
	}

	prvs.next = temp.next
	l.length--
}

// delete head node
func (l *List[T]) DeleteHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

// delete tail node
func (l *List[T]) DeleteTail() {
	if l.head == nil {
		return
	}
	temp := l.head
	var prev *node[T]
	for temp.next != nil {
		prev = temp
		temp = temp.next
	}
	prev.next = nil
}

// returns length of linked list
func (l *List[T]) Length() int {
	return l.length
}
