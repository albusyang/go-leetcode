package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 3, 3, 3}
	nums2 := []int{2, 2, 3}
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
    myNums := make([]int, 0)
	myMap := make(map[int]int)
	
	for _, v := range nums1 {
		_, ok := myMap[v] 
		if !ok {
			myMap[v] = 1
		} else {
			myMap[v] = myMap[v] + 1
		}
	}
	for k, v := range myMap {
		count := 0
		for _, v1 := range nums2 {
			if k == v1 {
				count++
			}
		}
		if count < v {
			myMap[k] = count
		}
	}
	for k, v := range myMap {
		for i := 0; i < v; i++ {
			myNums = append(myNums, k)
		}
	}
	return myNums
}