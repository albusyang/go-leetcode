package main

import "fmt"

func main() {
	arr := []int{3,3}
	t := 6
	fmt.Println(twoSum(arr, t))
}

func twoSum(nums []int, target int) []int {
	var towInt []int
    for i, v := range nums {
		for ii := i + 1; ii < len(nums); ii++ {
			if v + nums[ii] == target {
				towInt = append(towInt, i)
				towInt = append(towInt, ii)
			}
		}
	}
	return towInt
}