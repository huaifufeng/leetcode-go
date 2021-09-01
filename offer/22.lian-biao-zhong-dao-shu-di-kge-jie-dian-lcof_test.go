package offer

import (
	"testing"
)

//测试方法
func TestGetKthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	node := getKthFromEnd(head, 2)
	if node.Val != 4 {
		t.Error("链表中倒数第k个节点算法测试1错误")
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	node = getKthFromEnd1(head, 2)
	if node.Val != 4 {
		t.Error("链表中倒数第k个节点算法1测试1错误")
	}
}
