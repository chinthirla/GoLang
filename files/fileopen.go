package main
import (
	"os"
	"fmt"
)
func main() {
	file, err := os.Open("/home/GoLang/Loops/for.go")
	if err != nil {
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
//		fmt.Println(stat)
		return
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

