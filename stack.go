package collection

// A first-in-last-out [Stack] implemented by using [List] as the base.
// Since [List] is thread-safe [Stack] should also be thread-safe.
type Stack[T any] struct {
	list *List[T]
}

// Interface guard
var _ stack[any] = (*Stack[any])(nil)

// [NewStack] creates a new [Stack] of type T.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list: NewList[T]()}
}

// Length returns the number of values current in the stack.
func (s *Stack[T]) Length() int {
	return s.list.length
}

// Push a list of values in to the stack, starting from left to right.
func (s *Stack[T]) Push(values ...T)

// Pop get the value of the last push, and remove the value from the stack.
func (s *Stack[T]) Pop() (T, error)

// Top get the value of the last push but does not remove the value from the stack.
func (s *Stack[T]) Top() (T, error)
