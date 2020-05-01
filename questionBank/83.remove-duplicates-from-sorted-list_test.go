package questionBank

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	res := deleteDuplicates(head)
	for res.Val != 1 || res.Next.Val != 2 {
		t.Error("删除链表重复值错误")
	}

}
