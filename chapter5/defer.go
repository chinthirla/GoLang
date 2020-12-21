package main 
import "fmt"
func first() {
	fmt.Println("This is first function")
}
func second() {
	fmt.Println("This is second function")
}
func main() {
	defer second()
	first()
}
