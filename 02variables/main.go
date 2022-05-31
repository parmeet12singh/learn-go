package main

import "fmt"

func main() {

	var username string = "parmeet"
	fmt.Println(username)
	fmt.Printf("Varibale is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varibale is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Varibale is of type: %T \n", smallVal)

	var smallFloat float64 = 255.23423423423423423
	fmt.Println(smallFloat)
	fmt.Printf("Varibale is of type: %T \n", smallFloat)
}
