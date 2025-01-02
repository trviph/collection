package collection_test

import (
	"math/rand"
	"testing"

	"github.com/trviph/collection"
)

func TestAppend(t *testing.T) {
	list := collection.NewList[int]()
	for want := 1; want <= 10_000_000; want++ {
		list.Append(rand.Int())
		got := list.Length()
		if got != want {
			t.Fatalf("failed append to list length should be %d, got %d", want, got)
		}
	}
}

func TestPrepend(t *testing.T) {
	list := collection.NewList[int]()
	for want := 1; want <= 10_000_000; want++ {
		list.Prepend(rand.Int())
		got := list.Length()
		if got != want {
			t.Fatalf("failed prepend to list length should be %d, got %d", want, got)
		}
	}
}

func BenchmarkAppend(b *testing.B) {
	list := collection.NewList[int]()
	for i := 0; i <= b.N; i++ {
		list.Append(i)
	}
}
