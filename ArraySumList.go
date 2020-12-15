package main
import "fmt"
func main() {
	list := []int {1,2,3,4,5,6,7,8,9,10}
	sum := 0
	for i := 0; i<len(list); i++ {
	sum += list[i]
	}
	fmt.Println(sum)
}
