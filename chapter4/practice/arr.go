package main
import "fmt"
func main() {
	x := []float64{1,2,3,4,5}
	slice1 := make([]float64,4)
	copy(slice1, x)
	fmt.Println(slice1)
	fmt.Println(slice1[2])
}
