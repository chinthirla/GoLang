package main
import "fmt"
func main() {
	i := 0
	fmt.Printf("Enter a value : ")
	fmt.Scanf("%d", &i)
	fmt.Println(fib(i))
}
func fib(n int) int {
	x := 0
	y := 1
	temp := 0
	for i := 0; i <= n; i++ {
		temp += x
		x += y
		temp = x
		fmt.Println(temp)
	}
	return temp
}
