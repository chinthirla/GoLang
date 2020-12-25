package main
import "fmt"
func main() {
	a := 10
	ptr := &a
	*ptr++
	fmt.Println("values of a is:", a)
	fmt.Println("address of a is :", &a)
	fmt.Println("values of ptr is :", *ptr)
	fmt.Println("address of ptr is:", ptr)
}
