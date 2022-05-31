package main

import "fmt"

func main() {

	fmt.Println("Welcome to Array")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[3] = "Banana"

	fmt.Println("Fruits: ", fruits)
	fmt.Println("Lenght of Fruits List: ", len(fruits))

	var vegetables = [5]string{"Potato", "Tomato", "Cucumber", "Onion"}
	fmt.Println("Vegetables: ", vegetables)
	fmt.Println("Lenght of Vegetables List: ", len(vegetables))

}
