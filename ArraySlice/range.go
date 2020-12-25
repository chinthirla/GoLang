package main
import "fmt"
func main() {
	var total float64
//	var value float64
	x := [5]float64{10,22,32,42,44}
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
