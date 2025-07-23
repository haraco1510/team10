package main

import "fmt"

func input(revealed [][]bool, rows, cols, x, y int) bool{
	if (0 <= x && x <= rows-1) && (0 <= y && y <= cols-1){
		if(revealed[x][y] == false){
			return true
		}else{
			fmt.Println("入力済みだよ(^^♪")
			return false
		}
		
	} else {
		fmt.Println("範囲外だよ( ﾉД`)ｼｸｼｸ…")
		return false
	}
}

func Compare(board [][]int, x, y int) bool{
	if board[y][x] == Mine {
		return true
	}
	return false
}




func Level(Level int) (int, int, int){
	if(Level == 1){
		return 3, 3, 2
	}else if(Level == 2){
		return 7, 7, 7
	}else{
		return 10, 10, 15
	}
}
