package main
import "fmt"
func swap(ptr *int,ptr1 *int) {
	*ptr = 2
	*ptr1 = 1
}
func main() {
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println("value of a:", a)
	fmt.Println("value of b is:", b)
}

