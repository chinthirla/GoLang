package main
import ("fmt"; "math")
func circlearea(c *Circle) float64 {
	return math.Pi * *c**c
}
func main() {
	c := Circle{5}
	fmt.Println(circlearea(&c))
}
type Circle struct {
	x, y, r float64
}
