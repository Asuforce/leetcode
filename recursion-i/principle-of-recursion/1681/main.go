package main

import "fmt"

func main() {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(swapPairs(l))
}

// ListNode is linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nextNode := head.Next
	nextNextNode := nextNode.Next

	nextNode.Next = head
	head.Next = swapPairs(nextNextNode)

	return nextNode
}
