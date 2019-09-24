package main

import "fmt"

func main() {
	bills := []int{5,5,5,10,20}
	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
    countF, countT, countTW := 0, 0, 0
	for _, v := range bills {
		switch v {
			case 5:
			countF++
			case 10:
			if countF >= 1 {
				countT++
				countF--
			} else {
				return false
			}
			case 20:
			if countF >= 1 && countT >= 1 {
				countTW++
				countF--
				countT--
			} else if countF >= 3 {
				countTW++
				countF -= 3
			} else {
				return false
			}
		}
	}
	return true
}