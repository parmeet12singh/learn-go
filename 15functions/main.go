package main

import "fmt"

func main() {

	fmt.Println("Welcome to Functions")
	greeter()
	greeter2()

	result := add(1, 2)
	fmt.Println("The result is:", result)

	result2, myMessage := proAdd(1, 2, 3, 4, 5)
	fmt.Println("The result is:", result2)
	fmt.Println("My Message is:", myMessage)
}

func greeter() {
	fmt.Println("Hello World from greeter")
}

func greeter2() {
	fmt.Println("Hello World from greeter2")
}

func add(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdd(values ...int) (int, string) {

	total := 0

	for _, value := range values {
		total += value
	}

	return total, "ProAdd"

}
