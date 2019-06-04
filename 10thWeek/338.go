package main

import "fmt"

func main() {
    fmt.Println(countBits(5))
}

func countBits(num int) []int {
	bnArr := make([]int, num+1)
	for i := 0; i <= num; i++{
		bn := decimalToBinary(i)
		hw := 0
		for _, v := range bn {
			if v == 1 {
				hw++
			}
		}
		bnArr[i] = hw
	}
	return bnArr
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
    return raw
}
