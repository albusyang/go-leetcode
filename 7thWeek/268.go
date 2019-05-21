package main

import "fmt"

func main() {
	nums := []int{9,6,4,2,3,5,7,0,1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	if len(nums) == 1 {
		if nums[0] == 0 {
			return 1
		}
		return 0
	}
    max := len(nums)+1
	
	for i := 0; i < max; i++ {
		isExist := false
		for _, v := range nums {
			if i == v {
				isExist = true 
			}
		}
		if !isExist {
			return i
		}
	}
	return max
}