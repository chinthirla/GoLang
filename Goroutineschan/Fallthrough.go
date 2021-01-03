package main
import "fmt"
func main() {
	i := 2
	fmt.Println("swaitch output is:", ls(i))
}
func ls(n int) int {
	switch (n) {
	case 1: fmt.Printf("%d is lesser than 50\n", n)
case 2: fmt.Printf("%d is lesser than 50\n", n); fallthrough
case 3: fmt.Printf("%d is lesser than 50\n", n); fallthrough
default : fmt.Printf("Enter valid number\n")
}
return n
}

