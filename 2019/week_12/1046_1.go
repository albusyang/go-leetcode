package main

import "fmt"

func main() {
	stones := []int{5,5,5,10,20}
	fmt.Println(lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
    if len(stones) == 1 {
        return stones[0]
    }
	for {
		bubble(stones)
		if stones[1] == 0 {
			return stones[0]
		}
		if stones[0] == stones[1] {
			stones[0] = 0
			stones[1] = 0
		} else {
			stones[0] = stones[0] - stones[1]
			stones[1] = 0
		}
	}
	return 0
}

func bubble (stones []int) {
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(stones); j++ {
			if stones[i] > stones[j] {
				stones[i], stones[j] = stones[j], stones[i]
			}
		}
	}
}