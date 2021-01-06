package main
import (
	"fmt"
	"io/ioutil"
	"hash/crc32"
)
func main() {
	h1, err := getHash("/home/GoLang/select/select1.go")
	if err != nil {
		return
	}
	h2, err := getHash("/home/GoLang/select/select2.go")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}
func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

