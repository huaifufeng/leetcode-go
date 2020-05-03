package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

type SingleList struct {
	head *ListNode
}

func NewListNode(val int) *ListNode {
	listNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	return listNode
}

func NewList(head *ListNode) *SingleList {
	list := &SingleList{
		head: head,
	}

	return list
}

//查找指定值的节点
func (l *SingleList) findByValue(val int) *ListNode {
	cur := l.head
	for cur != nil {
		if cur.Val == val {
			return cur
		}

		cur = cur.Next
	}

	return nil
}

//查找指定索引的节点
func (l *SingleList) findByIndex(index int) *ListNode {
	cur := l.head
	pos := 0
	for cur != nil {
		if pos == index {
			return cur
		}
		pos++
		cur = cur.Next
	}

	return nil
}

//在指定值的元素后面插入节点
func (l *SingleList) insertAfter(node *ListNode, val int) bool {
	//不符合条件的节点不能插入
	if node == nil || node == l.head {
		return false
	}

	//头结点不存在的话，节点直接做头结点
	if l.head == nil {
		l.head = node
		return true
	}

	findNode := l.findByValue(val)
	//没有找到指定的值，不进行处理
	if findNode == nil {
		return false
	}

	node.Next = findNode.Next
	findNode.Next = node

	return true
}

//删除指定值的节点
func (l *SingleList) deleteNode(val int) bool {
	if l.head == nil {
		return false
	}

	//头结点单独处理，因为删除节点需要保留前一个节点的信息，和头结点的处理不同
	if l.head.Val == val {
		l.head = l.head.Next
		return true
	}

	cur := l.head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			return true
		}
		cur = cur.Next
	}

	return false
}
