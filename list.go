package collection

import (
	"fmt"
	"sync"
)

// A doubly linked list.
type List[T any] struct {
	length     int
	head, tail *Node[T]
	mux        sync.Mutex
}

// [NewList] creates a new doubly linked list.
//
//	emptyList := New[int]()
//	initializedList := New(1, 2, 3, 4, 5)
func NewList[T any](data ...T) *List[T] {
	l := &List[T]{}
	for _, d := range data {
		l.Append(d)
	}
	return l
}

// [Length] returns the number of node in the list.
func (l *List[T]) Length() int {
	l.mux.Lock()
	defer l.mux.Unlock()

	return l.length
}

// [Append] adds a new node at the end of the list.
func (l *List[T]) Append(data T) {
	l.mux.Lock()
	defer l.mux.Unlock()

	newNode := &Node[T]{data: data}

	// If there is no head, meaning the current list is empty,
	// then insert as head.
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.left = l.tail
		l.tail.right = newNode
		l.tail = newNode
	}
	l.length++
}

// [Prepend] adds a new node at the start of the list.
func (l *List[T]) Prepend(data T) {
	l.mux.Lock()
	defer l.mux.Unlock()

	newNode := &Node[T]{data: data}

	// If there is no head, meaning the current list is empty,
	// then insert as head.
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.right = l.head
		l.head.left = newNode
		l.head = newNode
	}
	l.length++
}

// [Insert] adds a new node after the node at a specified index.
// If the index is less than zero or greater than or equal the current length of the list,
// then this function will return an [ErrIndexOutOfRange] error.
// If you want to insert at the start or at the end of the list use [Prepend] or [Append] instead.
func (l *List[T]) Insert(data T, idx int) error {
	l.mux.Lock()
	defer l.mux.Unlock()

	if idx < 0 || idx >= l.length {
		return &ErrIndexOutOfRange{msg: fmt.Sprintf("index of %d is out of range for list of length of %d", idx, l.length)}
	}

	// Get the node at index specified by idx.
	newNode := &Node[T]{data: data}
	currNode := l.head
	for idx > 0 {
		currNode = currNode.right
		idx--
	}

	// Merge newNode in the the list after the currNode.
	newNode.right = currNode.right
	newNode.left = currNode
	currNode.right = newNode

	// Increase length by one.
	l.length++

	return nil
}
