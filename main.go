package main

import (
	"bufio"
	"fmt"
	"os"
  "strconv"
  "strings"
)

func main(){
	var Level int
	var num_x, num_y int
	var result int
	var err error
	
	fmt.Print("Start!\n")
	for{
		fmt.Print("難易度を選んでね(1~3で)：")
		fmt.Scanf("%d", &Level)
		if(1<=Level&&Level<=3){
			break
		}	
	}
	
	Rows, Cols, MineCount := selectLevel(Level)
	
	board, revealed, flagged := placeMines(Rows, Cols, MineCount)
	scanner := bufio.NewScanner(os.Stdin)
	setNumbers(Rows, Cols, board)
	printBoard(board, revealed, flagged)

	fmt.Println("~説明~")
	fmt.Println("爆弾に旗を立てたいときは、'f' の後に座標を入力（例: f 1 2）")
	fmt.Println("マスを開きたいときは、座標だけ入力（例: 1 2）")

	for{
		fmt.Print("座標を入力してね！：")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		if len(parts) != 2 && len(parts) != 3 {
			fmt.Println("入力形式が違うよ！（例: 1 0 または f 1 0）")
			continue
		}

		flagMode := false
		if len(parts) == 3 && strings.ToLower(parts[0]) == "f" {
				flagMode = true
				num_x, err = strconv.Atoi(parts[1])
				if err != nil {
						fmt.Println("数字の入力が不正だよ")
						continue
				}
				num_y, err = strconv.Atoi(parts[2])
				if err != nil {
						fmt.Println("数字の入力が不正だよ")
						continue
				}
		} else if len(parts) == 2 {
				num_x, err = strconv.Atoi(parts[0])
				if err != nil {
						fmt.Println("数字の入力が不正だよ")
						continue
				}
				num_y, err = strconv.Atoi(parts[1])
				if err != nil {
						fmt.Println("数字の入力が不正だよ")
						continue
				}
		} else {
				fmt.Println("入力形式が違うよ！（例: 1 0 または f 1 0）")
				continue
		}

		inputRange := input(revealed,Rows,Cols,num_x, num_y)
		if(inputRange == false){
			continue
		}

		if flagMode {
			flagged[num_x][num_y] = !flagged[num_x][num_y]
			printBoard(board, revealed, flagged)
		} else {
			if(board[num_x][num_y] == 0){
				opencell(board, revealed, num_x, num_y, Rows, Cols)
			}else{
				revealed[num_x][num_y] = true
			}

			explosion := Compare(board, num_x, num_y)
			printBoard(board, revealed, flagged)
			if explosion {
					result = -1
					break
			}
		}

		
		mul:=1
		num:=1
		for i:=0 ;i<Rows;i++{
			for j:=0;j<Cols;j++{
				if(board[i][j] != -1){
					if (revealed[i][j]){
						num=1
					}else{
						num=0
					}
				}

				mul = mul*num
			}
		}
		if(mul==1){
			result = 1
			break
		}
		//fmt.Println(explosion)
		
	}

	if(result == -1){
		fmt.Print("ドカーン！")
	}else{
		fmt.Print("せいこう！")
	}
	fmt.Print("\n")
	
	
}