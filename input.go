package main
import "fmt"


func input(x int, y int) bool{
	if (0 <=x && x<= 4) && (0 <= y && y <= 4){
		return true
	} else {
		fmt.Println("Input is invalid.")
		return false
	}
}

func Compare(bom_x int, bom_y int, currentX int, currentY int) bool{
	if (bom_x == currentX) && (bom_y == currentY) {
		return true
	}
	return false
}
