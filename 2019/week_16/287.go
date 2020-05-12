package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1,3,4,2,2}))
}

func findDuplicate(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    result := 0
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				result = nums[i]
			}
		}
	}
	return result
}