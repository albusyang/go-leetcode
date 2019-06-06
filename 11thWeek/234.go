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
	nn := isPalindrome(&n2)
	fmt.Println(nn)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
    lnl := make([]int, 0)
	lnl = append(lnl, head.Val)
	for head.Next != nil {
		lnl = append(lnl, head.Next.Val)
		head = head.Next
	}
	if len(lnl) == 1 {
		return true
	}
	l := 0
	r := 0
	isP := true
	if len(lnl) % 2 == 0 {
		l = len(lnl)/2 - 1
		r = len(lnl)/2
	} else {
		l = len(lnl)/2 - 1
		r = len(lnl)/2 + 1
	}
	for i := 0; i < len(lnl)/2; i++ {
			if lnl[l-i] != lnl[r+i] {
				isP = false
				break
			}			
		}
	return isP
}