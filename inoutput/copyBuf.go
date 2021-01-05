//func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
package main
import (
	"io"
	"os"
	"log"
	"strings"
)
func main() {
	r1 := strings.NewReader("First Reader\n")
	r2 := strings.NewReader("Second Reader\n")
	buf := make([]byte, 8)
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}
