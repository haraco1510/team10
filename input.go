package main

import "fmt"

func input(revealed [][]bool, rows, cols, x, y int) bool{
	if (0 <= x && x <= rows-1) && (0 <= y && y <= cols-1) && (revealed[x][y] == false){
		return true
	} else {
		fmt.Println("Input is invalid.")
		return false
	}
}

func Compare(board [][]int, x, y int) bool{
	if board[x][y] == Mine {
		return true
	}
	return false
}
