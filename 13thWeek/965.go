package main

import (
	"fmt"
)

func main() {
	n1 := TreeNode{0, nil, nil}
	n2 := TreeNode{1, nil, nil}
	root := TreeNode{1, &n1, &n2}
	fmt.Println(isUnivalTree(&root))
}

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
func isUnivalTree(root *TreeNode) bool {
	comV := root.Val
    nodeList := make([]*TreeNode, 0)
	//pre := 0
	cur := 0
	
	nodeList = append(nodeList, root)
	for cur < len(nodeList) {
		if nodeList[cur].Val != comV {
			return false
		}
		if nodeList[cur].Left != nil {
			nodeList = append(nodeList, nodeList[cur].Left)
		}
		if nodeList[cur].Right != nil {
			nodeList = append(nodeList, nodeList[cur].Right)
		}
		cur++
	}
	
	return true
}