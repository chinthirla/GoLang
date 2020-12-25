package main
import "fmt"
func main(){
	fmt.Print("Enter temparature in F:")
	var temp float64
	var c float64
	fmt.Scanf("%f", &temp)
	c = ((temp - 32) * 5/9)
	fmt.Println("Temparature in Celcious is :", c)
}

