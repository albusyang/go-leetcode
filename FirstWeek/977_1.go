package main

import (
	"fmt"
)

func main() {
	aa := []int{-7, -3, 2, 3, 1}
	arr := sortedSquares(aa)
	for _, v := range arr {
		fmt.Println(v)
	}
}

func sortedSquares(A []int) []int {
	arr := make([]int, 0)
	for _, v := range A {
		v *= v
		fmt.Println(v)
		arr = append(arr, v)
	}
	// sort.Ints(arr)
	quickSort(arr, 0, len(arr)-1)

	return arr
}

func bubleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func quickSort(arr []int, a, b int) {
	if a >= b {
		return
	}
	c := a
	p := arr[b]
	for i := a; i < b; i++ {
		if arr[i] < p {
			arr[i], arr[c] = arr[c], arr[i]
			c++
		}
	}
	arr[c], arr[b] = arr[b], arr[c]
	quickSort(arr, a, c-1)
	quickSort(arr, c+1, b)
}
