package main
import (
	"fmt"
	"io/ioutil"
)
func main() {
	file, err := ioutil.ReadFile("fileopen.go")
	if err != nil {
		return
	}
//	defer file.Close()
	bs := string(file)
	fmt.Println(bs)
}
