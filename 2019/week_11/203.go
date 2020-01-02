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
	nn := removeElements(&n1,6)
	fmt.Println(nn)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} 
    if head.Next == nil {
		if head.Val == val {
			return head.Next
		}
		return head
	} 
    if head.Next.Val == val {
			head.Next = head.Next.Next
			removeElements(head, val)
	} else {
			removeElements(head.Next, val)
	}
	if head.Val == val {
		return head.Next
	}
	return head
}