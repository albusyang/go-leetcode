package main

import "fmt"

func main() {
	fmt.Println(relativeSortArray([]int{2,3,1,3,2,4,6,7,9,2,19},[]int{2,1,4,3,9,6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
    result := make([]int, 0)
	tail := make([]int, 0)
	for _, v := range arr2 {
		for _, vv := range arr1 {
			if vv == v {
				result = append(result, vv)
			}
		}
	}
	isTail := false
	for _, v := range arr1 {
		isTail = true
		for _, vv := range result {
			if vv == v {
				isTail = false
			}
		}
		if isTail {
			tail = append(tail, v)
		}
	}
	for i := 0; i < len(tail); i++ {
		for j := i+1; j <len(tail); j++ {
			if tail[i] > tail[j] {
				tail[i], tail[j] = tail[j], tail[i]
			}
		}
	}
	result = append(result, tail...)
	return result
}