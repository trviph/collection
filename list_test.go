package collection_test

import (
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
	value, at = 99, list.Length()
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

func TestListSearch(t *testing.T) {
	list := collection.NewList(1, 2, 3, 5)
	equal := func(value, target int) bool {
		return value == target
	}

	gotIDX, gotErr := list.Search(4, equal)
	wantIDX := -1
	wantErr := error(&collection.ErrNotFound{})
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

func TestListIndex(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}

	for idx := 0; idx < list.Length(); idx++ {
		gotValue, gotErr := list.Index(idx)
		if want[idx] != gotValue {
			t.Errorf(testFailedMsg, "TestListIndex", want[idx], gotValue)
		}
		if gotErr != nil {
			t.Errorf(testFailedMsg, "TestListIndex", "nil error", gotErr)
		}
	}

	_, gotErr := list.Index(list.Length())
	wantErr := error(&collection.ErrIndexOutOfRange{})
	if _, ok := gotErr.(*collection.ErrIndexOutOfRange); !ok {
		t.Errorf(testFailedMsg, "TestListIndex", wantErr, gotErr)
	}
}

func TestListPop(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{5, 4, 3, 2, 1}

	for _, w := range want {
		gotValue, gotErr := list.Pop()
		if w != gotValue {
			t.Errorf(testFailedMsg, "TestListPop", w, gotValue)
		}
		if gotErr != nil {
			t.Errorf(testFailedMsg, "TestListPop", "nil error", gotErr)
		}
	}

	_, gotErr := list.Pop()
	wantErr := error(&collection.ErrIsEmpty{})
	if _, ok := gotErr.(*collection.ErrIsEmpty); !ok {
		t.Errorf(testFailedMsg, "TestListPop", wantErr, gotErr)
	}
}

func TestListDequeue(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)
	want := []int{1, 2, 3, 4, 5}

	for _, w := range want {
		gotValue, gotErr := list.Dequeue()
		if w != gotValue {
			t.Errorf(testFailedMsg, "TestListDequeue", w, gotValue)
		}
		if gotErr != nil {
			t.Errorf(testFailedMsg, "TestListDequeue", "nil error", gotErr)
		}
	}

	_, gotErr := list.Dequeue()
	wantErr := error(&collection.ErrIsEmpty{})
	if _, ok := gotErr.(*collection.ErrIsEmpty); !ok {
		t.Errorf(testFailedMsg, "TestListDequeue", wantErr, gotErr)
	}
}

func TestListRemove(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5)

	// This make the list become [1, 2, 4, 5]
	gotValue, gotErr := list.Remove(2)
	wantValue := 3
	if wantValue != gotValue {
		t.Errorf(testFailedMsg, "TestListRemove", wantValue, gotValue)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListRemove", "nil error", gotErr)
	}

	want := []int{1, 2, 4, 5}
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListRemove", want[idx], got)
		}
	}

	_, gotErr = list.Remove(list.Length())
	wantErr := error(&collection.ErrIndexOutOfRange{})
	if _, ok := gotErr.(*collection.ErrIndexOutOfRange); !ok {
		t.Errorf(testFailedMsg, "TestListRemove", wantErr, gotErr)
	}

	emptyList := collection.NewList[int]()
	_, gotErr = emptyList.Remove(emptyList.Length())
	wantErr = error(&collection.ErrIsEmpty{})
	if _, ok := gotErr.(*collection.ErrIsEmpty); !ok {
		t.Errorf(testFailedMsg, "TestListRemove", wantErr, gotErr)
	}
}
