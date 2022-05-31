package main

import "fmt"

func main() {

	fmt.Println("Welcome to Maps")

	languages := make(map[string]string)

	languages["Go"] = "Golang"
	languages["Py"] = "Python"
	languages["Js"] = "JavaScript"

	fmt.Println("Languages: ", languages)
	fmt.Println("Go: ", languages["Go"])

	delete(languages, "Go")
	fmt.Println("Languages: ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
