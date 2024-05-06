package main

import (
	"fmt"
	"time"
)

func getUserAge() {

	const userAge int = 18

	fmt.Print("How Old Are You? ")
	var age int
	fmt.Scan(&age)

	if age >= userAge {
		fmt.Println("You are", age, "years old", "You can attend Parties")
	} else if age < userAge {
		fmt.Println("You are", age, "years old", "You can not attend Parties!")
	} else {
		fmt.Println("Toddler")
	}

}

//	Pointers - holds memory address of a variable

func speeding() {
	isFast := false
	isSlow := false

	if isFast {
		fmt.Println("Your Going too fast slow down")
	} else if isSlow {
		fmt.Println("Speed up !!")
	} else {
		fmt.Println("Not Speeding")
	}

}

func week() {

	isWeekend := false // bool
	isHoliday := false

	if isWeekend && isHoliday {
		fmt.Println("Its a holiday on a weekend ")
	} else if isWeekend && !isHoliday {
		fmt.Println("Its a regular weekend ")
	} else {
		fmt.Println("Its a regular Week Day")
	}

}

func timeNow() {
	const user string = "Alexander"
	tme := time.Now()
	fmt.Println(tme)

	switch {
	case tme.Hour() < 12:
		fmt.Println("Good Morning,", user)
	case tme.Hour() >= 12 && tme.Hour() < 17:
		fmt.Println("Good Afternoon,", user)
	case tme.Hour() >= 17:
		fmt.Println("Good Evening,", user)
	default:
		fmt.Println("What's the time ?,", user)
	}
}
