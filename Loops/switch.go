package main
import "fmt"
func main() {
	fmt.Println("Enter a number :")
	var num int64
	fmt.Scanf("%i", &num)
	switch num  {
	case 0 : fmt.Println("Zero")
	case 1 : fmt.Println("One")
	case 2 : fmt.Println("Two")
	default : fmt.Println("not valid number")
}
}
