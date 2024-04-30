package main

import "fmt"

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
	//	Pointers - holds memory address of a variable

	// speeding()
}

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
