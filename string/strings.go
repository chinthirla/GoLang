package main
import (
	"fmt"
	"strings"
)
func main() {
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(arr)
	fmt.Println(str)
	fmt.Println(
		strings.Contains("test", "es"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaa", "a", "b", 2),
		strings.Split("a-s-d-g-f", "-"),
		strings.ToLower("TEST\n"),
		strings.ToUpper("vijaylovesbhavu"),
		)
}
