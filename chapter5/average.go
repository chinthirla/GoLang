package main
import "fmt"
func average(xs []float64) float64 {
	panic("Not Implemented")
}
func main() {
	x := []int{1,2,3,4,5}
	var total int
	for _, avg := range x {
		total += avg
		}
		fmt.Println(total / len(x))
	}
