package main
import "fmt"
func main() {
//	var i int64
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}
