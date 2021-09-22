/**
  题目地址：https://leetcode-cn.com/problems/split-linked-list-in-parts/

  解题：
  1、循环链表获取链表的长度，长度除以k获取每个子数组的长度，在取余k，获取多余的长度，从前向后依次添加一个值，获取整个数组的长度。 O(n)
    注意处理链表的索引
*/
package questionBank

func splitListToParts(head *ListNode, k int) []*ListNode {
	num := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		num++
	}

	if num < k {
		res := make([]*ListNode, k)
		cur := head
		for i := 0; i < num; i++ {
			res[i] = &ListNode{
				Val: cur.Val,
			}
			cur = cur.Next
		}

		return res
	}

	left := num / k
	reminder := num % k
	res := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		cur := &ListNode{
			Val: head.Val,
		}
		head = head.Next
		res[i] = cur
		for j := 1; j < left; j++ {
			cur.Next = &ListNode{
				Val: head.Val,
			}
			head = head.Next
			cur = cur.Next
		}

		if reminder > 0 {
			cur.Next = &ListNode{
				Val: head.Val,
			}
			head = head.Next
			reminder--
		}
	}

	return res
}
