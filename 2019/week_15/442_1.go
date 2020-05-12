package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{1,3,4,2,5,2,6,4}))
}

func findDuplicates(nums []int) []int {
    oneList := make([]int, len(nums)+1)
	result := make([]int, 0)
	for _, v := range nums {
		oneList[v]++
	}
	for i, v := range oneList {
		if v > 1 {
			result = append(result, i)
		}
	}
	return result
}