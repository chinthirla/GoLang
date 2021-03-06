package main
import "fmt"
type Shaper interface {
	Area() int
}
type Rectangle struct {
	width, length int
}
type Square struct {
	side int
}
func (r Rectangle) Area() int {
	return r.width * r.length
}
func (sq Square) Area() int {
	return sq.side * sq.side
}
func main() {
	r := Rectangle{width:5, length:4}
        q := Square{side:10}
	shapesArr := [...]Shaper{r, q}
        fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapesArr {
		fmt.Println("Shape details: ", shapesArr[n])
		fmt.Println("Area of this shape is: ", shapesArr[n].Area())
	}
}


