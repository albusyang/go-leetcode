package main

import (
	"fmt"
)

func main() {
	ss := []int{1, 10, 2, 4, 6, 7, 11, 8, 9}
	fmt.Println(lastStoneWeight(ss))
}

func lastStoneWeight(stones []int) int {
    if len(stones) < 2 {
		if len(stones) == 0 {
			return 0
		}
		return stones[0]
	}
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(stones); j++ {
			if stones[i] < stones[j] {
				stones[i], stones[j] = stones[j], stones[i]
			}
		}
	}
	if stones[len(stones)-1] != 0 && stones[len(stones)-2] != 0 {
		les := stones[len(stones)-1] - stones[len(stones)-2]
		if les == 0 {
			stones[len(stones)-1] = 0
			stones[len(stones)-2] = 0
		} else {
			stones[len(stones)-1] = les
			stones[len(stones)-2] = 0
		}
		
	}
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(stones); j++ {
			if stones[i] < stones[j] {
				stones[i], stones[j] = stones[j], stones[i]
			}
		}
	}
	if stones[len(stones)-2] != 0 {
		lastStoneWeight(stones)
	}
	return stones[len(stones)-1]
}