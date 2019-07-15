package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{1,3,4,2,5,2,6,4}))
}

func findDisappearedNumbers(nums []int) []int {
   result := make([]int, 0)
	rMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rMap[nums[i]] = 1
	}
	for i := 1; i <= len(nums); i++ {
		_, ok := rMap[i] 
		if !ok {
			result = append(result, i)
		}
	}
	return result
}