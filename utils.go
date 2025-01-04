package collection

// Orderable are types that support the following operations
//   - Equal
//   - Not equal
//   - Less than
//   - Less than or equal
//   - Greater than
//   - Greater than or equal
type Orderable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string
}

// Less than, can be use as cmp function of [Heap] to create a min heap.
func LessThan[T Orderable](current, other T) bool {
	return current < other
}

// Less than or equal, can be use as cmp function of [Heap] to create a min heap.
func LessThanOrEqual[T Orderable](current, other T) bool {
	return current <= other
}

// Greater than, can be use as cmp function of [Heap] to create a max heap.
func GreaterThan[T Orderable](current, other T) bool {
	return current > other
}

// Greater than or equal, can be use as cmp function of [Heap] to create a max heap.
func GreaterThanOrEqual[T Orderable](current, other T) bool {
	return current >= other
}

// Ensure f must return a nil error else this will panic.
func Must[T any](f func() (T, error)) T {
	val, err := f()
	if err != nil {
		panic(err)
	}
	return val
}
