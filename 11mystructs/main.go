package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	parmeet := User{"Parmeet", "parmeet@go.dev", true, 25}
	fmt.Println(parmeet)
	fmt.Printf("Parmeet details are %+v\n", parmeet)
	fmt.Printf("Name is %v and email is %v\n.", parmeet.Name, parmeet.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
