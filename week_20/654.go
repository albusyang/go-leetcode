package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	fmt.Println(constructMaximumBinaryTree(nums))
}


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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	root := &TreeNode{
		Val: nums[0],
		Left: nil, 
		Right: nil,
	}
	if len(nums) == 1 {
		return root
	}
	
	for i := 1; i < len(nums); i++{
		if nums[i] > root.Val {
			greaterThanRoot(root, nums[i])
		} else {
			smallerThanRoot(root, nums[i])
		}
	}
	return root
}

func greaterThanRoot(root *TreeNode, v int) {
	oldRoot := &TreeNode{root.Val, root.Left, root.Right}
	root.Left = oldRoot
	root.Right = nil
	root.Val = v
}

func smallerThanRoot(root *TreeNode, v int) {
	if root.Right == nil {
		newNode := &TreeNode{v, nil, nil}
		root.Right = newNode
	} else {
		if v > root.Right.Val {
			greaterThanRoot(root.Right, v)
		} else {
			smallerThanRoot(root.Right, v)
		}
	}
}