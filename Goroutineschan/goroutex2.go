package main
import "fmt"
func f(n int) {
	for i := 0; i<=10 ; i++ {
//		fmt.Println(n, ":", i)
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}
