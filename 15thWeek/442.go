package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{1,3,4,2,5,2,6,4}))
}

func findDuplicates(nums []int) []int {
    result := make([]int, 0)
	rMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := rMap[nums[i]]
		if ok {
			result = append(result, nums[i])
		} else {
			rMap[nums[i]] = 1
		}
	}
	return result
}