package main
import "fmt"
type Bus struct {
	l, b, h int
	row, seatperrow int
}
type Cuboider interface {
	CubicVolume() int
}
type PublicTransporter interface  {
    PassengerCapacity() int
}
type PersonalSpaceLaw interface {
	    IsCompliantWithLaw() bool
}
func (b Bus) IsCompliantWithLaw() bool {
	 return (b.l * b.b * b.h) / (b.row * b.seatperrow) >= 3
}
func (bus Bus) CubicVolume() int {
	 return bus.l *  bus.b * bus.h
}
func (bus Bus)  PassengerCapacity() int {
	return bus.row * bus.seatperrow
}
func main() {
	b := Bus{
		l:4, b:6, h:10,
		row:12, seatperrow:2}
	fmt.Println("Cubic volume of bus:", b.CubicVolume())
	fmt.Println("PassengerCapacity of bus:", b.PassengerCapacity())
	fmt.Println("PassengerCapacity of bus:", b.IsCompliantWithLaw())
}

