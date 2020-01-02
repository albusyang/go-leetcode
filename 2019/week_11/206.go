package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	n0 := ListNode{
		Val: 1,
		Next: nil,
	}
	n1 := ListNode{
		Val: 6,
		Next: &n0,
	}
	nn := reverseList(&n1)
	fmt.Println(nn)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
    lnl := make([]*ListNode, 0)
	lnl = append(lnl, head)
	for head.Next != nil {
		lnl = append(lnl, head.Next)
		head = head.Next
	}
	mid := head
	for i := len(lnl) - 2; i >= 0; i-- {
		mid.Next = lnl[i]
		mid = mid.Next
	}
	mid.Next = nil
	return head
}