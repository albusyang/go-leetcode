package main

import (
	"fmt"
)

func main() {
	t1 := TreeNode{2, nil, nil}
	t2 := TreeNode{1, nil, nil}
	t3 := TreeNode{3, &t1, &t2}
	fmt.Println(preorderTraversal(&t3))
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func preorderTraversal(root *TreeNode) []int {
	tvArr := make([]int, 0)
	if root == nil {
		return tvArr
	}
	preorder(root, &tvArr)
	return tvArr
}

func preorder(root *TreeNode, tvArr *[]int) {
	if root == nil{
		return
	}
	*tvArr = append(*tvArr, root.Val)
	preorder(root.Left, tvArr)
	preorder(root.Right, tvArr)
}