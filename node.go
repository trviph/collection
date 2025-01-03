// Collection is a Go library aim to implement some basic data structure such as List, Queue, Stack, Heap and more.
package collection

type Node[T any] struct {
	data T
	// Left node or previous node
	left *Node[T]
	// Right node or next node
	right *Node[T]
}

func (n *Node[T]) Val() T {
	return n.data
}
