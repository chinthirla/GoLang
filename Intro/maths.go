package main
import (
	"fmt"
	m "/home/GoLang/Intro/math"
)
func main() {
	cs := [...]float64{1,2,3,45,6,7}
	avg := m.Average(cs)
	fmt.Println(avg)
}
