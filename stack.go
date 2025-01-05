package collection

import "fmt"

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
	return s.list.Length()
}

// Push a list of values in to the stack, starting from left to right.
func (s *Stack[T]) Push(values ...T) {
	s.list.Append(values...)
}

// Pop get the value of the last push, and remove the value from the stack.
// If the stack is empty return [ErrIsEmpty] as an error.
func (s *Stack[T]) Pop() (T, error) {
	if value, err := s.list.Pop(); err != nil {
		return value, fmt.Errorf("failed to pop from stack, cause by %w", err)
	} else {
		return value, nil
	}
}

// Top get the value of the last push but does not remove the value from the stack.
// If the stack is empty return [ErrIsEmpty] as an error.
func (s *Stack[T]) Top() (T, error) {
	if value, err := s.list.Index(s.list.Length() - 1); err != nil {
		return value, fmt.Errorf("failed to peek at stack, cause by %w", err)
	} else {
		return value, nil
	}
}
