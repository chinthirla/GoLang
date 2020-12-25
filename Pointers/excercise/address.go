package main
import "fmt"
func zero(xptr *int) {
	*xptr = 30
	fmt.Println(&xptr)
}
func main() {
	x := new(int)
	*x = 40
	fmt.Println(*x)
	zero(x)
	fmt.Println(&x)
	fmt.Println("Value of the variables is :", *x)
	fmt.Println("address of variable is :", x)
}
