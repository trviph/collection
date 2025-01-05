package collection_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/trviph/collection"
)

func TestQueueRace(t *testing.T) {
	var wg sync.WaitGroup
	queue := collection.NewQueue[int]()
	functions := []func(){
		// Push to the queue
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				queue.Push(rand.Int())
			}
		},

		// Dequeue from the queue
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = queue.Dequeue()
			}
		},

		// Peek on the queue
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = queue.Front()
			}
		},

		// Peek on the queue
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = queue.Rear()
			}
		},
	}

	wg.Add(len(functions))
	for _, f := range functions {
		go f()
	}
	wg.Wait()
}
