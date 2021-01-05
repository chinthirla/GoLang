//copyN( dst writer, src Reader, n int64) (written int64, err error)
package main
import (
	"io"
	"os"
	"strings"
	"log"
)
func main() {
	r1 := strings.NewReader("First Reader\n")
///	r2 := strings.NewReader("second Reader\n")
	if _, err := io.CopyN(os.Stdout, r1, 20); err != nil {
		log.Fatal(err)
	}
}

