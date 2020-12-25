package main
import "fmt"
type Kitchen struct {
	nomlamps int
	table int
}
type Bedrooms struct {
	nomrooms int
	table int
}
type House struct {
	Kitchen
	Bedrooms
}
func main() {
	h := House{Kitchen{5, 20}, Bedrooms{10, 15}}
	fmt.Println("Number of lamps in house: ", h.nomlamps)
	fmt.Println("Number of rooms in house: ", h.nomrooms)
	fmt.Println("Number of tables in house: ", h.Kitchen.table)
	fmt.Println("Number of tables in house: ", h.Bedrooms.table)
}
