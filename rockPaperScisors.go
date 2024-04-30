package main

import "math/rand"

func game() {
	const rounds int = 3

	for i := 0; i < rounds; i++ {

		computerChoiceNum := rand.Int(rounds)

		var compChoice string

		switch computerChoiceNum {
		case 0:
			compChoice = "Rock"
		case 1:
			compChoice = "Paper"
		case 2:
			compChoice = "Scissors"
		}
	}
}
