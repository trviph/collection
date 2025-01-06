package cache

// A cache entry consist of key and value
type entry[K comparable, T any] struct {
	key   K
	value T
}
