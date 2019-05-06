package main

import "fmt"

type ListNode struct {
    Val int
	Next *ListNode
}

func main() {
	lo3 := ListNode{4, nil}
	lo2 := ListNode{2, &lo3}
	lo1 := ListNode{1, &lo2}
	
	lt3 := ListNode{4, nil}
	lt2 := ListNode{3, &lt3}
	lt1 := ListNode{1, &lt2}
	fmt.Println(mergeTwoLists(&lo1, &lt1))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}