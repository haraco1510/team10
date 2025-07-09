package main
import (
	"fmt"
)
func main(){
	x := 2
	y := 3
	bom_x := 2
	bom_y := 3

	result := input(x, y)
	fmt.Println(result)
	explosion := Compare(bom_x, bom_y, x, y)
	fmt.Println(explosion)
}