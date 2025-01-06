package cache

import (
	"fmt"
	"sync"

	"github.com/trviph/collection"
	"github.com/trviph/collection/internal"
)

// A cache using the Most Recently Used eviction policy.
//
// [Most Recently Used (MRU)]: https://en.wikipedia.org/wiki/Cache_replacement_policies#Most-recently-used_(MRU)
type MRU[K comparable, T any] struct {
	mu  sync.Mutex
	cap int

	// To keep track and quickly look up which node is holding an entry by its key.
	entryNodes map[K]*internal.Node[*entry[K, T]]

	// Keeping track of the recency of entries.
	// Entries are ordered from most recently used to least recently used,
	// going from head to tail.
	entryRecency *collection.List[*entry[K, T]]
}

var _ internal.Cache[int, any] = (*MRU[int, any])(nil)

// [NewLRU] creates a new cache with [LRU] eviction policy.
// It accepts cap as the only argument, specifying the maximum capacity of the cache.
// Return an error if cap is less than 1.
func NewMRU[K comparable, T any](cap int) (*MRU[K, T], error) {
	if cap < 1 {
		return nil, fmt.Errorf("failed to create mru; cause by invalid specified capacity")
	}
	return &MRU[K, T]{
		cap:          cap,
		entryNodes:   make(map[K]*internal.Node[*entry[K, T]]),
		entryRecency: collection.NewList[*entry[K, T]](),
	}, nil
}

// Like [NewMRU] but will panic on error.
func MustNewMRU[K comparable, T any](cap int) *MRU[K, T] {
	return collection.Must(
		func() (*MRU[K, T], error) {
			return NewMRU[K, T](cap)
		},
	)
}

// Put a new value with an associated key into the cache.
// Update the value if the key already exist.
// This marks the key as recently used.
func (c *MRU[K, T]) Put(key K, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If key already existed
	if _, ok := c.entryNodes[key]; ok {
		c.updateEntry(key, value)
	} else {
		c.evict()
		c.newEntry(key, value)
	}
}

func (c *MRU[K, T]) evict() {
	if c.cap > c.entryRecency.Length() {
		return
	}
	entry, err := c.entryRecency.Dequeue()
	if err != nil {
		// This should never happend
		panic(
			fmt.Errorf("something went very wrong; cannot drop LRU entry of cache with capacity of %d, entries length %d", c.cap, c.entryRecency.Length()),
		)
	}
	// Delete entry from lookup map
	delete(c.entryNodes, entry.key)
}

func (c *MRU[K, T]) newEntry(key K, value T) {
	// Mark it as recently used
	c.entryRecency.Prepend(&entry[K, T]{key: key, value: value})
	// Add new entry to map
	c.entryNodes[key] = c.entryRecency.Head()
}

func (c *MRU[K, T]) updateEntry(key K, value T) {
	// Get the node contains the entry
	oldNode := c.entryNodes[key]

	// Push new value to the recency list
	c.entryRecency.Prepend(&entry[K, T]{key: key, value: value})
	// Update the entries map
	c.entryNodes[key] = c.entryRecency.Head()

	// Remove the node from the recency list
	// If oldNode is tail then Pop, else just unlink
	if oldNode.Right == nil {
		_, _ = c.entryRecency.Pop()
	} else {
		oldNode.Unlink()
	}
}

// Get the value associated with the given key argument.
// If there is no such key returns [collection.ErrNotFound],
// or if the cache is empty then returns [collection.ErrIsEmpty].
// This marks the key as recently used.
func (c *MRU[K, T]) Get(key K) (T, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroValue T
	if c.entryRecency.Length() == 0 {
		return zeroValue, collection.ErrIsEmpty
	}
	node, ok := c.entryNodes[key]
	if !ok {
		return zeroValue, collection.ErrNotFound
	}

	entry := node.Value
	return entry.value, nil
}
