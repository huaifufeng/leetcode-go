/**
  题目地址：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

  给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

  <pre>
    输入: 1->1->2
    输出: 1->2
  </pre>

  解题：
	1、循环这个链表，判断当前节点和下个节点的值是否相等，如果相等将当前节点指向下下个节点；不相等继续下个节点的判断 O(n)
*/
package questionBank

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	tmp := head

	for tmp != nil && tmp.Next != nil {
		//两个值相等说明重复，将指针直接指向下下个节点，否则进行下个节点判断
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return head
}

func DeleteDuplicates(head *ListNode) *ListNode  {
	return deleteDuplicates(head)
}