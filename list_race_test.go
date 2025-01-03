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
	var wg sync.WaitGroup
	list := collection.NewList(1, 2, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			list.Append(rand.Int())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			list.Prepend(rand.Int())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			_ = list.Insert(rand.Int(), rand.Intn(list.Length()+1))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			for range list.All() {
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			for range list.Backward() {
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			_, _ = list.Search(rand.Int(), func(value, target int) bool { return value == target })
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			_, _ = list.Index(rand.Intn(list.Length() + 1))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			_, _ = list.Pop()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			_, _ = list.Dequeue()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < rand.Intn(1000); i++ {
			_, _ = list.Remove(rand.Intn(list.Length() + 1))
		}
	}()

	wg.Wait()
}
