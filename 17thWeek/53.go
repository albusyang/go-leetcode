package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func maxSubArray(nums []int) int {
    result := nums[0]
	for i := 0; i < len(nums); i++ {
		res := 0
		for j := i; j < len(nums); j++ {
			if res == 0 {
				if nums[j] > result {
					result = nums[j]
				}
			}
			res += nums[j]
			if res < 0 {
				break
			}
			if res > result {
				result = res
			}
		}
	}
	return result
}