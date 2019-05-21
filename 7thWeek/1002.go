package main

import "fmt"

func main() {
	A := []string{"cool","lock","cook"}
	fmt.Println(commonChars(A))
}

func commonChars(A []string) []string {
	if len(A) == 0 {
		return A
	}
	if len(A) == 1 {
		aa := make([]string, 0)
		for _, v := range A[0] {
			aa = append(aa, string(v))
		}
	}
	
    myStr := make(map[string]int)
	for _, v := range A {
		if len(myStr) == 0 {
			for _, v1 := range A[0] {
				sv, ok := myStr[string(v1)]
				if !ok {
					myStr[string(v1)] = 1
				} else {
					myStr[string(v1)] = sv +1
				}
			}
		} else {
			// myStrTemp := make(map[string]int)
			for k1 := range myStr {
				count := 0
				for _, v11 := range v {
					if k1 == string(v11) {
						count++
					}
				}
				if count < myStr[k1] {
					myStr[k1] = count
				}
			}
		}
	}
	Aa := make([]string, 0)
	for k, v := range myStr {
		for i := 0; i<v; i++ {
			Aa = append(Aa, k)
		}
	}
	return Aa
}