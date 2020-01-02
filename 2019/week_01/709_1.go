package main

import (
	"fmt"
)

func main() {
	fmt.Println(toLowerCase("Hello&HI"))
}

func toLowerCase(str string) string {
	var res []byte
	for _, v := range str {
		if v < 91 && v > 64{
			v += 32
		}
		res = append(res, byte(v))
	}
	return string(res)
} 
