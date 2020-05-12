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
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}