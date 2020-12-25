package main

import "fmt"

func main() {
        var creature string = "shark"
        var pointer *string = &creature

        fmt.Println("creature =", creature)
	fmt.Println("address of var :", &creature)
        fmt.Println("pointer =", pointer)
	fmt.Println("pointer value :", *pointer)
		}
