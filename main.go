package main

import (
	"leetcode-go/questionBank"
)

type ListNode = questionBank.ListNode

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil}}}

	list.Next.Next.Next = list.Next

}
