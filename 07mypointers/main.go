package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer: ", ptr)

	number := 12
	var ptr = &number

	fmt.Println("Value of actual pointer: ", ptr)
	fmt.Println("Value of actual pointer: ", *ptr)

	*ptr = *ptr + 1
	fmt.Println("Value of number: ", number)
}
