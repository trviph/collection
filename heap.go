package collection

// A [Heap] implemented using [slice] (dynamic array) as the base.
type Heap[T any] struct {
	values []T
	cmp    func(T, T) bool
}

var _ heap[any] = (*Heap[any])(nil)

// [NewHeap] creates a new [Heap].
func NewHeap[T any](cmp func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		values: make([]T, 0), cmp: cmp,
	}
}

func (h *Heap[T]) Push(values ...T) {}

func (h *Heap[T]) Pop() (T, error) {
	var res T
	return res, nil
}

func (h *Heap[T]) PushPop(value T) (T, error) {
	var res T
	return res, nil
}

func (h *Heap[T]) Top() (T, error) {
	var res T
	return res, nil
}
