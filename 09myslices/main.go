package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to Slices")

	var fruits = []string{"Apple", "Orange", "Banana"}

	fmt.Printf("Type of Fruits List: %T\n", fruits)

	fruits = append(fruits, "Mango", "Grapes")
	fmt.Println("Fruits: ", fruits)

	fruits = append(fruits[:2])
	fmt.Println("Fruits: ", fruits)

	highScores := make([]int, 4)

	highScores[0] = 100
	highScores[1] = 90
	highScores[2] = 80
	highScores[3] = 70

	highScores = append(highScores, 60, 50)
	fmt.Println("High Scores: ", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"Go", "Python", "Java", "C++", "C#"}
	fmt.Println("Courses: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses: ", courses)
}
