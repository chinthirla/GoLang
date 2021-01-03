package main
import (
	"fmt"
	"time"
)
func Pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func Another (c chan <- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func Printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go Pinger(c)
	go Another(c)
	go Printer(c)
	var input string
	fmt.Scanln(&input)
}

