package main

import "fmt"

func main() {
	t17 := TreeNode{
		Val: 7,
		Left: nil,
		Right: nil,
	}
	t16 := TreeNode{
		Val: 6,
		Left: nil,
		Right: nil,
	}
	t14 := TreeNode{
		Val: 5,
		Left: nil,
		Right: nil,
	}
	t13 := TreeNode{
		Val: 4,
		Left: &t17,
		Right: nil,
	}
	t12 := TreeNode{
		Val: 3,
		Left: nil,
		Right: &t16,
	}
	t11 := TreeNode{
		Val: 2,
		Left: &t13,
		Right: &t14,
	}
	t10 := TreeNode{
		Val: 1,
		Left: &t11,
		Right: &t12,
	}
	fmt.Println(minDepth(&t10))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if minDepth(root.Left) > minDepth(root.Right) {
		return minDepth(root.Right) + 1
	}
	return minDepth(root.Left) + 1
}