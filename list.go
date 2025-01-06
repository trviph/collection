package collection

import (
	"fmt"
	"iter"
	"sync"

	"github.com/trviph/collection/internal"
)

// [List] is a doubly linked list implementation.
// All operation on [List] is thread-safe,
// because it only allow one goroutine at a time to access it data.
type List[T any] struct {
	mu         sync.RWMutex
	length     int
	head, tail *internal.Node[T]
}

// Interface guard
var _ internal.List[any] = (*List[any])(nil)

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
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.length
}

// Append adds new nodes to at the end of the list
func (l *List[T]) Append(values ...T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, value := range values {
		l.append(value)
	}
}

func (l *List[T]) append(value T) {
	newNode := &internal.Node[T]{Value: value}
	if l.isEmpty() {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Insert(newNode)
		l.tail = newNode
	}
	l.length++
}

// Prepend adds a new node at the start of the list.
func (l *List[T]) Prepend(values ...T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, value := range values {
		l.prepend(value)
	}
}

func (l *List[T]) prepend(value T) {
	newNode := &internal.Node[T]{Value: value}
	if l.isEmpty() {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Insert(l.head)
		l.head = newNode
	}
	l.length++
}

// Insert adds a new node after the node at a specified index.
// If the index is less than zero or greater than or equal the current length of the list,
// then this function will return an [ErrIndexOutOfRange] error.
// If you want to insert at the start of the list use [List.Prepend] instead.
func (l *List[T]) Insert(value T, after int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if err := l.checkIndex(after); err != nil {
		return fmt.Errorf("failed to insert into list, cause by %w", err)
	}

	// If insert after tail, then append
	if after == l.length-1 {
		l.append(value)
		return nil
	}

	newNode := &internal.Node[T]{Value: value}
	idxNode := l.getNode(after)
	idxNode.Insert(newNode)
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
		l.mu.RLock()
		defer l.mu.RUnlock()

		curr := l.head
		idx := 0
		for curr != nil {
			if !yield(idx, curr.Value) {
				break
			}
			curr = curr.Right
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
func (l *List[T]) all() iter.Seq2[int, *internal.Node[T]] {
	return func(yield func(int, *internal.Node[T]) bool) {
		curr := l.head
		idx := 0
		for curr != nil {
			if !yield(idx, curr) {
				break
			}
			curr = curr.Right
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
		l.mu.RLock()
		defer l.mu.RUnlock()

		curr := l.tail
		idx := l.length - 1
		for curr != nil {
			if !yield(idx, curr.Value) {
				break
			}
			curr = curr.Left
			idx--
		}
	}
}

// This is similar to [Backward], there is however three differences.
// First this is private method,
// second that this function return a [node] instead of a value of type T,
// and third it does not using [sync.Mutex] to lock access.
// The [sync.Mutex] lock should only be called and managed by a public method to avoid
// race condition and deadlock.
func (l *List[T]) backward() iter.Seq2[int, *internal.Node[T]] {
	return func(yield func(int, *internal.Node[T]) bool) {
		curr := l.tail
		idx := l.length - 1
		for curr != nil {
			if !yield(idx, curr) {
				break
			}
			curr = curr.Left
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
	l.mu.RLock()
	defer l.mu.RUnlock()

	for idx, node := range l.all() {
		if equal(node.Value, target) {
			return idx, nil
		}
	}
	return -1, fmt.Errorf("target of of %v not existed in list, cause by %w", target, ErrNotFound)
}

// Index gets value at the specified index.
// If the index is out of range, it will return [ErrIndexOutOfRange] as error.
func (l *List[T]) Index(at int) (T, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var zeroValue T
	if l.isEmpty() {
		return zeroValue, fmt.Errorf("failed to get value at index %d from list, cause by %w", at, ErrIsEmpty)
	}
	if err := l.checkIndex(at); err != nil {
		return zeroValue, fmt.Errorf("failed to get value at index %d from list, cause by %w", at, err)
	}

	return l.getNode(at).Value, nil
}

// Pop removes and returns the last element of the list.
// If the list is empty then return [ErrIsEmpty] as an error.
func (l *List[T]) Pop() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.isEmpty() {
		var zeroValue T
		return zeroValue, fmt.Errorf("failed to pop from list, cause by %w", ErrIsEmpty)
	}
	return l.pop()
}

func (l *List[T]) pop() (T, error) {
	value := l.tail.Value
	l.tail = l.tail.Left
	if l.tail != nil {
		l.tail.Right = nil
	}

	l.length--
	// Popped the last value
	if l.isEmpty() {
		l.head = nil
		l.tail = nil
	}

	return value, nil
}

// Dequeue removes and returns the first element of the list.
// If the list is empty then return [ErrIsEmpty] as an error.
func (l *List[T]) Dequeue() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.isEmpty() {
		var zeroValue T
		return zeroValue, fmt.Errorf("failed to dequeue from list, cause by %w", ErrIsEmpty)
	}
	return l.dequeue()
}

func (l *List[T]) dequeue() (T, error) {
	value := l.head.Value
	l.head = l.head.Right
	if l.head != nil {
		l.head.Left = nil
	}

	// Dequeued the last value
	l.length--
	if l.isEmpty() {
		l.tail = nil
		l.head = nil
	}

	return value, nil
}

// Remove removes and returns the element at the specified index of the list.
// If the index is out of range then return [ErrIndexOutOfRange] as an error.
// Or if the list is empty then return [ErrIsEmpty] as an error.
func (l *List[T]) Remove(at int) (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Check if index is valid
	var zeroValue T
	if l.isEmpty() {
		return zeroValue, fmt.Errorf("failed to remove from list, cause by %w", ErrIsEmpty)
	}
	if err := l.checkIndex(at); err != nil {
		return zeroValue, fmt.Errorf("failed to remove from list, cause by %w", err)
	}

	// If Remove(0) then dequeue
	if at == 0 {
		return l.dequeue()
	}

	// If remove(lastIDX) then pop
	if at == l.length-1 {
		return l.pop()
	}

	// Get the node at the specified index and remove it
	idxNode := l.getNode(at)
	idxNode.Unlink()
	l.length--

	return idxNode.Value, nil
}

// Get the node at the specified index, should be called after [checkIndex]
// to avoid null pointer error.
func (l *List[T]) getNode(at int) *internal.Node[T] {
	// If the specified index in the left half of the list then
	// we should iterate from head -> tail.
	it := l.all
	// Else if the specified index in the right half of the list then
	// we should iterate from tail -> head.
	if at > (l.length / 2) {
		it = l.backward
	}

	for idx, node := range it() {
		if idx == at {
			return node
		}
	}
	// This panic have happened in the past due to:
	// - Jan 4th 25 (trviph) - Wrong implementation of List.Insert and List.Remove causing the nodes to not linked properly.
	panic(
		fmt.Errorf("something went very wrong: cannot find node with index %d in a list of length %d", at, l.length),
	)
}

func (l *List[T]) checkIndex(at int) error {
	if at < 0 || at >= l.length {
		return ErrIndexOutOfRange
	}
	return nil
}

func (l *List[T]) isEmpty() bool {
	return l.length == 0
}

// Return the head node of the [List].
//
// BUG(trviph): This function is needed for the cache package,
// but leaks [internal.Node] to the users.
// We could either find a way to remove this
// or could just ignore this and leave it hear.
func (l *List[T]) Head() *internal.Node[T] {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.head
}
