// Binary search using recursion
package main
import "fmt"
func BinarySearchRecursive(data []int, low int, high int, value int) int {
	mid := low + (high-low)/2 // To afunc the overflow
	if data[mid] == value {
	return mid
	} else if data[mid] < value {
	return BinarySearchRecursive(data, mid+1, high, value)
	} else {
	return BinarySearchRecursive(data, low, mid-1, value)
	}
}
func main () {
	array := []int {2,65,6,8,9,5,4,55,43,475,45,5745,54,}
	value := 4
	fmt.Println(array,value)
}
