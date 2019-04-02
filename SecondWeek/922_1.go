package main

import (
	"fmt"
)

func main() {
	arr := []int{4,1,1,0,1,0}
	fmt.Println(sortArrayByParityII(arr))
}
func sortArrayByParityII(A []int) []int {
    c := 0
	d := len(A) - 2
	
	for i := range A {
		if A[i] % 2 == 0 {
			A[i], A[c] = A[c], A[i]
			c = c + 1
		} 
	}
	fmt.Println(A)
	for i := 0; i < len(A) / 2; i++ {
		if i % 2 == 0 {
			continue
		}
		A[i], A[d] = A[d], A[i]
		d = d - 2 
	}
	return A
}