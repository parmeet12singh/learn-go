package main

import "fmt"

func main() {

	fmt.Println("Welcome to Loops")

	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	fmt.Println("Days of Week:", days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i, day := range days {
		fmt.Printf("%d: %s\n", i, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			break
			// rogueValue++
			// continue
		}
		fmt.Println("Rogue Value:", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Here is the GoTo Label")
}
