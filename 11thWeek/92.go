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
	n2 := ListNode{
		Val: 1,
		Next: &n1,
	}
	nn := reverseBetween(&n2, 1, 2)
	fmt.Println(nn)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    pre := &ListNode{0, head}
	preHead := pre
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	
	cur := pre.Next
	for i := m; i < n; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}
	return preHead.Next
}