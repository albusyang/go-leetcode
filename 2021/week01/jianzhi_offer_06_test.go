package week01

import (
	"fmt"
	"testing"
)

func Test_reversePrint(t *testing.T) {
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

	r := reversePrint(node0)

	fmt.Println(r)
}

func Test_reversePrint1(t *testing.T) {
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

	r := reversePrint1(node0)

	fmt.Println(r)
}

func Test_reversePrint2(t *testing.T) {
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

	r := reversePrint2(node0)

	fmt.Println(r)
}
