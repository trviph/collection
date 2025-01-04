package collection

import (
	"fmt"
	"sync"
)

// A [Heap] implemented using [slice] (dynamic array) as the base.
// Heap is thread-safe, because it only allow one goroutine at a time to access it data.
type Heap[T any] struct {
	mu     sync.RWMutex
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
func (h *Heap[T]) Push(values ...T) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, value := range values {
		h.values = append(h.values, value)
		// Swim the node that just got inserted to it approriate place
		h.swim()
	}
}

// Get a value at the root node, and remove it from the Heap.
// Returns [ErrIsEmpty] if the [Heap] is empty.
func (h *Heap[T]) Pop() (T, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	var res T
	if h.isEmpty() {
		return res, &ErrIsEmpty{"heap is empty"}
	}

	// Get value from the root
	res = h.values[0]

	// Swap the final node to the root
	h.values[0] = h.values[len(h.values)-1]
	// Shorten the underlying array
	h.values = h.values[:len(h.values)-1]
	// Sink the root node down to it apporiate place
	h.sink()

	return res, nil
}

// Push a value into the heap and then pop the root node.
// This function is equivalent to call a [Heap.Push] followed by a [Heap.Pop],
// but have a more efficient implementation.
// Returns [ErrIsEmpty] if the [Heap] is empty.
func (h *Heap[T]) PushPop(value T) (T, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	var res T
	if h.isEmpty() {
		return res, &ErrIsEmpty{"heap is empty"}
	}

	// If the inserted value satisfied cmp condition with the root,
	// then it means that this will become the new root. So set res
	// as value, and do nothing.
	//
	// Else take the root node and replace it with the value, then sink
	// the replaced root to its approriate place.
	if h.cmp(value, h.values[0]) {
		res = value
	} else {
		res = h.values[0]
		h.values[0] = value
		h.sink()
	}

	return res, nil
}

// Peek at the value at the root node without removing it from the Heap.
// Returns [ErrIsEmpty] if the [Heap] is empty.
func (h *Heap[T]) Top() (T, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var res T
	if h.isEmpty() {
		return res, &ErrIsEmpty{"heap is empty"}
	}
	return h.values[0], nil
}

// IsEmpty returns true if the heap does not hold any value.
func (h *Heap[T]) IsEmpty() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.isEmpty()
}

func (h *Heap[T]) isEmpty() bool {
	return len(h.values) == 0
}

// Swim/Heapify-up swim the bottom node toward the root.
func (h *Heap[T]) swim() {
	currIDX := len(h.values) - 1
	curr := h.values[currIDX]
	for currIDX > 0 {
		parentIDX := h.getParentIDX(currIDX)
		parent := h.values[parentIDX]
		if !h.cmp(curr, parent) {
			return
		}
		h.values[parentIDX], h.values[currIDX] = curr, parent
		currIDX = parentIDX
	}
}

// Sink/Heapify-down sink the root down to toward the bottom.
func (h *Heap[T]) sink() {
	currIDX := 0
	for currIDX < len(h.values) {
		if childIDX, ok := h.getChildToSwap(currIDX); !ok {
			return
		} else if ok := h.trySwap(currIDX, childIDX); !ok {
			return
		} else {
			currIDX = childIDX
		}
	}
}

// Get the index of the better child to swap
// If is a max heap get the max child, else if a min heap get min child.
// If there is no child to swap then ok is false.
func (h *Heap[T]) getChildToSwap(parentIDX int) (index int, ok bool) {
	leftIDX := h.getLeftIDX(parentIDX)

	// Because leftIDX always less than rightIDX,
	// so if leftIDX is greater than len(h.values)-1
	// then there is no need to check for rightIDX.
	if leftIDX >= len(h.values) {
		return 0, false
	}

	rightIDX := h.getRightIDX(parentIDX)
	if rightIDX >= len(h.values) {
		return leftIDX, true
	}

	if h.cmp(h.values[leftIDX], h.values[rightIDX]) {
		return leftIDX, true
	}
	return rightIDX, true
}

// Only swap when h.cmp condition is NOT satisfied.
func (h *Heap[T]) trySwap(parentIDX, childIDX int) bool {
	child := h.values[childIDX]
	parent := h.values[parentIDX]
	if h.cmp(parent, child) {
		return false
	}
	h.values[childIDX], h.values[parentIDX] = parent, child
	return true
}

func (h *Heap[T]) getParentIDX(idx int) int {
	return (idx - 1) / 2
}

func (h *Heap[T]) getLeftIDX(idx int) int {
	return idx*2 + 1
}

func (h *Heap[T]) getRightIDX(idx int) int {
	return idx*2 + 2
}
