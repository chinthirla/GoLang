package main
import "fmt"
func main() {
	i := 0
	for {
		if i >= 3 {
			break
		}
		i++
		fmt.Println("value of i is:", i)
	}
	fmt.Println("A statement just after for loop.")
}
