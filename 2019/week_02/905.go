package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 2}
	fmt.Println(sortArrayByParity(arr))
}

func sortArrayByParity(A []int) []int {
	c := 0
	for i := range A {
		if A[i]%2 == 0 {
			A[c], A[i] = A[i], A[c]
			c++
		}
	}
	return A
}
