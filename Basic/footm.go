package main
import "fmt"
func main() {
	var x float64
	fmt.Print("Enter value in feet :")
	fmt.Scanf("%f", &x)
	x *= 0.3048
	fmt.Println("Value in meters is : ", x)
}

