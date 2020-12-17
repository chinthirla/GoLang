package main
import "fmt"
func main() {
	elements :=map[string]map[string]string{
		"H" : map[string]string{
			"name" : "hydrogen",
			"state" : "gas",
		},
		"Li" : map[string]string{
			"name" : "Lithium",
			"state" : "gas",
		},
		"He" : map[string]string{
			"name" : "Helium",
			"state" : "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
	if ele, ok := elements["He"]; ok {
		fmt.Println(ele["name"], ele["state"])
	}
}



