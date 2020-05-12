package main

import "fmt"

func myPow(x float64, n int) float64 {
	var result float64
	if n < 0 {
		result = 1.0 / calculate(x, -n)
	} else {
		result = calculate(x, n)
	}
	return result
}

func calculate(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	var y float64
	y = calculate(x, n/2)

	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {
	fmt.Println(fmt.Sprintf("%f",myPow(2, -100)))
}
