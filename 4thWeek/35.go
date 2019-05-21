package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7}
	fmt.Println(searchInsert(nums, 7))
}

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v < target {
			if i < len(nums) - 1 {
				continue
			} else {
				return i + 1
			}
		} else {
			return i
		}
	}
	return 0
}