package questionBank

import "testing"

func TestSwapPairs(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	res := swapPairs(head)
	for res.Val != 2 || res.Next.Val != 1 || res.Next.Next.Val != 4 || res.Next.Next.Next.Val != 3 {
		t.Error("两两交换链表中节点方法错误")
	}
}

func TestSwapPairs2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	res := swapPairs2(head)
	for res.Val != 2 || res.Next.Val != 1 || res.Next.Next.Val != 4 || res.Next.Next.Next.Val != 3 {
		t.Error("两两交换链表中节点方法2错误")
	}
}
