package main

import "fmt"

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(nums, 2))    
}

func removeElement(nums []int, val int) int {
	index := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			sum++
		}
	}
    for i := 0; i < len(nums) - sum; i++ {
		index = i
		if nums[i] == val {
			for true {
				if nums[index] == nums[i] {
					if index < len(nums) - 1 {
						index++
					} else {
						break
					}
				} else {
					break				
				}
			}
			nums[i], nums[index] = nums[index], nums[i]
		} 
	}
	return len(nums) - sum
}
