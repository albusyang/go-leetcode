package main

import "fmt"

func main() {
	tn1 := &TreeNode{Val: 15, Left: nil, Right: nil}
	tn2 := &TreeNode{Val: 14, Left: nil, Right: nil}
	tn3 := &TreeNode{Val: 13, Left: tn1, Right: tn2}
	tn4 := &TreeNode{Val: 12, Left: tn3, Right: nil}
	fmt.Println(levelOrderBottom(tn4))
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func levelOrderBottom(root *TreeNode) [][]int {
    result := make([][]int, 0)
	if root == nil {
		return result
	}
	
	result = collect(0, result, root)
	
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	
	return result
}

func collect(layer int, result [][]int, node *TreeNode) [][]int{
	if layer < len(result) {
		result[layer] = append(result[layer], node.Val)
	} else {
		ss := make([]int, 0)
		ss = append(ss, node.Val)
		result = append(result, ss)
	}
	
	if node.Left != nil {
		result = collect(layer+1, result, node.Left)
	}
	
	if node.Right != nil {
		result = collect(layer+1, result, node.Right)
	}
	
	return result
}