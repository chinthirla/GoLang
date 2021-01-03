package main
import "fmt"
type Perimeter interface {
	Sahpe() int
}
type Rectangle struct {
	width, length int
}
type Circle struct {
	radius int
}
func (r Rectangle) Shape() int {
	return r.width * r.length
}
func (c Circle) Shape() int {
	return 3 * c.radius * c.radius
}
func main() {
	r := Rectangle{width:7, length:10}
	fmt.Println("Value of Rectangle:", r)
	fmt.Println("Sahpe of Rectangle is:", r.Shape())
	c := Circle{radius:5}
	fmt.Println("vlaue of circle is :", c)
	fmt.Println("value of radius is:", c.Shape())
}

