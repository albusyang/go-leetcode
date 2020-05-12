package main

import (
	"fmt"
)

func main() {
	arr := []int{9,9,9}
	fmt.Println(plusOne(arr))
}

func plusOne(digits []int) []int {
    dl := len(digits) - 1
	adv := 1
	for adv > 0 {
		if digits[dl] < 9 {
			digits[dl] += adv
			adv = 0
			return digits
		} else {
			if dl == 0 {
				newDs := digits[:]
				newDs[dl] = 1
				newDs = append(newDs, 0)
				adv = 0
				return newDs
			}
			digits[dl] = 0
		}
		dl--
	}
    return digits
}