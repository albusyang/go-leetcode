package main

import "fmt"

func main() {
	nums := []int{0,1,0,1,0,1,99}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
    for i := 0; i < len(nums); i++ {
		isSingle := true
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i] == nums[j] {
				isSingle = false
				break
			}
		}
		if isSingle {
			return nums[i]
		}
	}
	return 0
}
