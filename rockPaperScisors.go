package main

import (
	"fmt"
	"math/rand"
)

func game() {
	const rounds int = 3

	for i := 0; i < rounds; i++ {

		computerChoiceNum := rand.Int()

		var (
			compChoice string
		)

		switch computerChoiceNum {
		case 0:
			compChoice = "Rock"
		case 1:
			compChoice = "Paper"
		case 2:
			compChoice = "Scissors"
		}
		fmt.Println("You chose ", compChoice)
	}
}
