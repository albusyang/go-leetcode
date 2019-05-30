package main

import (
    "fmt"
)

func main() {
    fmt.Println(hasAlternatingBits(5)) 
}

func hasAlternatingBits(n int) bool {
	bn := decimalToBinary(n)
	for i := 0; i < len(bn) - 1; i++ {
		if bn[i] == bn[i+1] {
			return false
		}
	}
    return true
}

func decimalToBinary(num int) []int {
	raw := make([]int, 0)
    for true {
		if num != 0 {
			temp := num
			raw = append(raw, temp%2)
			num /= 2	
		} else {
			break
		}
	}
	sum := len(raw) - 1
	bba := make([]int, sum + 1)
	for k, v := range raw {
		bba[sum-k] = v
	}
    return bba
}
