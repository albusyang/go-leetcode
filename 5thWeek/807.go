package main

import "fmt"

func main() {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
	//fmt.Println(grid[1][2])
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
    count := 0
	lArr := make([]int, 0)
	rArr := make([]int, 0)
	// 获取左右边(每一行最大的值)
	for l := 0; l < len(grid); l++ {
		max := 0
		for r := 0; r < len(grid); r++ {
			if grid[l][r] > max {
				max = grid[l][r]
			}
		}
		lArr = append(lArr, max)
	}
	// 获取上下底
	for r := 0; r < len(grid); r++ {
		max := 0
		for l := 0; l < len(grid); l++ {
		if grid[l][r] > max {
				max = grid[l][r]
			}
		}
		rArr = append(rArr, max)
	}
	
	for l := 0; l < len(grid); l++ {
		for r := 0; r < len(grid); r++ {
			if lArr[l] > rArr[r] {
				count += rArr[r] - grid[l][r]
			} else {
				count += lArr[l] - grid[l][r]
			}
		}
	}
	
	return count
}