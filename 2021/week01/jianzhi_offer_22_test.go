package week01

import (
	"fmt"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
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

	r := getKthFromEnd(node0, 3)

	fmt.Println(r)
}

func Test_getKthFromEnd1(t *testing.T) {
	// n := 10
	// for i := 0; i < n; i++ {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	node0 := &ListNode{0, node1}
	// }
	n := node0
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}

	r := getKthFromEnd1(node0, 7)

	fmt.Println(r)
}

func Test_getKthFromEnd2(t *testing.T) {
	// n := 10
	// for i := 0; i < n; i++ {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	node0 := &ListNode{0, node1}
	// }
	n := node0
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}

	r := getKthFromEnd2(node0, 3)

	fmt.Println(r)
}
