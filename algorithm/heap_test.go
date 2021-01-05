package algorithm

import (
	"testing"
)

func TestHeapInsert(t *testing.T) {
	heap := NewHeap(5)
	heap.insert(1)
	heap.insert(2)
	heap.insert(3)
	heap.insert(4)
	heap.insert(10)
	heap.insert(14)

	heap.delete()
	heap.delete()
	heap.delete()
	heap.delete()

	heap.insert(2)
	heap.insert(21)

}
