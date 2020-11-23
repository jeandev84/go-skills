package main

import (
	"fmt"
)

func main() {
	weekday := 6 // 1 == Monday, 7 == Sunday
	fmt.Printf("Day od the weekd=%d. What's for today?\n", weekday)

	switch weekday {

	case 1:
		fmt.Println("Beginning of the week, let's get to work!")
	case 3:
		fmt.Println("Wednesday, the half is done!")
	case 6, 7: // si weekDay = 6 ou 7
		fmt.Println("It's the week-end!")
	default:
		fmt.Println("Nothing special about this day...")
	}

	// Timing
	hour := 10
	fmt.Printf("Current time=%d\n", hour)

	switch {
	case hour < 12:
		fmt.Println("It's the morning")
	case hour > 12 && hour < 18:
		fmt.Println("It's the afternoon")
	default:
		fmt.Println("It's the evening")
	}
}
