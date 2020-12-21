package main
import "fmt"
func f(x int) {
	fmt.Println(x)
}
func main() {
	c := 4
	v := 6
	f(c + v)
}
