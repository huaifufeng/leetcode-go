package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	head := &questionBank.ListNode{
		Val: 1,
		Next: &questionBank.ListNode{
			Val: 1,
			Next: &questionBank.ListNode{
				Val: 2,
				Next: nil,
			},
		},
	}
	res := questionBank.DeleteDuplicates(head)
	for res.Val != 1 || res.Next.Val != 2 {
		t.Log(res)
		t.Error("删除链表重复值错误")
	}

}