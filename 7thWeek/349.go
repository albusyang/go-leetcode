package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 3, 3, 3}
	nums2 := []int{2, 2, 3}
	fmt.Println(intersect(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
    myNums := make([]int, 0)
	
	for _, v := range nums1 {
		overlap := false
		for _, v1 := range nums2 {
			if v == v1 {
				overlap = true
			}
		}
		if overlap {
			exist := false
			for _, v2 := range myNums {
				if v == v2 {
					exist = true
				}
			}
			if !exist {
				myNums = append(myNums,v)
			}
		}
	}
	return myNums
}