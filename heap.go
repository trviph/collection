package collection

import "fmt"

// A [Heap] implemented using [slice] (dynamic array) as the base.
// Heap is thread-safe, because it only allow one goroutine at a time to access it data.
type Heap[T any] struct {
	values []T
	cmp    func(T, T) bool
}

var _ heap[any] = (*Heap[any])(nil)

// [NewHeap] creates a new [Heap].
// It takes a function that compares two values.
// During swim/heapify-up and sink/heapify-down operation
// if the function returns true, then the nodes will be swapped with eachother.
// Will return an error if cmp is nil, if you want to panic instead use [MustNewHeap].
//
// For example a max Heap[int] would need:
//
//	func maxHeapCmp(current, other int) {
//		return current >= other
//	}
//
// A min Heap[int] would need:
//
//	func mimHeapCmp(current, other int) {
//		return current <= other
//	}
func NewHeap[T any](cmp func(current, other T) bool) (*Heap[T], error) {
	if cmp == nil {
		return nil, fmt.Errorf("function argument is required to create a new heap")
	}

	return &Heap[T]{
		values: make([]T, 0), cmp: cmp,
	}, nil
}

// Like [NewHeap] but will panic if cmp is nil.
func MustNewHeap[T any](cmp func(current, other T) bool) *Heap[T] {
	return Must(func() (*Heap[T], error) {
		return NewHeap(cmp)
	})
}

// Push values into the Heap.
func (h *Heap[T]) Push(values ...T) {}

// Get a value at the root node, and remove it from the Heap.
func (h *Heap[T]) Pop() (T, error) {
	var res T
	return res, nil
}

// Push a value into the heap and then pop the root node.
// This function is equivalent to call a [Heap.Push] followed by a [Heap.Pop],
// but have a more efficient implementation.
func (h *Heap[T]) PushPop(value T) (T, error) {
	var res T
	return res, nil
}

// Peek at the value at the root node without removing it from the Heap.
func (h *Heap[T]) Top() (T, error) {
	var res T
	return res, nil
}

// IsEmpty returns true if the heap does not hold any value.
func (h *Heap[T]) IsEmpty() bool {
	return true
}
