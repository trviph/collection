package collection

import (
	"sync"
)

type List[T any] struct {
	length     int
	head, tail *Node[T]
	mux        sync.Mutex
}

// NewList creates a new doubly linked list
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

// Length returns the number of node in the list
func (l *List[T]) Length() int {
	l.mux.Lock()
	defer l.mux.Unlock()

	return l.length
}

// Append adds a new node after the end of the list
func (l *List[T]) Append(data T) {
	l.mux.Lock()
	defer l.mux.Unlock()

	newNode := &Node[T]{data: data}
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

// Prepend adds a new node before the start of the list
func (l *List[T]) Prepend(data T) {
	l.mux.Lock()
	defer l.mux.Unlock()

	newNode := &Node[T]{data: data}
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
