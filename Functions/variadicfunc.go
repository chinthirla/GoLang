package main
import "fmt"
func main() {
	fmt.Println(add(1,2,3,4))
}
func add(args ...int) int {
	total := 0
	for _, c := range args {
		total += c
	}
		return total
	}

//func Println(a ...interface{}) (n int, err error)
