package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &questionBank.ListNode{
		Val:1,
		Next: &questionBank.ListNode{
			Val:2,
			Next: &questionBank.ListNode{
				Val:4,
				Next:nil}}}
	l2 := &questionBank.ListNode{
		Val:1,
		Next: &questionBank.ListNode{
			Val:3,
			Next: &questionBank.ListNode{
				Val:4,
				Next:nil}}}
	value := questionBank.MergeTwoLists(l1, l2)
	res := &questionBank.ListNode{
		Val:1,
		Next: &questionBank.ListNode{
			Val:1,
			Next: &questionBank.ListNode{
				Val:2,
				Next:&questionBank.ListNode{
					Val:3,
					Next:&questionBank.ListNode{
						Val:4,
						Next:&questionBank.ListNode{
							Val:4,
							Next:nil,
						},
					},
				}}}}
	for value != nil && res != nil {
		if value.Val != res.Val {
			t.Log(value)
			t.Error("方法验证失败")
		}

		value = value.Next
		res = res.Next
	}
}