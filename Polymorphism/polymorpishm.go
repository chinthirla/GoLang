package main
import "fmt"
type Hobby interface {
	myStereotype() string
}
type Human struct {
}
type Man struct {
	Human
}
type Woman struct {
	Human
}
type Dog struct {
}
func (m Man) myStereotype() string {
	return "i am a man and i like hunting"
}
func (w Woman) myStereotype() string {
	return "i am a woman and i like shpooing"
}
func (h Human) myStereotype() string {
    return "I'm a Human, only an abstract concept, and I can have no hobby."
}
func (d Dog) myStereotype() string {
	return "bow bow bow, I'm chasing sticks."
}
func main() {
	m := new(Man)
	w := new(Woman)
	h := new(Human)
	d := new(Dog)
	arr := [...]Hobby{m, w, h, d}
	for n, _ := range (arr) {
		fmt.Println("My hobby?  Well,", arr[n].myStereotype())
	}
}
