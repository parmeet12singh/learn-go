package main

import "fmt"

func main() {

	fmt.Println("Welcome to If Else")

	loginCount := 10

	var result string

	if loginCount > 10 {
		result = "Login Count greater than 10"
	} else if loginCount == 10 {
		result = "Login Count equal to 10"
	} else {
		result = "Login Count less than 10"
	}

	fmt.Println("Result: ", result)

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num := 9; num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}
}
