package main
import "fmt"
func zero(val *int) {
	*val = 0
	fmt.Println(val)
}
func main() {
	val2 := 20
	zero(&val2)
	fmt.Println(val2)
}

