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
	fmt.Scanf("%d %d", &num_x, &num_y)
	var bom_x=2
	var bom_y=3
	
	result := input(num_x, num_y)
	fmt.Println(result)
	explosion := Compare(bom_x, bom_y, num_x, num_y)
	fmt.Println(explosion)
}