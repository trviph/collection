package cache

// A cache entry consist of a key of hashable type and a value of any type
type entry[K comparable, T any] struct {
	key   K
	value T
}
