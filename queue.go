package collection

// A first-in-first-out [Queue] implemented by using [List] as the base.
// Since [List] is thread-safe, [Queue] should also be thread-safe.
type Queue[T any] struct {
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
	return q.list.length
}

// Push a list of values in to the queue, starting from left to right.
func (q *Queue[T]) Push(values ...T) {
	q.list.Append(values...)
}

// Dequeue get the value from the front of the queue, and remove it from the queue.
// If the queue is empty return [ErrIsEmpty] as an error.
func (q *Queue[T]) Dequeue() (T, error) {
	var value T
	if q.list.Length() == 0 {
		return value, &ErrIsEmpty{msg: "queue is empty"}
	}
	return q.list.Dequeue()
}

// Front get the value from the front but does not remove it from the queue.
// If the queue is empty return [ErrIsEmpty] as an error.
func (q *Queue[T]) Front() (T, error) {
	var value T
	if q.list.Length() == 0 {
		return value, &ErrIsEmpty{msg: "queue is empty"}
	}
	return q.list.Index(0)
}

// Rear get the value from the rear/end but does not remove it from the queue.
// If the queue is empty return [ErrIsEmpty] as an error.
func (q *Queue[T]) Rear() (T, error) {
	var value T
	if q.list.Length() == 0 {
		return value, &ErrIsEmpty{msg: "queue is empty"}
	}
	return q.list.Index(q.Length() - 1)
}
