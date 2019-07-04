package main

import (
	"fmt"
)

func main() {
	t1 := &TreeNode{1, nil, nil}
	t2 := &TreeNode{2, nil, nil}
	t3 := &TreeNode{3, t1, t2}
	t4 := &TreeNode{4, t2, t1}
	fmt.Println(isSameTree(t3, t4))
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
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	
	pArr := make([]*TreeNode, 0)
	qArr := make([]*TreeNode, 0)
	pArr = append(pArr, p)
	qArr = append(qArr, q)
	
	pnum := 0
	qnum := 0
	
	for pnum < len(pArr) {
		if pArr[pnum] == nil {
			pnum++
			continue
		}
		pArr = append(pArr, pArr[pnum].Left)
		pArr = append(pArr, pArr[pnum].Right)
		pnum++
	}
	for qnum < len(qArr) {
		if qArr[qnum] == nil {
			qnum++
			continue
		}
		qArr = append(qArr, qArr[qnum].Left)
		qArr = append(qArr, qArr[qnum].Right)
		qnum++
	}
	
	if len(pArr) != len(qArr) {
		return false
	}
	
	for i:=0; i<len(pArr); i++ {
		if pArr[i] == nil && qArr[i] == nil {
			continue
		} 
		if pArr[i] == nil || qArr[i] == nil {
			return false
		}
		if pArr[i].Val != qArr[i].Val {
			return false
		}
	}
	return true
}