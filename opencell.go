package main



func opencell(board [][]int, revealed [][]bool, x, y, rows, cols int) {
	// 範囲外チェック（y: 行, x: 列）
	if y < 0 || y >= rows || x < 0 || x >= cols {
		return
	}

	// すでに開いていたら終了
	if revealed[y][x] {
		return
	}

	// 開ける
	revealed[y][x] = true

	// 0なら周囲を再帰的に開ける
	if board[y][x] == 0 {
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

