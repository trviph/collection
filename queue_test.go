package collection_test

import (
	"errors"
	"testing"

	"github.com/trviph/collection"
)

func TestNewQueue(t *testing.T) {
	// New empty queue
	_ = collection.NewQueue[int]()
	// New queue with values
	_ = collection.NewQueue[int](1, 2, 3, 4, 5)
}

func TestQueueLength(t *testing.T) {
	emptyQueue := collection.NewQueue[int]()
	if emptyQueue.Length() != 0 {
		t.Errorf(testFailedMsg, "TestQueueLength", 0, emptyQueue.Length())
	}

	queue := collection.NewQueue(1, 2, 3)
	if queue.Length() != 3 {
		t.Errorf(testFailedMsg, "TestQueueLength", 3, queue.Length())
	}
}

func TestQueuePush(t *testing.T) {
	queue := collection.NewQueue[int]()

	queue.Push(1)
	value, err := queue.Rear()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueuePush", "nil error", err)
	}
	if value != 1 {
		t.Errorf(testFailedMsg, "TestQueuePush", 1, value)
	}

	queue.Push(2)
	value, err = queue.Rear()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueuePush", "nil error", err)
	}
	if value != 2 {
		t.Errorf(testFailedMsg, "TestQueuePush", 2, value)
	}
}

func TestQueueDequeue(t *testing.T) {
	queue := collection.NewQueue[int]()

	_, err := queue.Dequeue()
	if !errors.Is(err, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestQueuePop", collection.ErrIsEmpty, err)
	}

	queue.Push(1)
	queue.Push(2)
	queue.Push(3, 4)

	// Should got 1 since 1 is pushed in first
	value, err := queue.Dequeue()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueuePop", "nil error", err)
	}
	if value != 1 {
		t.Errorf(testFailedMsg, "TestQueuePop", 1, value)
	}

	// Should got 2 next
	value, err = queue.Dequeue()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueuePop", "nil error", err)
	}
	if value != 2 {
		t.Errorf(testFailedMsg, "TestQueuePop", 2, value)
	}

	// Should got 3 next
	value, err = queue.Dequeue()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueuePop", "nil error", err)
	}
	if value != 3 {
		t.Errorf(testFailedMsg, "TestQueuePop", 3, value)
	}

	// And finally 4
	value, err = queue.Dequeue()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueuePop", "nil error", err)
	}
	if value != 4 {
		t.Errorf(testFailedMsg, "TestQueuePop", 4, value)
	}
}

func TestQueueFront(t *testing.T) {
	queue := collection.NewQueue[int]()

	_, err := queue.Front()
	if !errors.Is(err, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestQueueFront", collection.ErrIsEmpty, err)
	}

	queue.Push(1)
	queue.Push(2)

	// Should got 1 since 1 is pushed in first
	value, err := queue.Front()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueueFront", "nil error", err)
	}
	if value != 1 {
		t.Errorf(testFailedMsg, "TestQueueFront", 1, value)
	}

	// Should got 1 since 1 is not removed from the queue
	value, err = queue.Front()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueueFront", "nil error", err)
	}
	if value != 1 {
		t.Errorf(testFailedMsg, "TestQueueFront", 1, value)
	}
}

func TestQueueRear(t *testing.T) {
	queue := collection.NewQueue[int]()

	_, err := queue.Rear()
	if !errors.Is(err, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestQueueRear", collection.ErrIsEmpty, err)
	}

	queue.Push(1)
	queue.Push(2)

	// Should got 2 since 2 is pushed in last
	value, err := queue.Rear()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueueRear", "nil error", err)
	}
	if value != 2 {
		t.Errorf(testFailedMsg, "TestQueueRear", 2, value)
	}

	queue.Push(3, 4, 5)
	// Should got 5 since 5 is pushed in last
	value, err = queue.Rear()
	if err != nil {
		t.Errorf(testFailedMsg, "TestQueueRear", "nil error", err)
	}
	if value != 5 {
		t.Errorf(testFailedMsg, "TestQueueRear", 5, value)
	}
}
