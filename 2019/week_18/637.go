package main

import "fmt"

func main() {
	tn1 := &TreeNode{Val: 15, Left: nil, Right: nil}
	tn2 := &TreeNode{Val: 14, Left: nil, Right: nil}
	tn3 := &TreeNode{Val: 13, Left: tn1, Right: tn2}
	tn4 := &TreeNode{Val: 12, Left: tn3, Right: nil}
	fmt.Println(averageOfLevels(tn4))
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func averageOfLevels(root *TreeNode) []float64 {
    treeNodes := make([][]int, 0)
	if root == nil {
		return []float64{}
	}
	
	layerArr := collect(0, treeNodes, root)
	result := make([]float64, len(layerArr))
	for i, v := range layerArr {
		var rr float64
		var sum int
		for _, vv := range v {
			sum += vv
		}
		rr = float64(sum) / float64(len(v))
		result[i] = rr
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