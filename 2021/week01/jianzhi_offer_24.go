package week01

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// input: 1->2->3->4->5->NULL
//
// output: 5->4->3->2->1->NULL
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	inx := head
	for inx != nil {
		next := inx.Next
		inx.Next = pre
		pre = inx
		inx = next
	}
	return pre
}
