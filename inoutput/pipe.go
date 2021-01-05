//func Pipe() (*PipeReader, *PipeWriter)
package main
import (
	"io"
	"os"
	"log"
	"fmt"
)
func main() {
	r, w := io.Pipe()
	go func() {
		fmt.Fprint(w, "reading pipe\n")
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
