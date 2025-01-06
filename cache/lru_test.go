package cache_test

import (
	"errors"
	"math/rand"
	"sync"
	"testing"

	"github.com/trviph/collection"
	"github.com/trviph/collection/cache"
)

func TestNewLRU(t *testing.T) {
	_, err := cache.NewLRU[int, int](0)
	if err == nil {
		t.Errorf(testFailedMsg, "TestNewLRU", "error", err)
	}
	_, err = cache.NewLRU[int, int](1)
	if err != nil {
		t.Errorf(testFailedMsg, "TestNewLRU", "nil error", err)
	}
}

func TestMustNewLRU(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(testFailedMsg, "TestMustNewLRU", "panic", r)
		}
	}()
	_ = cache.MustNewLRU[int, any](-1)
}

func TestLRU(t *testing.T) {
	// Create a cache the only hold 2 values at maximum
	lru, err := cache.NewLRU[int, string](2)
	if err != nil {
		t.Errorf(testFailedMsg, "TestLRUPut", "nil error", err)
	}

	// Should get is empty error
	if _, err := lru.Get(1); !errors.Is(err, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestLRUPut", collection.ErrIsEmpty, err)
	}

	// Basic case this should get A back
	lru.Put(1, "A")
	if val, err := lru.Get(1); err != nil {
		t.Errorf(testFailedMsg, "TestLRUPut", "nil error", err)
	} else if val != "A" {
		t.Errorf(testFailedMsg, "TestLRUPut", "A", val)
	}

	// Should get error since A is evicted
	lru.Put(2, "B")
	lru.Put(3, "C")
	if _, err := lru.Get(1); !errors.Is(err, collection.ErrNotFound) {
		t.Errorf(testFailedMsg, "TestLRUPut", collection.ErrNotFound, err)
	}

	// Should update B to BB and mark it as recently used
	lru.Put(2, "BB")
	if val, err := lru.Get(2); err != nil {
		t.Errorf(testFailedMsg, "TestLRUPut", "nil error", err)
	} else if val != "BB" {
		t.Errorf(testFailedMsg, "TestLRUPut", "BB", val)
	}

	// Put A back in, this should cause C to be evicted
	lru.Put(1, "A")
	if _, err := lru.Get(3); !errors.Is(err, collection.ErrNotFound) {
		t.Errorf(testFailedMsg, "TestLRUPut", collection.ErrNotFound, err)
	}
	if val, err := lru.Get(1); err != nil {
		t.Errorf(testFailedMsg, "TestLRUPut", "nil error", err)
	} else if val != "A" {
		t.Errorf(testFailedMsg, "TestLRUPut", "A", val)
	}
}

func TestLRURace(t *testing.T) {
	var wg sync.WaitGroup
	lru := cache.MustNewLRU[int, int](randint(10, 50))
	functions := []func(){
		// Put to the cache
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				lru.Put(randint(0, 100), rand.Int())
			}
		},

		// Get from the cache
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = lru.Get(randint(0, 100))
			}
		},
	}

	wg.Add(len(functions))
	for _, f := range functions {
		go f()
	}
	wg.Wait()
}
