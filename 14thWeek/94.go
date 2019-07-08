package main

import (
	"fmt"
)

func main() {
	t1 := TreeNode{2, nil, nil}
	t2 := TreeNode{1, nil, nil}
	t3 := TreeNode{3, &t1, &t2}
	fmt.Println(inorderTraversal(&t3))
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func inorderTraversal(root *TreeNode) []int {
	tvArr := make([]int, 0)
	if root == nil {
		return tvArr
	}
	inorder(root, &tvArr)
	return tvArr
}

func inorder(root *TreeNode, tvArr *[]int) {
	if root == nil{
		return
	}
	inorder(root.Left, tvArr)
	*tvArr = append(*tvArr, root.Val)
	inorder(root.Right, tvArr)
}