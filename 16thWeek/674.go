package main 

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1,3,5,4,7}))
}

func findLengthOfLCIS(nums []int) int {
    result := 0
	if len(nums) == 0 {
		return result
	}
	temp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			temp++
		} else {
			if temp > result {
				result = temp
			}
			temp = 1
		}
	}
	if temp > result {
		return temp
	}
	return result
}