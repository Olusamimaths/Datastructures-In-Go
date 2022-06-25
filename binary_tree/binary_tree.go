package binarytree

type BinaryTree[T comparable] interface {
	Size() int
	Add(item T)
	Contains(item T) bool
	Delete(item T)
	Unlink(currentNode *node[T], newNode *node[T])
	GetNode(item T) *node[T]
	Insert(parent *node[T], child *node[T])
}

type binarytree[T comparable] struct {
	size int
	root Node[T]
}

func NewBinaryTree[T comparable](item T) BinaryTree[T] {
	return &binarytree[T]{
		size: 0,
		root: NewNode(item),
	}
}

func (bt *binarytree[T]) Size() int {
	return bt.size
}

func (bt *binarytree[T]) Add(item T) {
	panic("not implemented") // TODO: Implement
}

func (bt *binarytree[T]) Contains(item T) bool {
	panic("not implemented") // TODO: Implement
}

func (bt *binarytree[T]) Delete(item T) {
	panic("not implemented") // TODO: Implement
}

func (bt *binarytree[T]) Unlink(currentNode *node[T], newNode *node[T]) {
	panic("not implemented") // TODO: Implement
}

func (bt *binarytree[T]) GetNode(item T) *node[T] {
	panic("not implemented") // TODO: Implement
}

func (bt *binarytree[T]) Insert(parent *node[T], child *node[T]) {
	panic("not implemented") // TODO: Implement
}
