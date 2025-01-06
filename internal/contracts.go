package internal

// This file contains interfaces to guarantee backward compatibility of the package.
// We should only add to this file and not change or delete in it.
import "iter"

type List[T any] interface {
	Length() int
	Append(values ...T)
	Prepend(values ...T)
	Insert(value T, after int) error
	All() iter.Seq2[int, T]
	Backward() iter.Seq2[int, T]
	Search(target T, equal func(value, target T) bool) (int, error)
	Index(at int) (T, error)
	Pop() (T, error)
	Dequeue() (T, error)
	Remove(at int) (T, error)
}

type Stack[T any] interface {
	Length() int
	Push(values ...T)
	Pop() (T, error)
	Top() (T, error)
}

type Queue[T any] interface {
	Length() int
	Push(values ...T)
	Dequeue() (T, error)
	Front() (T, error)
	Rear() (T, error)
}

type Heap[T any] interface {
	Push(values ...T)
	Pop() (T, error)
	PushPop(value T) (T, error)
	Top() (T, error)
	IsEmpty() bool
}

type Cache[K comparable, T any] interface {
	Put(key K, value T) error
	Get(key K) (T, error)
}
