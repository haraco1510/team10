package main
import (
	"fmt"
)
func main(){
	var (
		Rows      = 5
		Cols      = 5
		MineCount = 5
	)

	board, revealed := placeMines(Rows, Cols, MineCount)
	setNumbers(Rows, Cols, board)
	printBoard(board, revealed)

	var num_x, num_y int
	var result int
	
	fmt.Print("Start!")

	for{
		fmt.Print("数字を入力してね！：")
		fmt.Scanf("%d %d", &num_x, &num_y)
		inputRange := input(num_x, num_y)
		if(inputRange == false){
			continue
		}
		revealed[num_x][num_y] = true
		explosion := Compare(board, num_x, num_y)
		if(explosion == true){
			result = -1
			break
		}
		mul=1
		for i:=0 ;i<Rows;i++{
			for j:=0;j<Cols;j++{
				mul = mul*revealed[i][j]
			}
		}
		if(mul){
			result = 1
			break
		}
		//fmt.Println(explosion)
		printBoard(board, revealed)
	}

	if(result == -1){
		fmt.Print("fail")
	}else{
		fmt.Print("success")
	}
	
	
}