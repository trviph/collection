package collection

import (
	"fmt"
	"sync"
)

// A first-in-first-out [Queue] implemented by using [List] as the base.
// Since [List] is thread-safe, [Queue] should also be thread-safe.
type Queue[T any] struct {
	mu   sync.Mutex
	list *List[T]
}

// Interface guard
var _ queue[any] = (*Queue[any])(nil)

// [NewQueue] creates a new [Queue] of type T.
func NewQueue[T any](values ...T) *Queue[T] {
	return &Queue[T]{list: NewList(values...)}
}

// Length returns the number of values current in the queue.
func (q *Queue[T]) Length() int {
	return q.list.Length()
}

// Push a list of values in to the queue, starting from left to right.
func (q *Queue[T]) Push(values ...T) {
	q.list.Append(values...)
}

// Dequeue get the value from the front of the queue, and remove it from the queue.
// If the queue is empty return [ErrIsEmpty] as an error.
func (q *Queue[T]) Dequeue() (T, error) {
	if value, err := q.list.Dequeue(); err != nil {
		return value, fmt.Errorf("failed to dequeue queue, cause by %w", err)
	} else {
		return value, nil
	}
}

// Front get the value from the front but does not remove it from the queue.
// If the queue is empty return [ErrIsEmpty] as an error.
func (q *Queue[T]) Front() (T, error) {
	if value, err := q.list.Index(0); err != nil {
		return value, fmt.Errorf("failed to peek at the front of the queue, cause by %w", err)
	} else {
		return value, nil
	}
}

// Rear get the value from the rear/end but does not remove it from the queue.
// If the queue is empty return [ErrIsEmpty] as an error.
func (q *Queue[T]) Rear() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if value, err := q.list.Index(q.list.Length() - 1); err != nil {
		return value, fmt.Errorf("failed to peek at the rear of the queue, cause by %w", err)
	} else {
		return value, nil
	}
}
