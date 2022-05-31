package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time package")

	presentTime := time.Now()
	fmt.Println("Present Time: ", presentTime)

	fmt.Println("Year: ", presentTime.Year())
	fmt.Println("Month: ", presentTime.Month())
	fmt.Println("Day: ", presentTime.Day())
	fmt.Println("Hour: ", presentTime.Hour())
	fmt.Println("Minute: ", presentTime.Minute())
	fmt.Println("Second: ", presentTime.Second())

	fmt.Println("Weekday: ", presentTime.Weekday())

	fmt.Println("Location: ", presentTime.Location())

	fmt.Println("Unix: ", presentTime.Unix())

	fmt.Println("My Format: ", presentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("My Format: ", presentTime.Format("2006-01-02 15:04:05 Monday"))

	createdDate := time.Date(1998, time.May, 12, 10, 30, 0, 0, time.UTC)
	fmt.Println("Created Date: ", createdDate)
	fmt.Println("Created Date: ", createdDate.Format("2006-01-02 15:04:05 Monday"))

}
