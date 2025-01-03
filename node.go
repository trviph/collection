// Collection is a Go library aim to implement some basic data structure such as List, Queue, Stack, Heap and more.
package collection

type node[T any] struct {
	value T
	// Left node or previous node
	left *node[T]
	// Right node or next node
	right *node[T]
}

func (n *node[T]) Val() T {
	return n.value
}
