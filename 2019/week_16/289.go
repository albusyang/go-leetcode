package main

import "fmt"

func main() {
	board := [][]int{{0,1,0},{0,0,1},{1,1,1},{0,0,0}}
	gameOfLife(board)
	fmt.Println(board)
}

func gameOfLife(board [][]int)  {
	result := [][]int{}
	//result = board 
	vv := len(board)
	hv := len(board[0])
	for m := 0; m < vv; m++ {
		ms := make([]int, 0)
		for n := 0; n < hv; n++ {
			ms = append(ms, board[m][n])
		}
		result = append(result, ms)
		for n := 0; n < hv; n++ {
			lives := 0
			// 正左方
			if n-1 >= 0 {
				lives += board[m][n-1]
			}
			// 正右方
			if n+1 < hv {
				lives += board[m][n+1]
			}
			// 上面一排
			if m-1 >= 0 {
				if n-1 >= 0 {
					// 左上角
					lives += board[m-1][n-1]
				}
				// 正上方
				lives += board[m-1][n]
				
				if n+1 < hv {
					// 右上角
					lives += board[m-1][n+1]
				}
			}
			// 下面一排
			if m+1 < vv {
				if n-1 >= 0 {
					// 左下角
					lives += board[m+1][n-1]
				}
				// 正下方
				lives += board[m+1][n]
				
				if n+1 < hv {
					// 右下角
					lives += board[m+1][n+1]
				}
			}
		
			if board[m][n] == 1 {
				if lives < 2 || lives > 3 {
					result[m][n] = 0
				}
			}
			if board[m][n] == 0 {
				if lives==3 {
					result[m][n] = 1
				}
			}
		}
	}

	for m := 0; m < vv; m++ {
		for n := 0; n < hv; n++ {
			board[m][n] = result[m][n]
		}
	}
}