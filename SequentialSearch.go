package main
import "fmt"
func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
	if value == data[i] {
	return true
	}
	}
	return false
}
func main() {
	array := []int {1,2,3,4,5,6,7,8,9,10}
	value := 5
	fmt.Println(SequentialSearch(array, value))
}
