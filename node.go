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
