/**
  题目地址：https://leetcode-cn.com/problems/reverse-linked-list/

  反转一个单链表。

  <pre>
    输入: 1->2->3->4->5->NULL
    输出: 5->4->3->2->1->NULL
  </pre>

  你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

  解题： 时间复杂度O(n) n为数组长度

*/
package questionBank

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//这里使用两个指针来完成操作 pre代表排好序的链表指针 cur代表当前处理的指针
//每次处理的时候，把当前指针的下一个节点指向 pre
//把下一个节点指向当前节点
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return pre
}
