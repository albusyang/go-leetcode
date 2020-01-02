package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	root := ListNode{1, nil}
	fmt.Println(middleNode(&root))
}

func middleNode(head *ListNode) *ListNode {
    var lnList []*ListNode
	myNode := head
	for myNode != nil {
		lnList = append(lnList, myNode)
		myNode = myNode.Next
	}
	
	index := len(lnList)/2
	return lnList[index]
}