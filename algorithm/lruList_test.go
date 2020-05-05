package algorithm

import (
	"fmt"
	"testing"
)

func TestLruList(t *testing.T) {
	lru := NewLruList(5)
	lru.getItem(1)
	lru.getItem(2)
	lru.getItem(3)

	lru.getItem(2)
	head := lru.head
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
