package main
import (
	"fmt"
	)
func sum(array []int) int {
	result := 0
	for _, v := range array {
	result += v
	}
	return result
}
func main() {
	numb := []int{0,1, 2, 3, 4, 5,6,7,8,9,10}
	fmt.Print("Result :", sum(numb),"\n")
}
