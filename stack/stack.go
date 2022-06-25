package stack

type BasicStack[T comparable] interface {
	Push(item T)
	Size() int
	Pop() T
	Contains(item T) bool
	Access(item T) T
}

func NewStack[T comparable](Size int) BasicStack[T] {
	return &basicStack[T]{
		data:         make([]T, Size),
		stackPointer: 0,
	}
}

type basicStack[T comparable] struct {
	data         []T
	stackPointer int
}

func (bs *basicStack[T]) Push(x T) {
	bs.data[bs.stackPointer] = x
	bs.stackPointer++
}

func (bs *basicStack[T]) Size() int {
	return bs.stackPointer
}

func (bs *basicStack[T]) Pop() T {
	if bs.stackPointer == 0 {
		panic("No more items on the stack")
	}
	bs.stackPointer--
	return bs.data[bs.stackPointer]
}

func (bs *basicStack[T]) Contains(item T) bool {
	found := false

	if bs.Size() == 0 {
		return found
	}

	for i := 0; i < bs.Size(); i++ {
		if item == bs.data[i] {
			found = true
			break
		}
	}

	return found
}

func (bs *basicStack[T]) Access(item T) T {
	for bs.stackPointer > 0 {
		x := bs.Pop()
		if x == item {
			return x
		}
	}
	panic("Item not found")
}
