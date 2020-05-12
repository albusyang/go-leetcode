package main

import (
	"fmt"
	"sort"
)

func main() {
	aa := []int{-7,-3,2,3,11}
	arr := sortedSquares(aa)
	for _, v := range arr {
		fmt.Println(v)
	}
}

func sortedSquares(A []int) []int {
	arr := make([]int, 0)
	for _, v := range A {
		v *= v
		arr = append(arr, v)
	}
	sort.Ints(arr)
	
    return arr
}