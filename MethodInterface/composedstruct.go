package main
import "fmt"
type Kitchen struct {
	nomoflamps int
}
type House struct {
	Kitchen
	nomofrooms int
}
func main() {
	h := House{Kitchen{20}, 5}
	fmt.Println("House h has this many lamps :", h.nomoflamps)
	fmt.Println("House h has this many rooms :", h.nomofrooms)
}
