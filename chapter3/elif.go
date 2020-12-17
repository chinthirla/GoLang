package main
import "fmt"
func main() {
	var n int
	fmt.Print("Enter number :")
	fmt.Scanf("%i", &n)
	if n == 0 {
		fmt.Println("Zero")
	} else if n == 1 {
		fmt.Println("One")
	} else {
		fmt.Println("Enter a valid number")
	}

}

