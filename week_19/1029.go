package main

import "fmt"

func main() {
	costs := [][]int{{10,20},{30,200},{400,50},{30,20},{1,20},{100,1}}
	fmt.Println(twoCitySchedCost(costs))
}

func twoCitySchedCost(costs [][]int) int {
	bF := make([]int, 0)
	
	// first A then B
	arr := make([][]int, 0)
	// costB - costA
	for i, v := range costs {
		bMinA := v[1] - v[0]
		temp := []int{i, bMinA}
		arr = append(arr, temp)
	}
	bubble(arr)
	
	count := 0 
	for i := range arr {	
		if count == len(costs)/2 {
			break
		}
		bF = append(bF, arr[i][0])
		count++
	}
	sum := 0 
	for i := range costs {
		isB := false
		for _, j:= range bF {
			if j == i {
				sum += costs[i][1]
				isB = true
				continue
			}
		}
		if !isB {
			sum += costs[i][0]
		}
	}
	return sum
}

func bubble (one [][]int) {
	for i := 0; i < len(one); i++ {
		for j := 0; j < len(one); j++ {
			if one[j][1] > one[i][1] {
				one[j], one[i] = one[i], one[j]
			}
		}
	}
	
}