package main
import "fmt"
type Human interface {
	 myStereotype() string
}
type Man struct {
}
type Woman struct {
}
func (m Man)  myStereotype() string {
	return "i like hunting bcz i am man"
}
func (w Woman)  myStereotype() string {
	return "I like shpooing bcz i am woman"
}
func main() {
	m := new(Man)
	w := new(Woman)
	arr := [...]Human{m, w}
	for n, _ := range (arr) {
		fmt.Println("I'm a human, and my stereotype is: ", arr[n].myStereotype())
	}
}
