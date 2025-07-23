package main

import (
	"fmt"
	"strings"
	"github.com/eiannone/keyboard"
)



func clearScreen() {
	// 画面を50行分の改行で押し流す（擬似クリア）
	fmt.Print(strings.Repeat("\n", 10))
}

func main() {
	
	fmt.Print("レベルを決めてね！\n")
	var level int
	fmt.Scanf("%d", &level)
	rows, cols, MineCount := Level(level)
	

	board, revealed,_ := placeMines(rows, cols, MineCount)
	setNumbers(rows, cols, board)
	// カーソルの位置
	cursorX := 0
	cursorY := 0

	// キーボード入力開始
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		clearScreen()
		printBoard(cursorX, cursorY, board, revealed)
		// 盤面表示
		/*for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				if x == cursorX && y == cursorY {
					fmt.Print("[X]") // カーソル
				} else {
					fmt.Print("[ ]")
				}
			}
			fmt.Println()
		}*/
 
		// キー入力待機
		_ ,key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		// ESCキーで終了
		if key == keyboard.KeyEsc {
			break
		}

		// カーソル移動
		switch key {
		case keyboard.KeyArrowUp:
			if cursorY > 0 {
				cursorY--
			}
		case keyboard.KeyArrowDown:
			if cursorY < rows-1 {
				cursorY++
			}
		case keyboard.KeyArrowLeft:
			if cursorX > 0 {
				cursorX--
			}
		case keyboard.KeyArrowRight:
			if cursorX < cols-1 {
				cursorX++
			}
		case keyboard.KeyEnter:
			if(board[cursorY][cursorX] == 0){
				for i := -1; i < 2; i++{
					for j := -1; j < 2; j++{
						if 0 <= cursorX+i && cursorX+i < cols && 0 <= cursorY+j && cursorY+j < rows {
							revealed[cursorY + j][cursorX + i] = true
						}
					}
				}
			}else{
				revealed[cursorY][cursorX] = true
			}

			explosion := Compare(board, cursorX, cursorY)
		
			mul:=1
			num:=1
			result:=0
			for i:=0 ;i<rows;i++{
				for j:=0;j<cols;j++{
					if (revealed[i][j] && board[i][j]!=-1){
						num=1
					}else{
						num=0
					}
					mul = mul*num
				}
			}
			if(mul==1){
				result = 1
			}else if(explosion == true){
				result = -1
			}

			if(result == -1){
				clearScreen()
				printBoard(-1, -1, board, revealed)
				fmt.Print("ドカーン！\n")

				return 

			}else if(result == 1){
				fmt.Print("せいこう！\n")
				return 
			}
		case keyboard.KeyEsc:
			return
	}
	}
}





