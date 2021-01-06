package main
import (
	"path/filepath"
	"os"
	"fmt"
)
func main() {
	filepath.Walk("/home/GoLang/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}


