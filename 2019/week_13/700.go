package main

import (
	"fmt"
)

func main() {
	n1 := TreeNode{0, nil, nil}
	n2 := TreeNode{1, nil, nil}
	root := TreeNode{2, &n1, &n2}
	fmt.Println(searchBST(&root, 1))
}

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
    nodeList := make([]*TreeNode, 0)
	cur := 0
	
	nodeList = append(nodeList, root)
	for cur < len(nodeList) {
		if nodeList[cur].Val == val {
			return nodeList[cur]
		}
		if nodeList[cur].Left != nil {
			nodeList = append(nodeList, nodeList[cur].Left)
		}
		if nodeList[cur].Right != nil {
			nodeList = append(nodeList, nodeList[cur].Right)
		}
		cur++
	}
	
	return nil
}