package main
import "fmt"
func main(){
	fmt.Println(f1())
}
func f1 () int {
	return f2()
}
func f2 () int {
	return f3 ()
}	
func f3() (r int) {
	r = 1
	return
}


