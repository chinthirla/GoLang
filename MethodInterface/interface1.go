package main
import "fmt"
type Shaper interface {
	Area() int
}
type Rectangle struct {
	width, length int
}
func (r Rectangle) Area() int {
	return r.width * r.length
}
func main() {
	r := Rectangle{width:4, length:5}
	fmt.Println("value of r:", r)
	fmt.Println("Vlue of area:", r.Area())
	s := Shaper(r)
	fmt.Println("Vlaue of shaper interface:", s.Area())
}



