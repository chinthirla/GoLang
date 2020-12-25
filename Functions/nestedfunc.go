package main
import "fmt"
func main() {
	e := 4
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1 + e))
}
