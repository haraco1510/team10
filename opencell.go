package main



func opencell(board [][]int, revealed [][]bool, x, y, rows, cols int) {
	// 範囲外チェック
	if x < 0 || x >= rows || y < 0 || y >= cols {
		return
	}
	// すでに開いていたら終了
	if revealed[x][y] {
		return
	}

	// 開ける
	revealed[x][y] = true

	// 開けたマスが0なら周囲も再帰で開く
	if board[x][y] == 0 {
		directions := [8][2]int{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, -1},           {0, 1},
			{1, -1},  {1, 0},  {1, 1},
		}
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			opencell(board, revealed, nx, ny, rows, cols)
		}
	}
}
