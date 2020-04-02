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
				Val: 2,
				Next: nil,
			},
		},
	}
	res := DeleteDuplicates(head)
	for res.Val != 1 || res.Next.Val != 2 {
		t.Log(res)
		t.Error("删除链表重复值错误")
	}

}