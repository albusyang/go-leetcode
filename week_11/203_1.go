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
    iH := head 
	for iH.Next != nil {
		if iH.Next.Val == val {
			iH.Next = iH.Next.Next
		} else {
			iH = iH.Next
		}
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}