package questionBank

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	node := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	head := &ListNode{
		Val:  4,
		Next: node,
	}

	deleteNode(node)
	if head.Next.Val != 1 {
		t.Error("删除链表中的节点算法测试1错误")
	}

	node = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
		},
	}

	head = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: node,
		},
	}

	deleteNode(node)
	if head.Next.Val != 5 {
		t.Error("删除链表中的节点算法测试2错误")
	}

	node = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
		},
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: node,
		},
	}

	deleteNode(node)
	if head.Next.Val != 2 {
		t.Error("删除链表中的节点算法测试3错误")
	}

	head = &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},
	}

	deleteNode(head)
	if head.Val != 1 {
		t.Error("删除链表中的节点算法测试4错误")
	}

	head = &ListNode{
		Val: -3,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: -99,
			},
		},
	}

	deleteNode(head)
	if head.Val != 5 {
		t.Error("删除链表中的节点算法测试5错误")
	}
}
