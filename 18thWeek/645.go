package main

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{1,5,3,2,2,7,6,4,8,9}))
}

func findErrorNums(nums []int) []int {
    if len(nums) < 1 {
		return nums
	}
	
	result := make([]int, 2)
	tempArr := make([]int, 0)
	
	for i := 0; i < len(nums); i++ {
		isExist := false
		for _, v := range tempArr {
			if v == nums[i] {
				isExist = true
			}
		}
		if !isExist {
			tempArr = append(tempArr, nums[i])
		}
		for j := i+1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				result[0] = nums[i]
				continue
			}
			
		}
	}
	
	for i := 0; i < len(tempArr); i++ {
		for j := i+1; j < len(tempArr); j++ {
			if tempArr[i] > tempArr[j] {
				tempArr[i], tempArr[j] = tempArr[j], tempArr[i]
			}
		}
	}
	
	ss := tempArr[0]
	
	for i := 0; i < len(tempArr); i++ {
		if tempArr[i] > ss+i {
			result[1] = ss+i
			return result
		}
		if tempArr[i] < ss+i {
			if tempArr[i] == 2 {
				result[1] = 1
				return result
			} else {
				result[1] = ss+i
				return result
			}
		}
	}

	if ss == 2 {
		result[1] = ss-1
		return result
	}
	result[1] = tempArr[len(tempArr)-1]+1
	return result
}