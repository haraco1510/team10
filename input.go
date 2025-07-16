package main

import "fmt"

var Mine = -1

func input(x int, y int) bool{
	if (0 <= x && x <= 4) && (0 <= y && y <= 4){
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
