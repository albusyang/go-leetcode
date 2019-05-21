package main

import "fmt"

func main() {
nums := []int{1,1,1,2,2,3,3,3,4,5,5,5,6,6,7,7,7,7,7,7,7,8,9,9}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	flag := 1
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j-1] == nums[j] && flag == 1{
			flag = 2
			i++
			nums[i] = nums[j]
			
		}
		if nums[j-1] != nums[j] {
			flag = 1
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums[:i+1])
	return i + 1
}