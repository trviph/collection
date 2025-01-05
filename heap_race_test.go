package collection_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/trviph/collection"
)

func TestHeapRace(t *testing.T) {
	var wg sync.WaitGroup
	heap := collection.MustNewHeap[int](collection.GreaterThan)
	functions := []func(){
		// Push to the heap
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				heap.Push(rand.Int())
			}
		},

		// Pop from the heap
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = heap.Pop()
			}
		},

		// PushPop on the heap
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = heap.PushPop(rand.Int())
			}
		},

		// Peek at the heap
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = heap.Top()
			}
		},

		// Check if heap is empty
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_ = heap.IsEmpty()
			}
		},
	}

	wg.Add(len(functions))
	for _, f := range functions {
		go f()
	}
	wg.Wait()
}
