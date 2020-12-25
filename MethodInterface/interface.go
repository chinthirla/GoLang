package main
import "fmt"
type shape interface {
	area() float64
	perimeter() float64
}
type rect struct {
	width float64
	length float64
}
func (r rect) area() float64 {
	return r.width * r.length
}
func (r rect) perimeter() float64 {
	return 2 * (r.width * r.length)
}
func main() {
        var s shape
	s = rect{5.0, 4.0}
	r := rect{6.0, 2.0}
	fmt.Printf("type of s is : %T \n", s)
	fmt.Printf("value of s is %v:\n ", s)
	fmt.Println("area of rectanglr s:", s.area())
	fmt.Println("s == r is :", s == r)
}
