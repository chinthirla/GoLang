package main
import ("fmt"; "math")
func (r rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y1)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func main() {
	r := rectangle{0, 0, 15, 2}
	fmt.Println(r.area())
}
type rectangle struct {
	x1, x2, y1, y2 float64
}
