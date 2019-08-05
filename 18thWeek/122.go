package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}

func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			result += prices[i] - prices[i-1]
		} 
	}
	return result
}