package main
import (
	"fmt"
	"strconv"
	"time"
)
func makeCakeAndSend(cs chan string, flavour string, count int) {
	for i := 0; i<=10; i++ {
		cakeName := flavour + " Cake " + strconv.Itoa(i)
		cs <- cakeName
	}
	close(cs)
}
func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false
	for {
		if (strbry_closed && choco_closed) { return }
	        fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbry_ok := <- strbry_cs:
			if (!strbry_ok) {
				strbry_closed = true
				fmt.Println(" ... Strbry channel closed!")
			} else {
				fmt.Println("Received from strbry channel.  Now packing", cakeName)
			}
		case cakeName, choco_ok := <- choco_cs:
			if(!choco_ok) {
				choco_closed = true
				fmt.Println("... chocolate channel closed!")
			} else {
				fmt.Println("Received from chacolate channel. Now Packing", cakeName)
			}
		}
	}
}
func main() {
	 strbry_cs := make(chan string)
	 choco_cs := make(chan string)
	go makeCakeAndSend(strbry_cs, "chacolate", 3)
	go makeCakeAndSend(choco_cs, "Strawberry", 4)
//	time.Sleep(4 * 1e9)
	go receiveCakeAndPack(strbry_cs, choco_cs)
	time.Sleep(4 * 1e9)
}


