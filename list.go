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

// Interface guard
var _ list[any] = (*List[any])(nil)

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

// Length returns the number of node in the list.
func (l *List[T]) Length() int {
	l.mux.Lock()
	defer l.mux.Unlock()

	return l.length
}

// Append adds new nodes to at the end of the list
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

// Prepend adds a new node at the start of the list.
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

// Insert adds a new node after the node at a specified index.
// If the index is less than zero or greater than or equal the current length of the list,
// then this function will return an [ErrIndexOutOfRange] error.
// If you want to insert at the start or at the end of the list use [List.Prepend] or [List.Append] instead.
func (l *List[T]) Insert(value T, at int) error {
	l.mux.Lock()
	defer l.mux.Unlock()

	if err := l.checkIndex(at); err != nil {
		return err
	}

	// Get the node at index specified by idx.
	idxNode := l.getNode(at)

	// Create a new node
	newNode := &node[T]{value: value}

	// Merge newNode in the the list after the currNode.
	newNode.right = idxNode.right
	newNode.left = idxNode
	idxNode.right = newNode

	// Increase length by one.
	l.length++

	return nil
}

// All return an iterator of elements in list going from head to tail.
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

// This is similar to [All], there is however three differences.
// First this is private method,
// second that this function return a [node] instead of a value of type T,
// and third it does not using [sync.Mutex] to lock access.
// The [sync.Mutex] lock should only be called and managed by a public method to avoid
// race condition and deadlock.
func (l *List[T]) all() iter.Seq2[int, *node[T]] {
	return func(yield func(int, *node[T]) bool) {
		curr := l.head
		idx := 0
		for curr != nil {
			if !yield(idx, curr) {
				return
			}
			curr = curr.right
			idx++
		}
	}
}

// Backward return an iterator of elements in list going from tail to head.
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

// Search searches for a value in the list.
// It takes the target to search for and the equal function.
// The equal function takes two arguments value and target,
// it should return true if the two arguments is considered to be equal.
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
	l.mux.Lock()
	defer l.mux.Unlock()

	for idx, node := range l.all() {
		if equal(node.value, target) {
			return idx, nil
		}
	}
	return -1, &ErrNotFound{msg: fmt.Sprintf("target of %v not existed in list", target)}
}

// Index gets value at the specified index.
// If the index is out of range, it will return [ErrIndexOutOfRange] as error.
func (l *List[T]) Index(at int) (T, error) {
	l.mux.Lock()
	defer l.mux.Unlock()

	var res T
	if err := l.checkIndex(at); err != nil {
		return res, err
	}
	for idx, node := range l.all() {
		if idx == at {
			res = node.value
			break
		}
	}
	return res, nil
}

// Pop removes and returns the last element of the list.
// If the list is empty then return [ErrIsEmpty] as an error.
func (l *List[T]) Pop() (T, error) {
	l.mux.Lock()
	defer l.mux.Unlock()

	var value T
	if l.length == 0 {
		return value, &ErrIsEmpty{msg: "list is empty"}
	}

	value = l.tail.value
	l.tail = l.tail.left
	l.length--

	// Popped the last value
	if l.length == 0 {
		l.head = nil
	}

	return value, nil
}

// Dequeue removes and returns the first element of the list.
// If the list is empty then return [ErrIsEmpty] as an error.
func (l *List[T]) Dequeue() (T, error) {
	l.mux.Lock()
	defer l.mux.Unlock()

	var value T
	if l.length == 0 {
		return value, &ErrIsEmpty{msg: "list is empty"}
	}

	value = l.head.value
	l.head = l.head.right
	l.length--

	// Dequeued the last value
	if l.length == 0 {
		l.tail = nil
	}

	return value, nil
}

// Remove removes and returns the element at the specified index of the list.
// If the index is out of range then return [ErrIndexOutOfRange] as an error.
// Or if the list is empty then return [ErrIsEmpty] as an error.
func (l *List[T]) Remove(at int) (T, error) {
	l.mux.Lock()
	defer l.mux.Unlock()

	var value T
	// Check if index is valid
	if l.length == 0 {
		return value, &ErrIsEmpty{msg: "list is empty"}
	}
	if err := l.checkIndex(at); err != nil {
		return value, err
	}

	// Get the node at the specified index
	idxNode := l.getNode(at)

	// Remove idxNode from the list
	left := idxNode.left
	right := idxNode.right
	left.right = right
	right.left = left

	value = idxNode.value
	l.length--
	return value, nil
}

// Get the node at the specified index, should be called after [checkIndex]
// to avoid null pointer error.
func (l *List[T]) getNode(at int) *node[T] {
	for idx, node := range l.all() {
		if idx == at {
			return node
		}
	}
	return nil
}

func (l *List[T]) checkIndex(at int) error {
	if at < 0 || at >= l.length {
		return &ErrIndexOutOfRange{
			msg: fmt.Sprintf(
				"index of %d is out of range for list of length of %d", at, l.length,
			),
		}
	}
	return nil
}
