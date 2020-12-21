package main
import "fmt"
func fact(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}
func main() {
	var c uint
	fmt.Printf("Enter number to get factorial :")
	fmt.Scanf("%d", &c)
	fmt.Println("factorial of given number is :",fact(c))
}

