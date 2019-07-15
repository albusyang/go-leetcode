package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{1,3,4,2,5,2,6,4}))
}

func findDisappearedNumbers(nums []int) []int {
    oneList := make([]int, len(nums)+1)
	result := make([]int, 0)
	for _, v := range nums {
		oneList[v]++
	}
	for i := 1; i < len(oneList); i++ {
		if oneList[i] == 0 {
			result = append(result, i)
		}
	}
	return result
}