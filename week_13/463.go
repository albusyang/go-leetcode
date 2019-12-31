package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}
	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	pm := 0
    for i, vi := range grid {
		for j, vj := range vi {
			if vj == 1 {
			if i == 0 || grid[i-1][j] == 0 {
					pm += 1
				}
				if j == 0 || grid[i][j-1] == 0 {
					pm += 1
				}
			}
		}
	}
	return pm << 1
}