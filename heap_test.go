package collection_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/trviph/collection"
)

func TestNewHeap(t *testing.T) {
	_, err := collection.NewHeap[any](nil)
	if err == nil {
		t.Errorf(testFailedMsg, "TestNewHeap", "error", nil)
	}

	// This create a min heap
	_, err = collection.NewHeap[int](collection.LessThan)
	if err != nil {
		t.Errorf(testFailedMsg, "TestNewHeap", "nil error", err)
	}

	// This create a min heap
	_, err = collection.NewHeap[int](collection.LessThanOrEqual)
	if err != nil {
		t.Errorf(testFailedMsg, "TestNewHeap", "nil error", err)
	}

	// This create a max heap
	_, err = collection.NewHeap[int](collection.GreaterThan)
	if err != nil {
		t.Errorf(testFailedMsg, "TestNewHeap", "nil error", err)
	}

	// This create a max heap
	_, err = collection.NewHeap[int](collection.GreaterThanOrEqual)
	if err != nil {
		t.Errorf(testFailedMsg, "TestNewHeap", "nil error", err)
	}
}

func TestMustNewHeap(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(testFailedMsg, "TestMustNewHeap", "panic", r)
		}
	}()
	_ = collection.MustNewHeap[any](nil)
}

func TestMinHeap(t *testing.T) {
	heap := collection.MustNewHeap[int](collection.LessThan)
	err := minHeapTest(heap)
	if err != nil {
		t.Errorf(testFailedMsg, "TestMinHeap", "nil error", err)
	}

	heap = collection.MustNewHeap[int](collection.LessThanOrEqual)
	err = minHeapTest(heap)
	if err != nil {
		t.Errorf(testFailedMsg, "TestMinHeap", "nil error", err)
	}
}

func minHeapTest(heap *collection.Heap[int]) error {
	pushed := 0
	for i := 0; i < rand.Intn(999)+1; i++ {
		heap.Push(rand.Int())
		pushed++
	}

	// Test if the heap always return min value
	previous, err := heap.Pop()
	if err != nil {
		return err
	}
	received := 1
	for !heap.IsEmpty() {
		current, err := heap.Pop()
		if err != nil {
			return err
		}
		if current < previous {
			return fmt.Errorf(
				"min heap return invalid value; current value %d is less than previous value %d",
				current,
				previous,
			)
		}
		previous = current
		received++
	}

	// Test if the heap return all the push value
	if pushed != received {
		return fmt.Errorf("pushed %d values but only received %d value, missing %d values", pushed, received, pushed-received)
	}

	return nil
}

func TestMaxHeap(t *testing.T) {
	heap := collection.MustNewHeap[int](collection.GreaterThan)
	err := maxHeapTest(heap)
	if err != nil {
		t.Errorf(testFailedMsg, "TestMaxHeap", "nil error", err)
	}

	heap = collection.MustNewHeap[int](collection.GreaterThanOrEqual)
	err = maxHeapTest(heap)
	if err != nil {
		t.Errorf(testFailedMsg, "TestMaxHeap", "nil error", err)
	}
}

func maxHeapTest(heap *collection.Heap[int]) error {
	pushed := 0
	for i := 0; i < rand.Intn(999)+1; i++ {
		heap.Push(rand.Int())
		pushed++
	}

	// Test if the heap always return min value
	previous, err := heap.Pop()
	if err != nil {
		return err
	}
	received := 1
	for !heap.IsEmpty() {
		current, err := heap.Pop()
		if err != nil {
			return err
		}
		if current > previous {
			return fmt.Errorf(
				"min heap return invalid value; current value %d is greater than previous value %d",
				current,
				previous,
			)
		}
		previous = current
		received++
	}

	// Test if the heap return all the push value
	if pushed != received {
		return fmt.Errorf("pushed %d values but only received %d value, missing %d values", pushed, received, pushed-received)
	}

	return nil
}

func TestHeapPushPop(t *testing.T) {
	maxHeap := collection.MustNewHeap[int](collection.GreaterThan)

	// Should return error since the heap is empty
	if _, err := maxHeap.PushPop(1); err == nil {
		t.Errorf(testFailedMsg, "TestHeapPushPop", "error", err)
	}

	maxHeap.Push(1)
	// Should return 2 since it is greater than what in the heap
	want := 2
	if got, err := maxHeap.PushPop(want); err != nil {
		t.Errorf(testFailedMsg, "TestHeapPushPop", "nil error", err)
	} else if got != want {
		t.Errorf(testFailedMsg, "TestHeapPushPop", want, got)
	}

	// Should return 1 since it is greater than the pushed value
	want = 1
	if got, err := maxHeap.PushPop(-1); err != nil {
		t.Errorf(testFailedMsg, "TestHeapPushPop", "nil error", err)
	} else if got != want {
		t.Errorf(testFailedMsg, "TestHeapPushPop", want, got)
	}

	// The next value should now be -1
	want = -1
	if got, err := maxHeap.Top(); err != nil {
		t.Errorf(testFailedMsg, "TestHeapPushPop", "nil error", err)
	} else if got != want {
		t.Errorf(testFailedMsg, "TestHeapPushPop", want, got)
	}
}

func TestHeapTop(t *testing.T) {
	maxHeap := collection.MustNewHeap[int](collection.GreaterThan)

	// Should return error since heap is empty
	if _, err := maxHeap.Top(); err == nil {
		t.Errorf(testFailedMsg, "TestHeapTop", "error", err)
	}

	maxHeap.Push(100)
	want := 100
	for i := 0; i < 10; i++ {
		// The value at top should not change
		if got, err := maxHeap.Top(); err != nil {
			t.Errorf(testFailedMsg, "TestHeapTop", "nil error", err)
		} else if got != want {
			t.Errorf(testFailedMsg, "TestHeapTop", want, got)
		}
	}
}
