package collection_test

import (
	"errors"
	"testing"

	"github.com/trviph/collection"
)

func TestNewStack(t *testing.T) {
	// New empty stack
	_ = collection.NewStack[int]()
	// New stack with values
	_ = collection.NewStack[int](1, 2, 3, 4, 5)
}

func TestStackLength(t *testing.T) {
	emptyStack := collection.NewStack[int]()
	if emptyStack.Length() != 0 {
		t.Errorf(testFailedMsg, "TestStackLength", 0, emptyStack.Length())
	}

	stack := collection.NewStack(1, 2, 3)
	if stack.Length() != 3 {
		t.Errorf(testFailedMsg, "TestStackLength", 3, stack.Length())
	}
}

func TestStackPush(t *testing.T) {
	stack := collection.NewStack[int]()

	stack.Push(1)
	value, err := stack.Top()
	if err != nil {
		t.Errorf(testFailedMsg, "TestStackPush", "nil error", err)
	}
	if value != 1 {
		t.Errorf(testFailedMsg, "TestStackPush", 1, value)
	}

	stack.Push(2)
	value, err = stack.Top()
	if err != nil {
		t.Errorf(testFailedMsg, "TestStackPush", "nil error", err)
	}
	if value != 2 {
		t.Errorf(testFailedMsg, "TestStackPush", 2, value)
	}
}

func TestStackPop(t *testing.T) {
	stack := collection.NewStack[int]()

	_, err := stack.Pop()
	if !errors.Is(err, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestStackPop", collection.ErrIsEmpty, err)
	}

	stack.Push(1)
	stack.Push(2)

	// Should got 2 since 2 is pushed in last
	value, err := stack.Pop()
	if err != nil {
		t.Errorf(testFailedMsg, "TestStackPop", "nil error", err)
	}
	if value != 2 {
		t.Errorf(testFailedMsg, "TestStackPop", 2, value)
	}

	// Should got 1 since 2 is removed from the stack
	value, err = stack.Pop()
	if err != nil {
		t.Errorf(testFailedMsg, "TestStackPop", "nil error", err)
	}
	if value != 1 {
		t.Errorf(testFailedMsg, "TestStackPop", 1, value)
	}
}

func TestStackTop(t *testing.T) {
	stack := collection.NewStack[int]()

	_, err := stack.Top()
	if !errors.Is(err, collection.ErrIsEmpty) {
		t.Errorf(testFailedMsg, "TestStackTop", collection.ErrIsEmpty, err)
	}

	stack.Push(1)
	stack.Push(2)

	// Should got 2 since 2 is pushed in last
	value, err := stack.Top()
	if err != nil {
		t.Errorf(testFailedMsg, "TestStackTop", "nil error", err)
	}
	if value != 2 {
		t.Errorf(testFailedMsg, "TestStackTop", 2, value)
	}

	// Should got 2 since 2 is not removed from the stack
	value, err = stack.Top()
	if err != nil {
		t.Errorf(testFailedMsg, "TestStackTop", "nil error", err)
	}
	if value != 2 {
		t.Errorf(testFailedMsg, "TestStackTop", 2, value)
	}
}
