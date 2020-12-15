package main
import "fmt"
func main() {
	array := []int {1,2,3,4,5,6,7}
	k := 3
//	z := Rotate(array,k)
	fmt.Println(Rotate(array,k))
}
func Rotate(data []int, k int)[]int  {
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)
	return data
}
func ReverseArray(data []int, start int, end int)[]int  {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	return data
}
