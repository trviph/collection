package cache

import (
	"github.com/trviph/collection"
	"github.com/trviph/collection/internal"
)

// A cache using the Least Recently Used eviction policy.
//
// [Least Recently Used (LRU)]: https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_Recently_Used_(LRU)
type LRU[K comparable, T any] struct {
	cap int
}

var _ internal.Cache[int, any] = (*LRU[int, any])(nil)

// [NewLRU] creates a new cache with [LRU] eviction policy.
// It accepts cap as the only argument, specifying the maximum capacity of the cache.
func NewLRU[K comparable, T any](cap int) *LRU[K, T] {
	return &LRU[K, T]{
		cap: cap,
	}
}

// Put a new value with an associated key into the cache.
// If the key already existed then it returns an [collection.ErrAlreadyExisted] as the error.
func (c *LRU[K, T]) Put(key K, value T) error {
	return nil
}

// Get the value associated with the given key argument.
// If there is no such key returns [collection.ErrNotFound], or if the cache is empty then returns [collection.ErrIsEmpty].
func (C *LRU[K, T]) Get(key K) (T, error) {
	var zeroValue T
	return zeroValue, collection.ErrIsEmpty
}
