package collection_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/trviph/collection"
)

func TestNewList(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestNewList", want[idx], got)
		}
	}
}

func TestListLength(t *testing.T) {
	emptyList := collection.NewList[int]()
	emptyListLen := emptyList.Length()
	zerolen := 0
	if emptyListLen != zerolen {
		t.Errorf(testFailedMsg, "TestListLength", zerolen, emptyListLen)
	}

	variableList := collection.NewList(1, 2, 3)
	variableListLen := variableList.Length()
	want := 3
	if variableListLen != want {
		t.Errorf(testFailedMsg, "TestListLength", want, variableListLen)
	}
}

func TestListAppend(t *testing.T) {
	list := collection.NewList[int]()
	list.Append()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4, 5, 6)
	want := []int{1, 2, 3, 4, 5, 6}

	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListAppend", want[idx], got)
		}
	}
}

func TestListPrepend(t *testing.T) {
	list := collection.NewList[int]()
	list.Prepend()
	list.Prepend(6)
	list.Prepend(5)
	list.Prepend(4)
	list.Prepend(3, 2, 1)
	want := []int{1, 2, 3, 4, 5, 6}

	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListPrepend", want[idx], got)
		}
	}
}

func TestListInsert(t *testing.T) {
	list := collection.NewList[int](1, 3, 6)

	// This make the list [1, 2, 3, 6]
	value, at := 2, 0
	if err := list.Insert(value, at); err != nil {
		t.Errorf(testFailedMsg, "TestListInsert", "nil error", err)
	}

	// This make the list [1, 2, 3, 4, 6]
	value, at = 4, 2
	if err := list.Insert(value, at); err != nil {
		t.Errorf(testFailedMsg, "TestListInsert", "nil error", err)
	}

	// This make the list [1, 2, 3, 4, 5, 6]
	value, at = 5, 3
	if err := list.Insert(value, at); err != nil {
		t.Errorf(testFailedMsg, "TestListInsert", "nil error", err)
	}

	// This should failed, because index is out of range
	value, at = 99, 99
	err := list.Insert(value, at)
	if _, ok := err.(*collection.ErrIndexOutOfRange); !ok {
		var wantErr error = &collection.ErrIndexOutOfRange{}
		t.Errorf(testFailedMsg, "TestListInsert", wantErr, err)
	}

	// This should failed, because index is negative
	value, at = -1, -1
	err = list.Insert(value, at)
	if _, ok := err.(*collection.ErrIndexOutOfRange); !ok {
		var wantErr error = &collection.ErrIndexOutOfRange{}
		t.Errorf(testFailedMsg, "TestListInsert", wantErr, err)
	}

	want := []int{1, 2, 3, 4, 5, 6}
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListPrepend", want[idx], got)
		}
	}
}

func TestListAll(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}

	prevIDX := -1
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListAll", want[idx], got)
		}
		if prevIDX >= idx {
			t.Errorf(
				"TestListAll failed; index should increasing; previous index is %d but current index is %d",
				prevIDX,
				idx,
			)
		}
		prevIDX = idx
	}

	// Test break early
	for idx := range list.All() {
		if idx > 2 {
			break
		}
	}
}

// Testing case of multiple goroutines access the list iterator
func TestListAllConcurrent(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	tf := func(id int) {
		defer wg.Done()
		for idx, got := range list.All() {
			t.Logf("[gorountine %d] value of %d at index %d", id, idx, got)
			if want[idx] != got {
				t.Errorf(testFailedMsg, "TestListAll", want[idx], got)
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

func TestListBackward(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}

	prevIDX := list.Length()
	for idx, got := range list.Backward() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListAll", want[idx], got)
		}
		if prevIDX <= idx {
			t.Errorf(
				"TestListAll failed; index should decreasing; previous index is %d but current index is %d",
				prevIDX,
				idx,
			)
		}
		prevIDX = idx
	}

	// Test break early
	for idx := range list.Backward() {
		if idx > 2 {
			break
		}
	}
}

// Testing case of multiple goroutines access the list iterator
func TestListBackwardConcurrent(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	tf := func(id int) {
		defer wg.Done()
		for idx, got := range list.Backward() {
			t.Logf("[gorountine %d] value of %d at index %d", id, idx, got)
			if want[idx] != got {
				t.Errorf(testFailedMsg, "TestListAll", want[idx], got)
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

func TestListSearch(t *testing.T) {
	list := collection.NewList(1, 2, 3, 5)
	equal := func(value, target int) bool {
		return value == target
	}

	gotIDX, gotErr := list.Search(4, equal)
	wantIDX := -1
	wantErr := &collection.ErrNotFound{}
	if gotIDX != wantIDX {
		t.Errorf(testFailedMsg, "TestListSearch", wantIDX, gotIDX)
	}
	if _, ok := gotErr.(*collection.ErrNotFound); !ok {
		t.Errorf(testFailedMsg, "TestListSearch", wantErr, gotErr)
	}

	gotIDX, gotErr = list.Search(2, equal)
	wantIDX = 1
	wantErr = nil
	if gotIDX != wantIDX {
		t.Errorf(testFailedMsg, "TestListSearch", wantIDX, gotIDX)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListSearch", wantErr, gotErr)
	}
}

func BenchmarkAppend(b *testing.B) {
	list := collection.NewList[int]()
	for i := 0; i <= b.N; i++ {
		list.Append(i)
	}
}
