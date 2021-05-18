package week01

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	node3 := &ListNode{6, nil}
	node2 := &ListNode{5, node3}
	node1 := &ListNode{4, node2}

	node03 := &ListNode{3, nil}
	node02 := &ListNode{2, node03}
	node01 := &ListNode{1, node02}

	node := mergeTwoLists(node1, node01)

	inx := node
	for inx != nil {
		fmt.Println(inx)
		inx = inx.Next
	}
}
