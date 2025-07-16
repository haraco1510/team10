package main
import (
	"fmt"
)
func main(){
	var Level int
	var num_x, num_y int
	var result int
	
	fmt.Print("Start!\n")
	for{
		fmt.Print("難易度を選んでね(1~3で)：")
		fmt.Scanf("%d", &Level)
		if(1<=Level&&Level<=3){
			break
		}	
	}
	
	Rows, Cols, MineCount := selectLevel(Level)
	
	board, revealed := placeMines(Rows, Cols, MineCount)
	setNumbers(Rows, Cols, board)
	printBoard(board, revealed)

	for{
		fmt.Print("座標を入力してね！：")
		fmt.Scanf("%d %d", &num_x, &num_y)
		inputRange := input(revealed,Rows,Cols,num_x, num_y)
		if(inputRange == false){
			continue
		}

		if(board[num_x][num_y] == 0){
			for i := -1; i < 2; i++{
				for j := -1; j < 2; j++{
					if 0 <= num_x+i && num_x+i < Cols && 0 <= num_y+j && num_y+j < Rows {
						revealed[num_x + i][num_y + j] = true
					}
				}
			}
		}else{
			revealed[num_x][num_y] = true
		}


		explosion := Compare(board, num_x, num_y)

		printBoard(board, revealed)
		
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
		
		if(explosion == true){
			result = -1
			break
		}
	}

	if(result == -1){
		fmt.Print("ドカーン！")
	}else{
		fmt.Print("せいこう！")
	}
	fmt.Print("\n")
	
	
}