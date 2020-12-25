package main
import "fmt"
func zero(xptr *int) {
	*xptr = 20
}
func main() {
	xs := new(int)
	zero(xs)
	fmt.Println(*xs)
}
