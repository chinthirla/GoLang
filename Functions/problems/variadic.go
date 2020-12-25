/*Write a function with one variadic parameter
that finds the greatest number in a list of numbers.*/
package main
import "fmt"
func arr(x ...int) int {
	temp := x[0]
	fmt.Println(x)
	var xs int
	for _, xs := range x {
		if xs < temp {
		   temp = xs
		}
	}
	return xs
}
func main(){
	array := []int{
		48, 96, 86, 68,
	        57, 82, 63, 70,
	        37, 34, 83, 27,
	        19, 97, 9, 17,
		}
		fmt.Println("smalest number is : ", arr(array...))
	}
