/**
  题目地址：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

  输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
  例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点

  <pre>
     给定一个链表: 1->2->3->4->5, 和 k = 2.
     返回链表 4->5.
  </pre>

  解题：
    1、统计链表的长度n，然后获取第n-k个元素的内容 O(n), 循环两次链表
    2、双链表法。第一个链表先走k个节点，然后两个链表同时走，当第一个链表到达尾部的时候，第二个链表到达了倒数k的位置。 O(n)至多会循环一遍数组长度
*/
package offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	length := 0
	node := head
	for head != nil {
		length++
		head = head.Next
	}

	for i := 0; i < length-k; i++ {
		node = node.Next
	}

	return node
}

func getKthFromEnd1(head *ListNode, k int) *ListNode {
	node := head
	for k > 0 {
		head = head.Next
		k--
	}

	for head != nil {
		node = node.Next
		head = head.Next
	}

	return node
}
