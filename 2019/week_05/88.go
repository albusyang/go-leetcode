package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{1,3,5}
	merge1(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    newArr1 := nums1[:m]
	newList := append(newArr1, nums2...)
	sort.Ints(newList)
	nums1 = newList
}

func merge1(nums1 []int, m int, nums2 []int, n int)  {
    newArr1 := nums1[:m]
	newList := append(newArr1, nums2...)
	for i := range newList {
		for j := 1; j < len(newList) - i; j++ {
			if newList[j-1] > newList[j] {
				newList[j], newList[j-1] = newList[j-1], newList[j]
			}
		}
	}
	nums1 = newList
}