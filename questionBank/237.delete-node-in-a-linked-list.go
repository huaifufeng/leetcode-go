/**
  题目地址：https://leetcode-cn.com/problems/delete-node-in-a-linked-list/

  解题：
  1、循环链表往后走，直到找到指定的值，O(n)
*/
package questionBank

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}
