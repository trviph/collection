package cache

import (
	"fmt"
	"sync"

	"github.com/trviph/collection"
	"github.com/trviph/collection/internal"
)

// A cache using the Least Recently Used eviction policy.
//
// [Least Recently Used (LRU)]: https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_Recently_Used_(LRU)
type LRU[K comparable, T any] struct {
	mu  sync.Mutex
	cap int

	// To keep track and quickly look up which node is holding an entry by its key.
	entries map[K]*internal.Node[*entry[K, T]]

	// Keeping track of the recency of entries.
	// Entries are ordered from most recently used to least recently used,
	// going from head to tail.
	recency *collection.List[*entry[K, T]]
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
		cap:     cap,
		entries: make(map[K]*internal.Node[*entry[K, T]]),
		recency: collection.NewList[*entry[K, T]](),
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
// This marks the key as recently used.
func (c *LRU[K, T]) Put(key K, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If key already existed
	if _, ok := c.entries[key]; ok {
		c.updateEntry(key, value)
	} else {
		c.newEntry(key, value)
	}
	c.tryDrop()
}

func (c *LRU[K, T]) tryDrop() {
	if c.cap >= c.recency.Length() {
		return
	}
	entry, err := c.recency.Pop()
	if err != nil {
		// This should never happend
		panic(
			fmt.Errorf("something went very wrong; cannot drop LRU entry of cache with capacity of %d, entries lenght %d", c.cap, c.recency.Length()),
		)
	}
	// Delete entry from lookup map
	delete(c.entries, entry.key)
}

func (c *LRU[K, T]) newEntry(key K, value T) {
	// Mark it as recently used
	c.recency.Prepend(&entry[K, T]{key: key, value: value})
	// Add new entry to map
	c.entries[key] = c.recency.Head()
}

func (c *LRU[K, T]) updateEntry(key K, value T) {
	// Get the node contains the entry
	oldNode := c.entries[key]

	// Push new value to the recency list
	c.recency.Prepend(&entry[K, T]{key: key, value: value})
	// Update the entries map
	c.entries[key] = c.recency.Head()

	// Remove the node from the recency list
	// If oldNode is tail then Pop, else just unlink
	if oldNode.Right == nil {
		_, _ = c.recency.Pop()
	} else {
		oldNode.Unlink()
	}
}

// Get the value associated with the given key argument.
// If there is no such key returns [collection.ErrNotFound],
// or if the cache is empty then returns [collection.ErrIsEmpty].
// This marks the key as recently used.
func (c *LRU[K, T]) Get(key K) (T, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroValue T
	if c.recency.Length() == 0 {
		return zeroValue, collection.ErrIsEmpty
	}
	node, ok := c.entries[key]
	if !ok {
		return zeroValue, collection.ErrNotFound
	}

	entry := node.Value
	return entry.value, nil
}
