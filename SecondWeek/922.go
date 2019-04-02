package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,3,4}
	fmt.Println(sortArrayByParityII(arr))
}
func sortArrayByParityII(A []int) []int {
    arr := make([]int, len(A))
    c := 0
	d := 1
	for _, v := range A {
		if v % 2 == 0 {
			arr[c] = v
			c = c + 2
		} else {
			arr[d] = v
			d = d + 2
		}
	}
	return arr
}