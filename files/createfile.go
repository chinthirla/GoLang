package main
import (
	"os"
//	"fmt"
)
func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("This is test file")
}

