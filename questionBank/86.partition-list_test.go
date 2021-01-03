package questionBank

import (
	"testing"
)

func TestPartitionList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}

	res := partition(head, 3)

	rightRes := []int{1, 2, 2, 4, 3, 5}
	index := 0
	for res != nil {
		if res.Val != rightRes[index] {
			t.Error("分隔链表出现错误")
		}

		index++
		res = res.Next
	}
}
