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
