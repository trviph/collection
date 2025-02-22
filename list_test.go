package collection_test

import (
	"errors"
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

	// Test to see if nodes are linked properly
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListAppend", want[idx], got)
		}
	}

	// Test to see if nodes are linked properly
	for idx, got := range list.Backward() {
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

	// Test to see if nodes are linked properly
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListPrepend", want[idx], got)
		}
	}

	// Test to see if nodes are linked properly
	for idx, got := range list.Backward() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListPrepend", want[idx], got)
		}
	}
}

func TestListInsert(t *testing.T) {
	list := collection.NewList(1, 3, 6)

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

	// This make the list [1, 2, 3, 4, 5, 6, 7]
	value, at = 7, list.Length()-1
	if err := list.Insert(value, at); err != nil {
		t.Errorf(testFailedMsg, "TestListInsert", "nil error", err)
	}

	// This should failed, because index is out of range
	value, at = 99, list.Length()
	err := list.Insert(value, at)
	if !errors.Is(err, collection.ErrIndexOutOfRange) {
		t.Errorf(testFailedMsg, "TestListInsert", collection.ErrIndexOutOfRange, err)
	}

	// This should failed, because index is negative
	value, at = -1, -1
	err = list.Insert(value, at)
	if !errors.Is(err, collection.ErrIndexOutOfRange) {
		t.Errorf(testFailedMsg, "TestListInsert", collection.ErrIndexOutOfRange, err)
	}

	want := []int{1, 2, 3, 4, 5, 6, 7}
	// Test to see if nodes are linked properly
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListInsert", want[idx], got)
		}
	}

	// Test to see if nodes are linked properly
	for idx, got := range list.Backward() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListInsert", want[idx], got)
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
	if gotIDX != wantIDX {
		t.Errorf(testFailedMsg, "TestListSearch", wantIDX, gotIDX)
	}
	if !errors.Is(gotErr, collection.ErrNotFound) {
		t.Errorf(testFailedMsg, "TestListSearch", collection.ErrNotFound, gotErr)
	}

	gotIDX, gotErr = list.Search(2, equal)
	wantIDX = 1
	if gotIDX != wantIDX {
		t.Errorf(testFailedMsg, "TestListSearch", wantIDX, gotIDX)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListSearch", "nil error", gotErr)
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
	if !errors.Is(gotErr, collection.ErrIndexOutOfRange) {
		t.Errorf(testFailedMsg, "TestListIndex", collection.ErrIndexOutOfRange, gotErr)
	}
}

func TestListPop(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5, 6)

	// Pop a value
	wantPopValue := 6
	gotValue, gotErr := list.Pop()
	if wantPopValue != gotValue {
		t.Errorf(testFailedMsg, "TestListPop", wantPopValue, gotValue)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListPop", "nil error", gotErr)
	}

	// Test to see if nodes are linked properly
	want := []int{1, 2, 3, 4, 5}
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListPop", want[idx], got)
		}
	}

	// Test to see if nodes are linked properly
	for idx, got := range list.Backward() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListPop", want[idx], got)
		}
	}

	// Pop on empty list
	emptyList := collection.NewList[int]()
	_, gotErr = emptyList.Pop()
	if !errors.Is(gotErr, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestListPop", collection.ErrIsEmpty, gotErr)
	}
}

func TestListDequeue(t *testing.T) {
	list := collection.NewList(1, 2, 3, 4, 5, 6)

	// Dequeue a value
	wantDequeueValue := 1
	gotValue, gotErr := list.Dequeue()
	if wantDequeueValue != gotValue {
		t.Errorf(testFailedMsg, "TestListDequeue", wantDequeueValue, gotValue)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListDequeue", "nil error", gotErr)
	}

	// Test to see if nodes are linked properly
	want := []int{2, 3, 4, 5, 6}
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListDequeue", want[idx], got)
		}
	}

	// Test to see if nodes are linked properly
	for idx, got := range list.Backward() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListDequeue", want[idx], got)
		}
	}

	// Pop on empty list
	emptyList := collection.NewList[int]()
	_, gotErr = emptyList.Dequeue()
	if !errors.Is(gotErr, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestListDequeue", collection.ErrIsEmpty, gotErr)
	}
}

func TestListRemove(t *testing.T) {
	list := collection.NewList(1, 1, 2, 2, 2, 3, 4, 5, 5)
	// Use Remove(0) same effect as Dequeue(), this make the list become [1, 2, 2, 2, 3, 4, 5, 5]
	gotValue, gotErr := list.Remove(0)
	wantValue := 1
	if wantValue != gotValue {
		t.Errorf(testFailedMsg, "TestListRemove", wantValue, gotValue)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListRemove", "nil error", gotErr)
	}

	// Use Remove(length-1) same effect as Pop(), this make the list become [1, 2, 2, 2, 3, 4, 5]
	gotValue, gotErr = list.Remove(list.Length() - 1)
	wantValue = 5
	if wantValue != gotValue {
		t.Errorf(testFailedMsg, "TestListRemove", wantValue, gotValue)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListRemove", "nil error", gotErr)
	}

	// This make the list become [1, 2, 2, 3, 4, 5]
	gotValue, gotErr = list.Remove(1)
	wantValue = 2
	if wantValue != gotValue {
		t.Errorf(testFailedMsg, "TestListRemove", wantValue, gotValue)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListRemove", "nil error", gotErr)
	}

	// This make the list become [1, 2, 3, 4, 5]
	gotValue, gotErr = list.Remove(1)
	wantValue = 2
	if wantValue != gotValue {
		t.Errorf(testFailedMsg, "TestListRemove", wantValue, gotValue)
	}
	if gotErr != nil {
		t.Errorf(testFailedMsg, "TestListRemove", "nil error", gotErr)
	}

	// Test to see if nodes are linked properly
	want := []int{1, 2, 3, 4, 5}
	for idx, got := range list.All() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListRemove", want[idx], got)
		}
	}

	// Test to see if nodes are linked properly
	for idx, got := range list.Backward() {
		if want[idx] != got {
			t.Errorf(testFailedMsg, "TestListRemove", want[idx], got)
		}
	}

	// Remove with out-of-range index
	_, gotErr = list.Remove(list.Length())
	if !errors.Is(gotErr, collection.ErrIndexOutOfRange) {
		t.Errorf(testFailedMsg, "TestListRemove", collection.ErrIndexOutOfRange, gotErr)
	}

	// Remove on empty list
	emptyList := collection.NewList[int]()
	_, gotErr = emptyList.Remove(emptyList.Length())
	if !errors.Is(gotErr, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestListRemove", collection.ErrIsEmpty, gotErr)
	}
}
