package main

import "fmt"

var Mine2 = -1

func input(x int, y int) bool{
	if (0 <= x && x <= 4) && (0 <= y && y <= 4){
		return true
	} else {
		fmt.Println("Input is invalid.")
		return false
	}
}

func Compare(board [][]int, x, y int) bool{
	if board[y][x] == Mine2 {
		return true
	}
	return false
}

func Level(level int)(int, int, int){
	if(level == 1){
		return 3,3,2
	}else if(level == 2){
		return 7,7,7
	}else{
		return 10,10,15
	}

}

