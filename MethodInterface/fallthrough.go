package main
import "fmt"
func main() {
	i := 6
	switch i {
	case 3: fmt.Println("was > 3"); fallthrough;
case 4: fmt.Println("was > 4"); fallthrough;
case 5: fmt.Println("was > 5"); fallthrough;
case 6: fmt.Println("was > 6"); fallthrough;
case 7: fmt.Println("was > 7"); fallthrough;
case 8: fmt.Println("was > 8"); fallthrough;
case 9: fmt.Println("was > 9"); fallthrough;
default: fmt.Println("This is default output")
}
}

