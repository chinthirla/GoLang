package main
import ("fmt"; "math")
func circlearea(c *Circle) float64 {
	return math.Pi * c.y*c.x
}
func main() {
	c := Circle{9,2}
	fmt.Println(circlearea(&c))
}
type Circle struct {
	x, y float64
}
/*func circleArea(c Circle) float64 {
	return math.Pi * c.r*c.r
}
func main() {
c := Circle{0, 0, 5}
fmt.Println(circleArea(c))
}
type Circle struct {
	x, y, r float64
}*/
