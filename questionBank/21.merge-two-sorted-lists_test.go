package questionBank

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}
	value := MergeTwoLists(l1, l2)
	res := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
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
