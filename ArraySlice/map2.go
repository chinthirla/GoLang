package main
import "fmt"
func main() {
	elements := make(map[string]string)
	ele := make(map[int]int)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	ele[0] = 12
	ele[1] = 32
	ele[2] = 34
	fmt.Println(ele[0])
	fmt.Println(elements["Li"])
	name, ok := elements["Un"]
	fmt.Println(name, ok)
}
//shorterway to creat map like array
/*elements := map[string]string{
	"H": "Hydrogen",
	"He": "Helium",
	"Li": "Lithium",
	"Be": "Beryllium",
	"B": "Boron",
	"C": "Carbon",
	"N": "Nitrogen",
	"O": "Oxygen",
	"F": "Fluorine",
	"Ne": "Neon",
}
*/

