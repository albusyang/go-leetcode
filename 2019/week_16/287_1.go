package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1,3,4,2,2}))
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		index := 0
		if nums[i] > 0 {
			index = nums[i]
		} else {
			index = -nums[i]
		}
		if nums[index] < 0 {
			return index
		} else {
			nums[index] *= -1
		}
	}
	return -1
}