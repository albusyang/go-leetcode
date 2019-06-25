package main

import "fmt"

func main() {
	fmt.Println(addBinary("111", "111"))
}

func addBinary(a string, b string) string {
	adv := "0"
	rArr := make([]byte, 0)
	if len(a)-1 > len(b)-1 {
		return exec(a, b, adv, rArr)
	} else {
		return exec(b, a, adv, rArr)
	}
}

func exec(a, b, adv string, rArr []byte) string {
	for i := 0; i < len(a); i++ {
		rArr = append(rArr, a[len(a)-1-i])
	}
		
	for i := 0; i < len(b); i++ {
		if adv == "1" {
			switch rArr[i] + b[len(b) - 1-i] {
			case 98:
				adv = "1"
			case 97:
				adv = "1"
				rArr[i] = 48
			case 96:
				adv = "0"
				rArr[i] = 49
			}
		} else {
			switch rArr[i] + b[len(b) - 1-i] {
			case 98:
				adv = "1"
				rArr[i] = 48
			case 97:
				rArr[i] = 49
			}
		}
	}
		
	if adv == "1" {
		for i := 0; i < len(a)-len(b); i++{
			if adv == "1" {
				switch rArr[len(b)+i] + 49 {
				case 98:
					adv = "1"
					rArr[len(b)+i] = 48
				case 97:
					adv = "0"
					rArr[len(b)+i] = 49
				}
			}
		}
	}
	
	if adv == "1" {
		rArr = append(rArr, 49)
	}
	
	rss := ""
	for i := 0; i < len(rArr); i++ {
		rss = rss + string(rArr[len(rArr)-1-i])
	}
	return rss
}