package main

import "fmt"

func main() {
	nums := []int{1,2,1,3,2,5}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) []int {
	arr := make([]int, 0)
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
			arr = append(arr, nums[i])
		}
	}
	return arr
}
