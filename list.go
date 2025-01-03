// Collection is a Go library aim to implement some basic data structure such as List, Queue, Stack, Heap and more.
package collection

import (
	"fmt"
	"iter"
	"sync"
)

// [List] is a doubly linked list implementation.
// All operation on [List] should be thread-safe,
// because it only allow one goroutine at a time to access it data.
type List[T any] struct {
	length     int
	head, tail *node[T]
	mux        sync.Mutex
}

// [NewList] creates a new doubly linked [List].
// All operation on [List] should be thread-safe,
// because it only allow one goroutine at a time to access it data.
//
//	emptyList := New[int]()
//	initializedList := New(1, 2, 3, 4, 5)
func NewList[T any](values ...T) *List[T] {
	l := &List[T]{}
	for _, value := range values {
		l.Append(value)
	}
	return l
}

// [Length] returns the number of node in the list.
func (l *List[T]) Length() int {
	l.mux.Lock()
	defer l.mux.Unlock()

	return l.length
}

// [Append] adds new nodes to at the end of the list
func (l *List[T]) Append(values ...T) {
	l.mux.Lock()
	defer l.mux.Unlock()

	for _, value := range values {
		newNode := &node[T]{value: value}
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
	}
	l.length += len(values)
}

// [Prepend] adds a new node at the start of the list.
func (l *List[T]) Prepend(values ...T) {
	l.mux.Lock()
	defer l.mux.Unlock()

	for _, value := range values {
		newNode := &node[T]{value: value}
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
	}
	l.length += len(values)
}

// [Insert] adds a new node after the node at a specified index.
// If the index is less than zero or greater than or equal the current length of the list,
// then this function will return an [ErrIndexOutOfRange] error.
// If you want to insert at the start or at the end of the list use [Prepend] or [Append] instead.
func (l *List[T]) Insert(value T, at int) error {
	l.mux.Lock()
	defer l.mux.Unlock()

	if at < 0 || at >= l.length {
		return &ErrIndexOutOfRange{
			msg: fmt.Sprintf(
				"index of %d is out of range for list of length of %d", at, l.length,
			),
		}
	}

	// Get the node at index specified by idx.
	newNode := &node[T]{value: value}
	currNode := l.head
	for at > 0 {
		currNode = currNode.right
		at--
	}

	// Merge newNode in the the list after the currNode.
	newNode.right = currNode.right
	newNode.left = currNode
	currNode.right = newNode

	// Increase length by one.
	l.length++

	return nil
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
			if !yield(idx, curr.value) {
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
			if !yield(idx, curr.value) {
				return
			}
			curr = curr.left
			idx--
		}
	}
}

// [Search] searches for a value in the list.
// It takes the target to search for and the equal function.
// The equal function takes two arguments value and target,
// it should return [true] if the two arguments is considered to be equal.
//
// It returns an index greater or equal to zero and a nil error if the value existed inside the list,
// else return the index of -1 and error of [ErrNotFound].
//
//	 type user struct {
//	   name string
//	 }
//
//	 func main() {
//		  user1 := user{name: "User 1"}
//		  user2 := user{name: "User 2"}
//		  users := NewList(user1, user2)
//
//		  equal := func (value, target user) bool {
//		    return value.name == target.name
//		  }
//		  idx, err := users.Search(user1, equal)
//		  // code goes here
//	 }
func (l *List[T]) Search(target T, equal func(value, target T) bool) (int, error) {
	for idx, value := range l.All() {
		if equal(value, target) {
			return idx, nil
		}
	}
	return -1, &ErrNotFound{msg: fmt.Sprintf("target of %v not existed in list", target)}
}
