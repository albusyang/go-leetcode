package main

import "fmt"

func main() {
	arr := []int{0, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(A []int) int {
    top := 0
	for i := range A {
		if A[i] > A[i + 1] {
			top = i 
			break
		}
	}
	return top
}