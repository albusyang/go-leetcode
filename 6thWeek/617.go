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
	
	
	t27 := TreeNode{
		Val: 7,
		Left: nil,
		Right: nil,
	}
	t26 := TreeNode{
		Val: 6,
		Left: nil,
		Right: nil,
	}
	t24 := TreeNode{
		Val: 5,
		Left: nil,
		Right: nil,
	}
	t23 := TreeNode{
		Val: 4,
		Left: nil,
		Right: &t27,
	}
	t22 := TreeNode{
		Val: 3,
		Left: nil,
		Right: &t26,
	}
	t21 := TreeNode{
		Val: 2,
		Left: &t23,
		Right: &t24,
	}
	t20 := TreeNode{
		Val: 1,
		Left: &t21,
		Right: &t22,
	}

	fmt.Println(mergeTrees(&t10, &t20))
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
 
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode{
	if t1 != nil && t2 != nil {
		t1.Val += t2.Val
		t1.Left = mergeTrees(t1.Left, t2.Left)
		t1.Right = mergeTrees(t1.Right, t2.Right)
	}
	if t1 == nil && t2 != nil {
		t1 = t2
	}
    return t1
}