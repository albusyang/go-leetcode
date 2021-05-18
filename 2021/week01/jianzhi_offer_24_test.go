package week01

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	node0 := &ListNode{0, node1}
	n := node0
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}

	node := reverseList(node0)

	inx := node
	for inx != nil {
		fmt.Println(inx)
		inx = inx.Next
	}
}
