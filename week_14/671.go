package main

import (
	"fmt"
)

func main() {
	t1 := TreeNode{1, nil, nil}
	t2 := TreeNode{3, nil, nil}
	t3 := TreeNode{3, &t1, &t2}
	fmt.Println(findSecondMinimumValue(&t3))
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
    tnArr := make([]*TreeNode, 0)
	index := 0
	tnArr = append(tnArr, root)
	for index < len(tnArr) {
		if tnArr[index].Left != nil {
			tnArr = append(tnArr, tnArr[index].Left)
		}
		if tnArr[index].Right != nil {
			tnArr = append(tnArr, tnArr[index].Right)
		}
		index++
	}
	
	if len(tnArr) < 2 {
		return -1
	}
	
	for i := 0; i < len(tnArr); i++ {
		for j := 0; j <len(tnArr); j++ {
			if tnArr[i].Val < tnArr[j].Val {
				tnArr[i], tnArr[j] = tnArr[j], tnArr[i]
			}
		}
	}
	
	if tnArr[0].Val < tnArr[1].Val {
		return tnArr[1].Val
	} else {
		for ii := 1; ii < len(tnArr)-1; ii++ {
			if tnArr[ii].Val < tnArr[ii+1].Val {
				return tnArr[ii+1].Val
			}
		}
		return -1
	}
}