package main

import (
		"fmt"
		"math/rand"
		"time"
)

// --- 設定項目 ---
const (
	Rows      = 5
	Cols      = 5
	Mine      = -1
	MineCount = 5
)

func placeMines(rows, cols, mineCount int) [][]int {
	// 盤面の初期化（すべて0）
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}

	// ランダムに地雷を配置
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

func printBoard(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			if cell == Mine {
				fmt.Print(" * ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
}

func main() {
	board := placeMines(Rows, Cols, MineCount)
	printBoard(board)
}