package main

import (
	"fmt"
)

func main() {
	Arr := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	fmt.Println(flipAndInvertImage(Arr))
}

func flipAndInvertImage(A [][]int) [][]int {

	var transformArr = make([][]int, len(A))

	for i := range A {
		transformArrPart := make([]int, len(A[i]))
		for j := range A[i] {
			if A[i][j] == 1 {
				transformArrPart[len(A[i])-j-1] = 0
			}
			if A[i][j] == 0 {
				transformArrPart[len(A)-j-1] = 1
			}
		}
		transformArr[i] = transformArrPart
	}
	return transformArr
}
