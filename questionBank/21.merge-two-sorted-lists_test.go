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
	value := mergeTwoLists(l1, l2)
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
			t.Error("合并有序链表方法验证失败")
		}

		value = value.Next
		res = res.Next
	}
}

func TestMergeTwoLists2(t *testing.T) {
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
	value := mergeTwoLists2(l1, l2)
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
			t.Error("合并有序链表方法2验证失败")
		}

		value = value.Next
		res = res.Next
	}
}
