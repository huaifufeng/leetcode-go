package questionBank

import "testing"

func TestReverseList(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}

	res := ReverseList(list)

	value := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil}}}

	for value != nil && res != nil {
		if value.Val != res.Val {
			t.Log(value)
			t.Error("逆转链表方法验证失败")
		}

		value = value.Next
		res = res.Next
	}
}

func TestReverseList2(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}

	res := ReverseList2(list)

	value := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil}}}

	for value != nil && res != nil {
		if value.Val != res.Val {
			t.Log(value)
			t.Error("逆转链表方法2验证失败")
		}

		value = value.Next
		res = res.Next
	}
}
