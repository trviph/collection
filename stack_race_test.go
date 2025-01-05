package collection_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/trviph/collection"
)

func TestStackRace(t *testing.T) {
	var wg sync.WaitGroup
	stack := collection.NewStack[int]()
	functions := []func(){
		// Push to the stack
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				stack.Push(rand.Int())
			}
		},

		// Pop from the stack
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = stack.Pop()
			}
		},

		// Peek on the stack
		func() {
			defer wg.Done()
			for i := 0; i < randint(10, 1000); i++ {
				_, _ = stack.Top()
			}
		},
	}

	wg.Add(len(functions))
	for _, f := range functions {
		go f()
	}
	wg.Wait()
}
