package main

import "fmt"

func main() {
    fmt.Println(findComplement(5))
}

func findComplement(num int) int {
	rn := decimalToBinary(num)
	bn := make([]int, len(rn))
	for k, v := range rn{
		if v == 0 {
			bn[k] = 1		
		} else {
			bn[k] = 0
		}
	}	

	var hw int
    if bn[len(bn)-1] == 1 {
		hw += 1
	}
	for k, v := range bn {
		if v == 1 {
			hw += 2 << uint(len(bn)-k-2) //why '-2': bn list max index is 'len(bn)-1' & base is '2'
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
