package main

import (
	"fmt"
)

func main() {
	t1 := TreeNode{2, nil, nil}
	t2 := TreeNode{1, nil, nil}
	t3 := TreeNode{3, &t1, &t2}
	fmt.Println(kthSmallest(&t3,2))
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func kthSmallest(root *TreeNode, k int) int {
	if root == nil || k == 0{
		return 0
	}
	tnArr := make([]int, 0)
	inOrder(root, &tnArr)
	return tnArr[k-1]
}

func inOrder(root *TreeNode, tnArr *[]int) {
	if root == nil{
		return
	}
	inOrder(root.Left, tnArr)
	*tnArr = append(*tnArr, root.Val)
	inOrder(root.Right, tnArr)
}