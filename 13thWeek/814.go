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
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
    
	for true {
		result, isClean := adjustTree(root, false)
		if isClean {
			return result
		}
	}
	return nil
}

func adjustTree(root *TreeNode, isClean bool) (*TreeNode, bool) {
	nodeList := make([]*TreeNode, 0)
	cur := 0
	
	iscur := false
	nodeList = append(nodeList, root)
	for cur < len(nodeList) {
		if nodeList[cur].Left != nil {
			if nodeList[cur].Left.Left == nil && nodeList[cur].Left.Right == nil && nodeList[cur].Left.Val == 0 {
				nodeList[cur].Left = nil
				iscur = true
			} else {
				nodeList = append(nodeList, nodeList[cur].Left)
			}
		}
		if nodeList[cur].Right != nil {
			if nodeList[cur].Right.Left == nil && nodeList[cur].Right.Right == nil && nodeList[cur].Right.Val == 0 {
				nodeList[cur].Right = nil
				iscur = true
			} else {
				nodeList = append(nodeList, nodeList[cur].Right)
			}
		}
		cur++
	}
	if iscur {
		isClean = false
	} else {
		isClean = true
	}
	return root, isClean
}