package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1,2,3}, {4,5,6}, {7,8,9}}))
}

func transpose(A [][]int) [][]int {
    y := len(A)
	if y == 0 {
		return A
	}
	x := len(A[0])
	
	result := make([][]int, 0)
	for i := 0; i < x; i ++ {
		rr := make([]int, 0)
		for j := 0; j < y; j++ {
			rr = append(rr, A[j][i])
		}
		result = append(result, rr)
	}
	
	return result
}