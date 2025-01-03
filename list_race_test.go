package collection_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/trviph/collection"
)

func TestListAllRace(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	tf := func(id int) {
		defer wg.Done()
		for idx, got := range list.All() {
			t.Logf("[gorountine %d] value of %d at index %d", id, idx, got)
			if want[idx] != got {
				t.Errorf(testFailedMsg, "TestListAllRace", want[idx], got)
			}
		}
	}

	// When using go test -v -race, the output goroutine id should be continously like the below,
	// notice how the goroutine id is not mixing-up in between run:
	//
	//  === RUN   TestListAll
	//  list_test.go:62: [gorountine 3] 0 - 1
	//  list_test.go:62: [gorountine 3] 1 - 2
	//  list_test.go:62: [gorountine 3] 2 - 3
	//  list_test.go:62: [gorountine 3] 3 - 4
	//  list_test.go:62: [gorountine 3] 4 - 5
	//  list_test.go:62: [gorountine 1] 0 - 1
	//  list_test.go:62: [gorountine 1] 1 - 2
	//  list_test.go:62: [gorountine 1] 2 - 3
	//  list_test.go:62: [gorountine 1] 3 - 4
	//  list_test.go:62: [gorountine 1] 4 - 5
	//  list_test.go:62: [gorountine 2] 0 - 1
	//  list_test.go:62: [gorountine 2] 1 - 2
	//  list_test.go:62: [gorountine 2] 2 - 3
	//  list_test.go:62: [gorountine 2] 3 - 4
	//  list_test.go:62: [gorountine 2] 4 - 5
	amount := rand.Intn(18) + 2
	wg.Add(amount)
	for i := 1; i <= amount; i++ {
		go tf(i)
	}
	wg.Wait()
}

func TestListBackwardRace(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	tf := func(id int) {
		defer wg.Done()
		for idx, got := range list.Backward() {
			t.Logf("[gorountine %d] value of %d at index %d", id, idx, got)
			if want[idx] != got {
				t.Errorf(testFailedMsg, "TestListBackwardRace", want[idx], got)
			}
		}
	}

	// When using go test -v -race, the output goroutine id should be continously like the below,
	// notice how the goroutine id is not mixing-up in between run:
	//
	//  === RUN   TestListAll
	//  list_test.go:62: [gorountine 3] 0 - 1
	//  list_test.go:62: [gorountine 3] 1 - 2
	//  list_test.go:62: [gorountine 3] 2 - 3
	//  list_test.go:62: [gorountine 3] 3 - 4
	//  list_test.go:62: [gorountine 3] 4 - 5
	//  list_test.go:62: [gorountine 1] 0 - 1
	//  list_test.go:62: [gorountine 1] 1 - 2
	//  list_test.go:62: [gorountine 1] 2 - 3
	//  list_test.go:62: [gorountine 1] 3 - 4
	//  list_test.go:62: [gorountine 1] 4 - 5
	//  list_test.go:62: [gorountine 2] 0 - 1
	//  list_test.go:62: [gorountine 2] 1 - 2
	//  list_test.go:62: [gorountine 2] 2 - 3
	//  list_test.go:62: [gorountine 2] 3 - 4
	//  list_test.go:62: [gorountine 2] 4 - 5
	amount := rand.Intn(18) + 2
	wg.Add(amount)
	for i := 1; i <= amount; i++ {
		go tf(i)
	}
	wg.Wait()
}

func TestListRace(t *testing.T) {
	list := collection.NewList(1, 2, 3)
	go list.Length()
	go list.Append(4, 5, 6)
	go list.Prepend(0, -1, -2)
	go func() {
		_ = list.Insert(-1, 0)
	}()
	go func() {
		_, _ = list.Index(0)
	}()
	go func() {
		for range list.All() {
		}
	}()
	go func() {
		for range list.Backward() {
		}
	}()
	go func() {
		_, _ = list.Search(1, func(value, target int) bool { return value == target })
	}()
	go func() {
		_, _ = list.Index(1)
	}()
	go func() {
		_, _ = list.Pop()
	}()
	go func() {
		_, _ = list.Dequeue()
	}()
	go func() {
		_, _ = list.Remove(0)
	}()
}
