package main

import (
	"fmt"
)

func main() {
	Arr := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	fmt.Println(flipAndInvertImage(Arr))
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := range A {
		for j := range A[i] {
			k := len(A[i]) - 1 - j
			if j >= k {
				break
			}
			A[i][j], A[i][k] = A[i][k], A[i][j]
		}
	}

	for a := range A {
		for b := range A[a] {
			if A[a][b] == 1 {
				A[a][b] = 0
				continue
			}
			if A[a][b] == 0 {
				A[a][b] = 1
			}
		}
	}
	return A
}
