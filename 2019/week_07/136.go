package main

import "fmt"

func main() {
	nums := []int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
    for i:= 0; i<len(nums); i++ {
		isDup := false
		for j:=0; j<len(nums); j++ {
			if i == j {
				continue
			}
			if nums[j] == nums[i] {
				isDup = true
			}
		}
		if isDup {
			continue
		}
		return nums[i]
	}
	return 0
}