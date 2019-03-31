//题目地址：https://leetcode-cn.com/problems/add-two-numbers/
package 题库

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) print() {
	currentNode := l
	for currentNode != nil {
		currentNode = currentNode.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	l3Last := l3
	upValue := 0  //相加后上一位的值
	for l1 != nil || l2 != nil || upValue != 0 {
		if l1 != nil {
			upValue += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			upValue += l2.Val
			l2 = l2.Next
		}

		currentVal := upValue % 10
		upValue = upValue / 10

		currentNode := &ListNode{currentVal, nil}
		if l3 == nil {
			l3 = currentNode
			l3Last = currentNode
			continue
		}

		l3Last.Next = currentNode
		l3Last = currentNode
	}

	return l3
}

