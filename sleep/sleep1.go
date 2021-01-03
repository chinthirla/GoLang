package main
import (
	"fmt"
	"time"
)
func main() {
	time.Sleep(8 * time.Second)
	fmt.Println("Printing every 8 seconds")
	time.Sleep(5 * time.Second)
	fmt.Println("second print after sleep")
	}

