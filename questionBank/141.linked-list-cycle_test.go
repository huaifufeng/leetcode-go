package questionBank

import "testing"

//判断有环
func TestHasCycleHas(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}

	list.Next.Next.Next = list.Next

	has := hasCycle(list)
	if has != true {
		t.Error("判断链表是否有环方法错误")
	}
}

//判断没有环
func TestHasCycleNotHas(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}

	has := hasCycle(list)
	if has != false {
		t.Error("判断链表是否有环方法错误")
	}
}

//判断有环
func TestHasCycleHas2(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}

	list.Next.Next.Next = list.Next

	has := hasCycle2(list)
	if has != true {
		t.Error("判断链表是否有环方法错误")
	}
}

//判断没有环
func TestHasCycleNotHas2(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}

	has := hasCycle2(list)
	if has != false {
		t.Error("判断链表是否有环方法错误")
	}
}
