package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Rows      = 5
	Cols      = 5
	Mine      = -1
	MineCount = 5
)

func placeMines(rows, cols, mineCount int) [][]int {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
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
	return board
}

func setNumbers(board [][]int) {
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

func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			if cell == Mine {
				fmt.Print(" * ")
			} else {
				fmt.Printf(" %d ", cell)
			}
		}
		fmt.Println()
	}
}

func main() {
	board := placeMines(Rows, Cols, MineCount)
	setNumbers(board)
	printBoard(board)
}
