// Collection is a Go library aim to implement some basic data structure such as List, Queue, Stack, Heap and more.
package collection

import (
	"iter"
	"sync"
)

// [List] is a doubly linked list implementation.
// All operation on [List] should be thread-safe,
// because it only allow one goroutine at a time to access it data.
type List[T any] struct {
	length     int
	head, tail *Node[T]
	mux        sync.Mutex
}

// Create a new doubly linked [List].
// All operation on [List] should be thread-safe,
// because it only allow one goroutine at a time to access it data.
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

// Return the number of node in the list
func (l *List[T]) Length() int {
	l.mux.Lock()
	defer l.mux.Unlock()

	return l.length
}

// Append new nodes to at the end of the list
func (l *List[T]) Append(data ...T) {
	l.mux.Lock()
	defer l.mux.Unlock()

	for _, d := range data {
		newNode := &Node[T]{data: d}
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
}

// [All] return an iterator of elements in list going from head to tail.
// The iterator returns the index and value of the node.
//
//	for idx, val := range list.All() {
//	   // code goes here
//	}
func (l *List[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		l.mux.Lock()
		defer l.mux.Unlock()

		curr := l.head
		idx := 0
		for curr != nil {
			if !yield(idx, curr.data) {
				return
			}
			curr = curr.right
			idx++
		}
	}
}

// [Backward] return an iterator of elements in list going from tail to head.
// The iterator returns the index and value of the node.
//
//	for idx, val := range list.Backward() {
//	   // code goes here
//	}
func (l *List[T]) Backward() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		l.mux.Lock()
		defer l.mux.Unlock()

		curr := l.tail
		idx := l.length - 1
		for curr != nil {
			if !yield(idx, curr.data) {
				return
			}
			curr = curr.left
			idx--
		}
	}
}
