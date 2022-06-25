package binarytree

type Node[T comparable] interface {
	GetLeft() *node[T]
	SetLeft(left *node[T])

	GetRight() *node[T]
	SetRight(right *node[T])

	GetParent() *node[T]
	SetParent(parent *node[T])

	GetItem() T
	SetItem(item T)
}

type node[T comparable] struct {
	left   *node[T]
	right  *node[T]
	parent *node[T]

	item T
}

func NewNode[T comparable](item T) Node[T] {
	return &node[T]{}
}

func (n *node[T]) GetLeft() *node[T] {
	return n.left
}

func (n *node[T]) SetLeft(left *node[T]) {
	n.left = left
}

func (n *node[T]) GetRight() *node[T] {
	return n.right
}

func (n *node[T]) SetRight(right *node[T]) {
	n.right = right
}

func (n *node[T]) GetParent() *node[T] {
	return n.parent
}

func (n *node[T]) SetParent(parent *node[T]) {
	n.parent = parent
}

func (n *node[T]) GetItem() T {
	return n.item
}

func (n *node[T]) SetItem(item T) {
	n.item = item
}
