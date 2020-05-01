/**
  题目地址：https://leetcode-cn.com/problems/linked-list-cycle/

  给定一个链表，判断链表中是否有环。
  为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
  如果 pos 是 -1，则在该链表中没有环。


  <pre>
    输入：head = [3,2,0,-4], pos = 1
    输出：true
    解释：链表中有一个环，其尾部连接到第二个节点。
  </pre>

  <pre>
    输入：head = [1,2], pos = 0
    输出：true
    解释：链表中有一个环，其尾部连接到第一个节点。
  </pre>

  <pre>
    输入：head = [1], pos = -1
    输出：false
    解释：链表中没有环。
  </pre>

  解题：
	1、第一种就是暴力破解，依次对比数组的每个元素和之后的元素 O(n^2)
    2、第二种是循环，获取当前的最小价格，拿着当前的价格和这个最小价格比较，获取最大的收益值 O(n)
*/
package questionBank

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	//没有节点或只有一个节点 不是环
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func hasCycle2(head *ListNode) bool {
	//没有节点或只有一个节点 不是环
	if head == nil || head.Next == nil {
		return false
	}

	nodeMap := make(map[*ListNode]int)

	for head != nil {
		fmt.Print(nodeMap)
		//如果节点在hash中存在，说明之前遇到过
		if _, ok := nodeMap[head]; ok {
			return true
		}

		nodeMap[head] = 1
		head = head.Next
	}

	return false
}
