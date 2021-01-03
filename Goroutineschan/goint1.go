package main
import (
	"fmt"
	"time"
	)
func simulateEvent(name string, timeInSecs int64) { 
    // sleep for a while to simulate time consumed by event
       fmt.Println("Started ", name, ": Should take", timeInSecs, "seconds.")
       time.Sleep(time.Duration(timeInSecs * 1e9))
//       time.Sleep(12 * 1e9)
      // time.Sleep(time.Duration(timeInSecs.Int31n(1000)) * time.Millisecond)
       fmt.Println("Finished ", name)
}
func add2numbers(a, b int) {
	fmt.Println(a + b)
}
func main() {
	add2numbers(1, 3)
	go add2numbers(4, 5)
	simulateEvent("100m sprint", 10) //start 100m sprint, it should take 10 seconds
        simulateEvent("Long jump", 6) //start long jump, it should take 6 seconds
        simulateEvent("High jump", 3)
	time.Sleep(12 * 1e9)
}

