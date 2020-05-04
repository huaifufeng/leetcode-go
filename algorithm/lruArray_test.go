package algorithm

import "testing"

func TestLruArray(t *testing.T) {
	lru := NewLruArray(5)
	lru.getItem(1)
	lru.getItem(2)
	lru.getItem(3)
	lru.getItem(4)
	lru.getItem(5)

	lru.getItem(2)
	lru.getItem(6)
}
