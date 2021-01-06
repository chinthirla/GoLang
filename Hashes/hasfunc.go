package main
import (
	"fmt"
	"hash/crc32"
)
func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("vijay"))
	v := h.Sum32()
	fmt.Println(v)
}
