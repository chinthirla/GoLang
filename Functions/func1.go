package main
import "fmt"
func average(x[]float64) float64 {
//	x := []int{1,2,3,4,5}
	total := 0.0
	for _, avg := range x {
		total += avg
	}
	return total / float64(len(x))
}
func main() {
	v := []float64{1,2,3,4,5}
	fmt.Println(average(v))
}


