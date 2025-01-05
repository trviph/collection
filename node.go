// Collection is a Go library that aims to implement basic data structures such as List, Queue, Stack, Heap, and more.
package collection

type node[T any] struct {
	value T
	// Left node or previous node
	left *node[T]
	// Right node or next node
	right *node[T]
}

func (n *node[T]) unlink() {
	if n.left != nil && n.left.right == n {
		n.left.right = n.right
	}
	if n.right != nil && n.right.left == n {
		n.right.left = n.left
	}
	n.left = nil
	n.right = nil
}

func (n *node[T]) insert(newRight *node[T]) {
	newRight.left = n
	if n.right != nil && n.right.left == n {
		newRight.right = n.right
		n.right.left = newRight
	}
	n.right = newRight
}
