package algorithm

import "testing"

func TestList(t *testing.T) {
	//初始化头结点为空的节点
	list := NewList(NewListNode(0))

	//在值为0的节点后插入元素
	node1 := NewListNode(1)
	list.insertAfter(node1, 0)

	//在值为1的节点后插入元素
	node2 := NewListNode(2)
	list.insertAfter(node2, 1)

	//在值为0的节点后插入元素
	node3 := NewListNode(3)
	list.insertAfter(node3, 0)

	node4 := list.findByValue(2)
	if node4 != node2 {
		t.Error("查找指定值的节点错误")
	}

	res := list.deleteNode(2)
	if res != true {
		t.Error("删除数据错误")
	}
}
