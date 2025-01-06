package cache

import (
	"fmt"

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
// Return an error if cap is less than 1.
func NewLRU[K comparable, T any](cap int) (*LRU[K, T], error) {
	if cap < 1 {
		return nil, fmt.Errorf("failed to create lru; cause by invalid specified capacity")
	}
	return &LRU[K, T]{
		cap: cap,
	}, nil
}

// Like [NewLRU] but will panic on error.
func MustNewLRU[K comparable, T any](cap int) *LRU[K, T] {
	return collection.Must(
		func() (*LRU[K, T], error) {
			return NewLRU[K, T](cap)
		},
	)
}

// Put a new value with an associated key into the cache.
// Update the value if the key already exist.
// This mark the key as recently used.
func (c *LRU[K, T]) Put(key K, value T) {}

// Get the value associated with the given key argument.
// If there is no such key returns [collection.ErrNotFound],
// or if the cache is empty then returns [collection.ErrIsEmpty].
// This mark the key as recently used.
func (C *LRU[K, T]) Get(key K) (T, error) {
	var zeroValue T
	return zeroValue, collection.ErrIsEmpty
}
