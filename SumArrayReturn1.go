package main
import "fmt"
func SumArray(data []int) int {
	size := len(data)
	total := 0
	for index := 0; index < size; index++ {
	total = total + data[index]
	}
	return total
}
func main() {
	numb := []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(SumArray(numb))
}
