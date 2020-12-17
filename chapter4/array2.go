package main
import "fmt"
func main() {
	var x [5]float64
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 90
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
//	fmt.Println(total / 5)
	}
	fmt.Println(total / float64(len(x)))
}
