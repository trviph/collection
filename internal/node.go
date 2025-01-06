// Collection is a Go library that aims to implement basic data structures such as List, Queue, Stack, Heap, and more.
package internal

type Node[T any] struct {
	Value T
	// Left node or previous node
	Left *Node[T]
	// Right node or next node
	Right *Node[T]
}

func (n *Node[T]) Unlink() {
	if n.Left != nil && n.Left.Right == n {
		n.Left.Right = n.Right
	}
	if n.Right != nil && n.Right.Left == n {
		n.Right.Left = n.Left
	}
	n.Left = nil
	n.Right = nil
}

func (n *Node[T]) Insert(newRight *Node[T]) {
	newRight.Left = n
	if n.Right != nil && n.Right.Left == n {
		newRight.Right = n.Right
		n.Right.Left = newRight
	}
	n.Right = newRight
}
