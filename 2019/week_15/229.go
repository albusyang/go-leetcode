package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}

func majorityElement(nums []int) []int {
	result := make([]int, 0)
	rmap := make(map[int]int)
	for _, v := range nums {
		_, ok := rmap[v]
		if ok {
			rmap[v] = rmap[v]+1
		} else {
			rmap[v] = 1
		}
	}
	for k := range rmap {
		if rmap[k] > len(nums)/3 {
			result = append(result, k)
		}
	}	
	return result
}