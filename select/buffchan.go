package main
import (
	"fmt"
	"time"
)
func made(cs chan string) {
	cs <- "from made"
	fmt.Println(cs)
	time.Sleep(time.Second * 4)
}
func made1(cs chan string) {
	cs <- "from made2"
	time.Sleep(time.Second * 3)
	fmt.Println(cs)
}
func main() {
	c1 := make(chan string, 20)
	go made(c1)
	go made1(c1)
	var input string
	fmt.Scanln(&input)
}
