/**
  题目地址：https://leetcode-cn.com/problems/merge-two-sorted-lists/

  将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

  <pre>
     输入：1->2->4, 1->3->4
     输出：1->1->2->3->4->4
  </pre>

  解题： 时间复杂度O(n) n为最短的链表的元素个数。
	这里使用递归的方法，比较两个链表的值，将值小的链表赋值给目标链表l，然后继续递归处理值大的链表和值小链表的下一个链表值
*/
package questionBank

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//如果一个是nil，直接返回另一个链表
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	//初始化一个链表节点指针
	var l *ListNode
	//比较两个链表的值，将其中小的值赋值给链表，然后进行下一次循环处理
	if l1.Val > l2.Val {
		l = l2
		l.Next = mergeTwoLists(l1, l2.Next)
	} else {
		l = l1
		l.Next = mergeTwoLists(l1.Next, l2)
	}

	return l
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoLists(l1, l2)
}
