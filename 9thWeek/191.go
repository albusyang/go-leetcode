package main

import "fmt"

func main() {
    fmt.Println(hammingWeight(3))
}

func hammingWeight(num uint32) int {
	bn := decimalToBinary(int(num))
	hw := 0
	for _, v := range bn {
		if v == 1 {
			hw++
		}
	}
	return hw
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
