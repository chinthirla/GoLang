package main
import "fmt"
func main() {
	xs := []int{1,2,3,4,5}
	fmt.Println(add(xs ...))
}
func add(c ...int) int {
	total := 0
	for _, fs := range c {
		total += fs 
	}
	return total
}
