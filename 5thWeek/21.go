package main

import "fmt"

type ListNode struct {
    Val int
	Next *ListNode
}

func main() {
	//lo3 := ListNode{4, nil}
	//lo2 := ListNode{2, &lo3}
	lo1 := ListNode{1, nil}
	
	lt4 := ListNode{5, nil}
	lt3 := ListNode{4, &lt4}
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
	var ln ListNode
	pln := &ln
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pln.Next = l1
			l1 = l1.Next
			fmt.Println("merge1")
		} else {
			pln.Next = l2
			l2 = l2.Next
			fmt.Println("merge2")
		}
		pln = pln.Next
	}
	if (l1 == nil) { 
		pln.Next = l2;
		fmt.Println("merge3")
	}
    if (l2 == nil) {
		pln.Next = l1;
		fmt.Println("merge4")
	}
	return ln.Next
}