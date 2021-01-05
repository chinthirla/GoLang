//func Copy(dst Writer, src Reader) (written int64, err error)
package main
import (
	"os"
	"log"
	"strings"
	"io"
)
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
					}
				}
