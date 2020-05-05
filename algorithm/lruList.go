package algorithm

type LruList struct {
	head   *ListNode
	length int
	count  int
}

func NewLruList(len int) *LruList {
	lru := &LruList{
		head:   nil,
		length: len,
		count:  0,
	}

	return lru
}

func (ll *LruList) getItem(val int) int {
	if ll.exists(val) {
		ll.moveToTail(val)
		return val
	} else {
		if ll.count == ll.length {
			ll.head = ll.head.Next
			ll.addNote(val)
			return val
		} else {
			ll.addNote(val)
			return val
		}
	}
}

func (ll *LruList) moveToTail(val int) {
	if ll.count == 1 {
		return
	}

	//头指针单独处理
	if ll.head.Val == val {
		node := ll.head

		ll.head = ll.head.Next
		head := ll.head

		for head.Next != nil {
			head = head.Next
		}

		node.Next = nil
		head.Next = node
		return
	}

	head := ll.head
	var node *ListNode
	for head.Next != nil {
		if head.Next.Val == val {
			//如果是链表的尾部节点，不进行处理
			if head.Next.Next == nil {
				return
			}

			node = head.Next
			head.Next = head.Next.Next
			node.Next = nil
		}

		head = head.Next
	}

	head.Next = node
}

func (ll *LruList) addNote(val int) {
	if ll.head == nil {
		ll.head = NewListNode(val)
		ll.count++
		return
	}

	head := ll.head
	for head.Next != nil {
		head = head.Next
	}

	head.Next = NewListNode(val)
	ll.count++

	return
}

func (ll *LruList) exists(val int) bool {
	head := ll.head
	for head != nil {
		if head.Val == val {
			return true
		}
		head = head.Next
	}

	return false
}
