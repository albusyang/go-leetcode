package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}

func maxProfit(prices []int) int {
	result := 0
	for i := 0; i < len(prices); i++ {
		profit := 0
		for j := i+1; j < len(prices); j++ {
            if prices[i] > prices[j] {
                break
            }
			if prices[j]-prices[i] > profit {
				profit = prices[j]-prices[i]
			}
		}
		if profit > result {
			result = profit
		}
	}
	return result
}