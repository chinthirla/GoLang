package main
import "fmt"
func main() {
	x := [6]string{"a","b","c","d","e","f"}
	fmt.Println(x[2:5], "\n")
	fmt.Println(x[2:], "\n")
	fmt.Println(x[:], "\n")
	fmt.Println(x[:4])
}
