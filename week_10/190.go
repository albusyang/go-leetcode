package main

import "fmt"

func main() {
    fmt.Println(reverseBits(3))
}

func reverseBits(num uint32) uint32 {
	bn := decimalToBinary(int(num))
	var hw uint32
    if bn[31] == 1 {
		hw += 1
	}
	for k, v := range bn {
		if v == 1 {
			hw += 2 << uint(30-k)
		}
	}
	return uint32(hw)
}

func decimalToBinary(num int) []int {
	raw := make([]int, 32)
	index := 0
    for true {
		if num != 0 {
			temp := num
			raw[index] = temp%2
			index++
			num /= 2	
		} else {
			break
		}
	}
    return raw
}
