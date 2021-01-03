/**
  题目地址：https://leetcode-cn.com/problems/partition-list/

  给你一个链表和一个特定值 x ，请你对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。
  你应当保留两个分区中每个节点的初始相对位置。

  <pre>
     输入：head = 1->4->3->2->5->2, x = 3
	 输出：1->2->2->4->3->5
  </pre>

  解题：
	1、
*/
package questionBank

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	smallList := &ListNode{}
	bigList := &ListNode{}

	smallHead := smallList
	bigHead := bigList

	for head != nil {
		if head.Val < x {
			smallList.Next = head
			smallList = smallList.Next
		} else {
			bigList.Next = head
			bigList = bigList.Next
		}

		head = head.Next
	}

	bigList.Next = nil
	smallList.Next = bigHead.Next

	return smallHead.Next
}
