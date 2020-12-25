package main
import "fmt"
type Cameraphone struct {
	Phone
	Camera
}
type Phone struct { }
func (_ Phone) call() string {
	return  "Ring Ring"
}
type Camera struct { }
func (_ Camera) takepic() string {
		return "Click"
	}
func main() {
	cp := new (Cameraphone)
	fmt.Println("starting step of multi interface")
	fmt.Println("Take pictures : ", cp.takepic())
	fmt.Println("Call the phone :", cp.call())
}
