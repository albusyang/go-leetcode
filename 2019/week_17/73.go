package main

import "fmt"

func main() {
	mm := [][]int{{1,1,1},{1,0,1},{1,1,1}}
  setZeroes(mm)
  fmt.Println(mm)
}

func setZeroes(matrix [][]int)  {
    xArr := make([]int, 0)
	yArr := make([]int, 0)
	
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == 0 {
				xArr = append(xArr, x)
				yArr = append(yArr, y)
			}
		} 
	}
	
	for i := range matrix {
		for j := range matrix[i] {
			for _, y := range yArr {
				if y == i {
					matrix[i][j] = 0
				}
			}
			for _, x := range xArr {
				if x == j {
					matrix[i][j] = 0
				}
			}
		}
	}
}