/*Write a function which takes an integer and
halves it and returns true if it was even or false
if it was odd. For example half(1) should return
(0, false) and half(2) should return (1,
true).*/
package main
import "fmt"
func half1(x int) int {
	xs := int( x / 2 )
	return xs
}
func main() {
	var num int
	fmt.Printf("Enter number :")
	fmt.Scanf("%d", &num)
	str := half1(num)
	if str % 2 == 0 {
		fmt.Println("(", str, "False )")
	} else {
		fmt.Println("(", str, "True )")
	}
}

