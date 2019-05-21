package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {
    rArr := make([]int, 0)
	for ii := left; ii <= right; ii++ {
		if  ii < 10 {
			rArr = append(rArr, ii)
			continue
		}
		if ii == 0 {
			continue
		}
		i := ii
		if i >= 10000 {
			k := i / 10000
			if ii % k != 0 {
				continue
			}
			i %= 10000
			if i < 1000 {
				continue
			}
		}
		if i >= 1000 {
			k := i / 1000
			if ii % k != 0 {
				continue
			}
			i %= 1000
			if i < 100 {
				continue
			}
		}
		if i >= 100 {
			k := i / 100
			if ii % k != 0 {
				continue
			}
			i %= 100
			if i < 10 {
				continue
			}
		}
		if i >= 10 {
			k := i / 10
			if ii % k != 0 {
				continue
			}
			i %= 10
		}
		if i == 0 {
			continue
		}
		if ii % i != 0 {
			continue
		}
		rArr = append(rArr, ii)
	}
	return rArr
}