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
	air := make([]int, 0)
	for _, v := range A {
		v *= v
		fmt.Println(v)
		air = append(air, v)
	}
	sort.Ints(air)
	
    return air
}