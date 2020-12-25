//sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?
package main
import "fmt"
func sum(x ...int) int {
	total := 0
	for _, xs := range x {
	total += xs
}
	return total
}
func main() {
	arr := []int{1,2,3,4,5,6,7,8}
	slice := make([]int, 4)
	copy(slice, arr)
	fmt.Println(sum(slice...))
}

