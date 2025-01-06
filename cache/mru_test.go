package cache_test

import (
	"errors"
	"math/rand"
	"sync"
	"testing"

	"github.com/trviph/collection"
	"github.com/trviph/collection/cache"
)

func TestNewMRU(t *testing.T) {
	_, err := cache.NewMRU[int, int](0)
	if err == nil {
		t.Errorf(testFailedMsg, "TestNewMRU", "error", err)
	}
	_, err = cache.NewMRU[int, int](1)
	if err != nil {
		t.Errorf(testFailedMsg, "TestNewMRU", "nil error", err)
	}
}

func TestMustNewMRU(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(testFailedMsg, "TestMustNewMRU", "panic", r)
		}
	}()
	_ = cache.MustNewMRU[int, any](-1)
}

func TestMRU(t *testing.T) {
	// Create a cache the only hold 2 values at maximum
	MRU, err := cache.NewMRU[int, string](2)
	if err != nil {
		t.Errorf(testFailedMsg, "TestMRUPut", "nil error", err)
	}

	// Should get is empty error
	if _, err := MRU.Get(1); !errors.Is(err, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestMRUPut", collection.ErrIsEmpty, err)
	}

	// Basic case this should get A back
	MRU.Put(1, "A")
	if val, err := MRU.Get(1); err != nil {
		t.Errorf(testFailedMsg, "TestMRUPut", "nil error", err)
	} else if val != "A" {
		t.Errorf(testFailedMsg, "TestMRUPut", "A", val)
	}

	// Should get error since B is evicted
	MRU.Put(2, "B")
	MRU.Put(3, "C")
	if _, err := MRU.Get(2); !errors.Is(err, collection.ErrNotFound) {
		t.Errorf(testFailedMsg, "TestMRUPut", collection.ErrNotFound, err)
	}

	// Should update A to AA and mark it as recently used
	MRU.Put(1, "AA")
	if val, err := MRU.Get(1); err != nil {
		t.Errorf(testFailedMsg, "TestMRUPut", "nil error", err)
	} else if val != "AA" {
		t.Errorf(testFailedMsg, "TestMRUPut", "AA", val)
	}

	// Put B back in, this should cause AA to be evicted
	MRU.Put(2, "B")
	if _, err := MRU.Get(1); !errors.Is(err, collection.ErrNotFound) {
		t.Errorf(testFailedMsg, "TestMRUPut", collection.ErrNotFound, err)
	}

	if val, err := MRU.Get(2); err != nil {
		t.Errorf(testFailedMsg, "TestMRUPut", "nil error", err)
	} else if val != "B" {
		t.Errorf(testFailedMsg, "TestMRUPut", "B", val)
	}

	if val, err := MRU.Get(3); err != nil {
		t.Errorf(testFailedMsg, "TestMRUPut", "nil error", err)
	} else if val != "C" {
		t.Errorf(testFailedMsg, "TestMRUPut", "C", val)
	}
}

func TestMRURace(t *testing.T) {
	var wg sync.WaitGroup
	MRU := cache.MustNewMRU[int, int](randint(10, 50))
	functions := []func(){
		// Put to the cache
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				MRU.Put(randint(0, 100), rand.Int())
			}
		},

		// Get from the cache
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = MRU.Get(randint(0, 100))
			}
		},
	}

	wg.Add(len(functions))
	for _, f := range functions {
		go f()
	}
	wg.Wait()
}
