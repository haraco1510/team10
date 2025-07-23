package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Mine = -1


func placeMines(rows, cols, mineCount int) ([][]int, [][]bool, [][]bool) {
	board := make([][]int, rows)
	revealed := make([][]bool, rows)
	flagged := make([][]bool, rows)

	for i := range board {
		board[i] = make([]int, cols)
		revealed[i] = make([]bool, cols) 
		flagged[i] = make([]bool, cols)
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

	return board, revealed, flagged
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

func printBoard(board [][]int, revealed [][]bool, flagged [][]bool) {
	cols := len(board[0])
	rows := len(board)

	fmt.Print("   ")
	for c := 0; c < cols; c++ {
		fmt.Printf(" %d ", c)
	}
	fmt.Println()

	fmt.Print("   ")
	for c := 0; c < cols; c++ {
		fmt.Print("---")
	}
	fmt.Println()

	for r := 0; r < rows; r++ {
		fmt.Printf(" %d|", r)
		for c := 0; c < cols; c++ {
			if revealed[r][c] {
				if board[r][c] == Mine {
					fmt.Print(" * ")
				} else {
					fmt.Printf(" %d ", board[r][c])
				}
			} else if flagged[r][c] {
				fmt.Print(" x ")
			} else {
				fmt.Print(" ■ ")
			}
		}
		fmt.Println()
	}
}


