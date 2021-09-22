package questionBank

import "testing"

func TestSplitListToParts(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	k := 5
	res := splitListToParts(head, k)
	if res[2].Val != 3 {
		t.Error("分隔链表算法测试1错误")
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
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 8,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 10,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	k = 3
	res = splitListToParts(head, k)
	if res[2].Val != 8 {
		t.Error("分隔链表算法测试2错误", res[2])
	}
}
