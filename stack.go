package collection

// A first-in-last-out [Stack] implemented by using [List] as the base.
// Since [List] is thread-safe, [Stack] should also be thread-safe.
type Stack[T any] struct {
	list *List[T]
}

// Interface guard
var _ stack[any] = (*Stack[any])(nil)

// [NewStack] creates a new [Stack] of type T.
func NewStack[T any](values ...T) *Stack[T] {
	return &Stack[T]{list: NewList(values...)}
}

// Length returns the number of values current in the stack.
func (s *Stack[T]) Length() int {
	return s.list.length
}

// Push a list of values in to the stack, starting from left to right.
func (s *Stack[T]) Push(values ...T) {
	s.list.Append(values...)
}

// Pop get the value of the last push, and remove the value from the stack.
// If the stack is empty return [ErrIsEmpty] as an error.
func (s *Stack[T]) Pop() (T, error) {
	var value T
	if s.list.Length() == 0 {
		return value, &ErrIsEmpty{msg: "stack is empty"}
	}
	return s.list.Pop()
}

// Top get the value of the last push but does not remove the value from the stack.
// If the stack is empty return [ErrIsEmpty] as an error.
func (s *Stack[T]) Top() (T, error) {
	var value T
	if s.list.Length() == 0 {
		return value, &ErrIsEmpty{msg: "stack is empty"}
	}
	return s.list.Index(s.Length() - 1)
}
