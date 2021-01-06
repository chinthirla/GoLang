package main
import (
	"fmt"
	"sort"
)
type Person struct {
	Name string
	age int
}
type ByName []Person
func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].age < this[j].age
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func main() {
	kids := []Person{
		{"jill", 9},
		{"jack", 2},
		}
		sort.Sort(ByName(kids))
		fmt.Println(kids)
}
