package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Mine = -1


func placeMines(rows, cols, mineCount int) ([][]int, [][]bool) {
	board := make([][]int, rows)
	revealed := make([][]bool, rows)

	for i := range board {
		board[i] = make([]int, cols)
		revealed[i] = make([]bool, cols) 
	}

	rand.Seed(time.Now().UnixNano())
	placed := 0
	for placed < mineCount {
		r := rand.Intn(rows)
		c := rand.Intn(cols)
		if board[r][c] != Mine {
			board[r][c] = Mine
			placed++
		}
	}

	return board, revealed
}


func setNumbers(Rows int, Cols int, board [][]int) {
	dirs := [8][2]int{{-1,-1}, {-1,0}, {-1,1}, {0,-1}, {0,1}, {1,-1}, {1,0}, {1,1}}
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if board[r][c] == Mine {
				continue
			}
			count := 0
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < Rows && nc >= 0 && nc < Cols {
					if board[nr][nc] == Mine {
						count++
					}
				}
			}
			board[r][c] = count
		}
	}
}

func printBoard(cursorX int, cursorY int, board [][]int, revealed [][]bool) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if c == cursorX && r == cursorY {
				if revealed[r][c] {
					if board[r][c] == Mine {
						fmt.Print("[*]")
					} else {
						fmt.Printf("[%d]", board[r][c])
					}
				} else {
					fmt.Print("[■]")
				}
			}else{
				if revealed[r][c] {
					if board[r][c] == Mine {
						fmt.Print(" * ")
					} else {
						fmt.Printf(" %d ", board[r][c])
					}
				} else {
					fmt.Print(" ■ ")
				}
			}
			
		}
		fmt.Println()
	}
}
