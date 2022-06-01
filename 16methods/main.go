package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	parmeet := User{"Parmeet", "parmeet@go.dev", true, 25}
	fmt.Println(parmeet)
	fmt.Printf("Parmeet details are %+v\n", parmeet)
	fmt.Printf("Name is %v and email is %v\n", parmeet.Name, parmeet.Email)

	parmeet.GetStatus()
	parmeet.NewMail()

	fmt.Printf("Name is %v and email is %v\n", parmeet.Name, parmeet.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New mail : ", u.Email)
}
