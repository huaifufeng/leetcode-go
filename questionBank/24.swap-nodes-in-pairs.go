/**
  题目地址：https://leetcode-cn.com/problems/swap-nodes-in-pairs/

  给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
  你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

  <pre>
    给定 1->2->3->4, 你应该返回 2->1->4->3.
  </pre>

  解题： 时间复杂度O(n) n为数组长度
	循环整个数组，比较当前元素和pos-1位置的元素的值，如果不相等，就将pos的值替换为当前的元素的值
*/
package questionBank

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	//进行异常情况的判断
	if head == nil || head.Next == nil {
		return head
	}

	//因为在交换的时候需要获取元素对的前一个元素，这里声明链表头节点前面的元素节点
	prev := &ListNode{
		Val:  -1,
		Next: head,
	}
	//保存链表的头指针地址
	first := head.Next
	//当后面只有一个元素或没有元素时  不在继续处理
	for prev.Next != nil && prev.Next.Next != nil {
		left := prev.Next
		right := prev.Next.Next

		//左节点的下一个是右节点的下一个
		//右节点的下一个是左节点
		//之前元素的下一个是右节点
		//之前元素下次循环是左节点
		left.Next = right.Next
		right.Next = left
		prev.Next = right
		prev = left
	}

	return first
}
